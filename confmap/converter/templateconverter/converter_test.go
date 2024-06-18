package templateconverter

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/confmap/confmaptest"
)

func TestTemplateConverter(t *testing.T) {
	for _, tc := range []string{
		"no-templates",
		"single-receiver",
		"single-processor",
		"single-exporter",
		"multiple-receivers-no-pipelines",
		"all-components",
		"multiple-templates",
	} {
		t.Run(tc, func(t *testing.T) {
			conf, err := confmaptest.LoadConf(filepath.Join("testdata", "valid", tc, "config.yaml"))
			require.NoError(t, err)

			expected, err := confmaptest.LoadConf(filepath.Join("testdata", "valid", tc, "expected.yaml"))
			require.NoError(t, err)

			require.NoError(t, New().Convert(context.Background(), conf))
			expectedMap := expected.ToStringMap()
			actualMap := conf.ToStringMap()
			assert.Equal(t, expectedMap, actualMap)
		})
	}
}
