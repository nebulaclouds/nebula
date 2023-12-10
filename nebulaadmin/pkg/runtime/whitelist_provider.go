package runtime

import (
	"github.com/nebulaclouds/nebula/nebulaadmin/pkg/runtime/interfaces"
	"github.com/nebulaclouds/nebula/nebulastdlib/config"
)

const whitelistKey = "task_type_whitelist"

var whiteListProviderDefault = make(map[string][]interfaces.WhitelistScope)

var whitelistConfig = config.MustRegisterSection(whitelistKey, &whiteListProviderDefault)

// Implementation of an interfaces.QueueConfiguration
type WhitelistConfigurationProvider struct{}

func (p *WhitelistConfigurationProvider) GetTaskTypeWhitelist() interfaces.TaskTypeWhitelist {
	whitelists := whitelistConfig.GetConfig().(*interfaces.TaskTypeWhitelist)
	return *whitelists
}

func NewWhitelistConfigurationProvider() interfaces.WhitelistConfiguration {
	return &WhitelistConfigurationProvider{}
}
