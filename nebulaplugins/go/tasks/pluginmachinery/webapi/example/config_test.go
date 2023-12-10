package example

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nebulaclouds/nebula/nebulastdlib/config"
	"github.com/nebulaclouds/nebula/nebulastdlib/config/viper"
)

func TestGetConfig(t *testing.T) {
	configAccessor := viper.NewAccessor(config.Options{
		StrictMode:  true,
		SearchPaths: []string{"testdata/admin_plugin.yaml"},
	})

	err := configAccessor.UpdateConfig(context.TODO())
	assert.NoError(t, err)

	cfg := GetConfig()
	assert.Len(t, cfg.WebAPI.ResourceQuotas, 1)
}
