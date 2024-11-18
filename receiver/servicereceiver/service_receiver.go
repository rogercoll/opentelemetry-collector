// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package servicereceiver // import "go.opentelemetry.io/collector/receiver/nopreceiver"

import (
	"context"
	"fmt"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/confmap"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/otelcol"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/service"
)

// NewFactory returns a receiver.Factory that constructs nop receivers.
func NewFactory() receiver.Factory {
	return receiver.NewFactory(
		component.MustNewType("service"),
		func() component.Config { return &struct{}{} },
		receiver.WithTraces(createTraces, component.StabilityLevelDevelopment),
		receiver.WithMetrics(createMetrics, component.StabilityLevelDevelopment),
		receiver.WithLogs(createLogs, component.StabilityLevelDevelopment))
}

func createTraces(ctx_ context.Context, set receiver.Settings, c component.Config, ct consumer.Traces) (receiver.Traces, error) {
	return NewReceiver(set.TelemetrySettings, *confmap.NewFromStringMap(map[string]any{
		"receivers": map[string]any{
			"otlp": "",
		},
		"processors": map[string]any{
			"batch": "",
		},
	})), nil
}

func createMetrics(ctx_ context.Context, set receiver.Settings, c component.Config, ct consumer.Metrics) (receiver.Metrics, error) {
	return NewReceiver(set.TelemetrySettings, *confmap.NewFromStringMap(map[string]any{
		"receivers": map[string]any{
			"otlp": "",
		},
		"processors": map[string]any{
			"batch": "",
		},
	})), nil
}

func createLogs(ctx_ context.Context, set receiver.Settings, c component.Config, ct consumer.Logs) (receiver.Logs, error) {
	return NewReceiver(set.TelemetrySettings, *confmap.NewFromStringMap(map[string]any{
		"receivers": map[string]any{
			"otlp": "",
		},
		"processors": map[string]any{
			"batch": "",
		},
	})), nil
}

var nopInstance = &serviceReceiver{}

type serviceReceiver struct {
	serviceConfig otelcol.Config
	innterCtx     context.Context
	innerService  *service.Service

	rawConfig confmap.Conf
}

// factoryGetter is an interface that the component.Host passed to receivercreator's Start function must implement
// GetFactory is optional in hosts since 107.0, but we require it.
type factoryGetter interface {
	component.Host
	GetFactory(component.Kind, component.Type) component.Factory
}

func NewReceiver(telemetrySettings component.TelemetrySettings, rawConfig confmap.Conf) *serviceReceiver {
	return &serviceReceiver{
		rawConfig: rawConfig,
	}
}

// func componentsFactory[C any]() func() map[component.Type]C {
// 	return func() map[component.Type]C {
// 		factories := make(map[component.Type]C)
// 		for id, innerRecevier := range s.serviceConfig.Receivers {
// 			factory := factoryGetter.GetFactory(component.KindReceiver, id.Type())
// 			recfactory := factory.(C)
// 			factories[id.Type()] = recfactory
// 		}
// 		return factories
// 	}
// }

// Shutdown is a method to turn off receiving.
func (s *serviceReceiver) Start(ctx context.Context, host component.Host) error {
	factoryGetter, ok := host.(factoryGetter)
	if !ok {
		return fmt.Errorf("integrationreceiver is not compatible with the provided component.Host")
	}

	if err := s.rawConfig.Marshal(s.serviceConfig); err != nil {
		return fmt.Errorf("could not marshal configuration: %w", err)
	}

	receversConfig := func() map[component.Type]receiver.Factory {
		factories := make(map[component.Type]receiver.Factory)
		for id := range s.serviceConfig.Receivers {
			factory := factoryGetter.GetFactory(component.KindReceiver, id.Type())
			recfactory := factory.(receiver.Factory)
			factories[id.Type()] = recfactory
		}
		return factories
	}

	processorsFactories := func() map[component.Type]processor.Factory {
		factories := make(map[component.Type]processor.Factory)
		for id := range s.serviceConfig.Receivers {
			factory := factoryGetter.GetFactory(component.KindProcessor, id.Type())
			recfactory := factory.(processor.Factory)
			factories[id.Type()] = recfactory
		}
		return factories
	}

	var err error
	s.innerService, err = service.New(ctx, service.Settings{
		ReceiversConfigs:   s.serviceConfig.Receivers,
		ReceiversFactories: receversConfig(),

		ProcessorsConfigs:   s.serviceConfig.Processors,
		ProcessorsFactories: processorsFactories(),
	}, service.Config{})
	if err != nil {
		return err
	}

	return s.innerService.Start(ctx)
}

// Shutdown is a method to turn off receiving.
func (s *serviceReceiver) Shutdown(ctx context.Context) error {
	return s.innerService.Shutdown(ctx)
}
