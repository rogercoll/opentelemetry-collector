package templateconverter

import (
	"context"
	"fmt"
	"sort"
	"strings"
	"text/template"

	"go.opentelemetry.io/collector/confmap"
)

// Define a template type by adding a "templates" section to the config.
//
// templates:
//
//	my_template: |
//	  receivers: ... # required, one or more
//	  processors: ... # optional
//	  pipelines: ...  # required if processors are used, otherwise optional
//
// Instantiate a templated receivers:
//
// receivers:
//
//	template/my_template[/name]:
//	  parameter_one: value_one
//	  parameter_two: value_two
type converter struct {
	cfg       map[string]any
	templates map[string]map[string]*template.Template
}

// New returns a confmap.Converter, that renders all templates and inserts them into the given confmap.Conf.
func New() confmap.Converter {
	return &converter{
		templates: make(map[string]map[string]*template.Template),
	}
}

// TODO: use component.Type
func (c *converter) expandComponents(componentType string) error {
	componentCfgs, ok := c.cfg[componentType].(map[string]any)
	if !ok {
		return nil // invalid, but let the unmarshaler handle it
	}
	for templateID, parameters := range componentCfgs {
		if !strings.HasPrefix(templateID, "template") {
			continue
		}

		id, err := newInstanceID(templateID)
		if err != nil {
			return err
		}

		tmpl, ok := c.templates[id.tmplType]
		if !ok {
			return fmt.Errorf("template type %q not found", id.tmplType)
		}

		cfg, err := newTemplateConfig(tmpl[componentType], parameters)
		if err != nil {
			return err
		}

		c.expandTemplate(componentType, cfg, id)
	}
	return nil
}

func (c *converter) expandPipeline() error {
	pipelinesMap, ok := c.cfg["service"].(map[string]any)["pipelines"].(map[string]any)
	if !ok {
		return nil
	}
	for templateID, parameters := range pipelinesMap {
		if !strings.HasPrefix(templateID, "template") {
			continue
		}

		id, err := newInstanceID(templateID)
		if err != nil {
			return err
		}

		tmpl, ok := c.templates[id.tmplType]
		if !ok {
			return fmt.Errorf("template type %q not found", id.tmplType)
		}

		cfg, err := newTemplateConfig(tmpl["pipelines"], parameters)
		if err != nil {
			return err
		}

		delete(c.cfg["service"].(map[string]any)["pipelines"].(map[string]any), id.withPrefix("template"))

		for componentID, componentCfg := range *cfg {
			c.cfg["service"].(map[string]any)["pipelines"].(map[string]any)[componentID] = componentCfg
		}
	}
	return nil
}

func (c *converter) expandTemplate(componentType string, cfg *templateConfig, id *instanceID) {
	// Delete the reference to this template instance and
	// replace it with the rendered receiver(s).
	delete(c.cfg[componentType].(map[string]any), id.withPrefix("template"))

	componentIDs := make([]any, len(*cfg))
	i := 0
	for componentID, componentCfg := range *cfg {
		c.cfg[componentType].(map[string]any)[componentID] = componentCfg
		componentIDs[i] = componentID
		i++
	}

	// update pipeline references
	pipelinesMap := c.cfg["service"].(map[string]any)["pipelines"].(map[string]any)
	for _, pipelineMap := range pipelinesMap {
		oldCompIDs, ok := pipelineMap.(map[string]any)[componentType].([]any)
		if !ok {
			continue
		}
		newCompIDs := make([]any, 0, len(oldCompIDs))
		for _, receiverID := range oldCompIDs {
			if receiverID != id.withPrefix("template") {
				newCompIDs = append(newCompIDs, receiverID)
			}
		}

		if len(newCompIDs) == len(oldCompIDs) {
			// This pipeline did not use the template
			continue
		}

		newCompIDs = append(newCompIDs, componentIDs...)

		// This makes tests deterministic
		sort.Slice(newCompIDs, func(i, j int) bool {
			return newCompIDs[i].(string) < newCompIDs[j].(string)
		})
		pipelineMap.(map[string]any)[componentType] = newCompIDs
	}
}

func (c *converter) Convert(_ context.Context, conf *confmap.Conf) error {
	if err := c.parseTemplates(conf); err != nil {
		return err
	}

	c.cfg = conf.ToStringMap()

	// expand pipelines templates
	err := c.expandPipeline()
	if err != nil {
		return err
	}

	// replace receivers configurations + pipeline usages
	err = c.expandComponents("receivers")
	if err != nil {
		return err
	}

	// replace processors configurations + pipeline usages
	err = c.expandComponents("processors")
	if err != nil {
		return err
	}

	// replace exporters configurations + pipeline usages
	err = c.expandComponents("exporters")
	if err != nil {
		return err
	}

	delete(c.cfg, "templates")

	*conf = *confmap.NewFromStringMap(c.cfg)
	return nil
}

func (c *converter) parseTemplates(conf *confmap.Conf) error {
	if !conf.IsSet("templates") {
		return nil
	}

	templatesMap, ok := conf.ToStringMap()["templates"].(map[string]any)
	if !ok {
		return fmt.Errorf("'templates' must be a map")
	}

	c.templates = make(map[string]map[string]*template.Template, len(templatesMap))

	strToComponentsTemplate := func(componentType string, templateVal map[string]any) (*template.Template, error) {
		templateStr, ok := templateVal[componentType].(string)
		if !ok {
			return nil, fmt.Errorf("'templates::*::%s' must be a string", componentType)
		}
		parsedTemplate, err := template.New(componentType).Parse(templateStr)
		if err != nil {
			return nil, err
		}

		return parsedTemplate, nil
	}

	for templateName, templateVal := range templatesMap {
		componentsRawTemplate, ok := templateVal.(map[string]any)
		if !ok {
			return fmt.Errorf("'templates::%s::' must be a map", templateName)
		}
		// skip error for now
		receiverGroup, _ := strToComponentsTemplate("receivers", componentsRawTemplate)
		processorGroup, _ := strToComponentsTemplate("processors", componentsRawTemplate)
		exportersGroup, _ := strToComponentsTemplate("exporters", componentsRawTemplate)
		pipelinesGroup, _ := strToComponentsTemplate("pipelines", componentsRawTemplate)

		c.templates[templateName] = map[string]*template.Template{
			"receivers":  receiverGroup,
			"processors": processorGroup,
			"exporters":  exportersGroup,
			"pipelines":  pipelinesGroup,
		}

	}

	return nil
}
