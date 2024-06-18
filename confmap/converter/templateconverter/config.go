// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package templateconverter // import "go.opentelemetry.io/collector/confmap/converter/templateconverter"

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"
)

type instanceID struct {
	tmplType string
	instName string
}

func newInstanceID(id string) (*instanceID, error) {
	parts := strings.SplitN(id, "/", 3)
	switch len(parts) {
	case 2: // template/type
		return &instanceID{
			tmplType: parts[1],
		}, nil
	case 3: // template/type/name
		return &instanceID{
			tmplType: parts[1],
			instName: parts[2],
		}, nil
	}
	return nil, fmt.Errorf("'template' must be followed by type")
}

func (id *instanceID) withPrefix(prefix string) string {
	if id.instName == "" {
		return prefix + "/" + id.tmplType
	}
	return prefix + "/" + id.tmplType + "/" + id.instName
}

type templateConfig = map[string]any

func newTemplateConfig(tmpl *template.Template, parameters any) (*templateConfig, error) {
	var rendered bytes.Buffer
	var err error
	if err = tmpl.Execute(&rendered, parameters); err != nil {
		return nil, fmt.Errorf("render: %w", err)
	}

	cfg := new(templateConfig)
	if err = yaml.Unmarshal(rendered.Bytes(), cfg); err != nil {
		return nil, fmt.Errorf("malformed: %w", err)
	}

	return cfg, nil
}
