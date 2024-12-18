// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

package builders // import "go.opentelemetry.io/collector/service/internal/builders"

import (
	"errors"

	"go.uber.org/zap"

	"go.opentelemetry.io/collector/component"
)

var (
	errNilNextConsumer = errors.New("nil next Consumer")
	nopType            = component.MustNewType("nop")
)

// logStabilityLevel logs the stability level of a component. The log level is set to info for
// undefined, unmaintained, deprecated and development. The log level is set to debug
// for alpha, beta and stable.
func logStabilityLevel(logger *zap.Logger, sl component.StabilityLevel) {
	if sl >= component.StabilityLevelAlpha {
		logger.Debug(sl.LogMessage())
	} else {
		logger.Info(sl.LogMessage())
	}
}

// type ConfigStore struct {
// 	cfgs map[component.ID]component.Config
// }

type ConfigStore map[component.ID]component.Config

func (c ConfigStore) AddComponent(recvID component.ID, conf component.Config) {
	// TODO check if already exists
	c[recvID] = conf
}

func (c ConfigStore) RemoveComponent(recvID component.ID) {
	// TODO check if already exists
	delete(c, recvID)
}
