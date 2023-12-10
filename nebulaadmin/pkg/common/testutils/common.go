package testutils

import "github.com/nebulaclouds/nebula/nebulaidl/gen/pb-go/nebulaidl/admin"

// Convenience method to wrap verbose boilerplate for initializing a PluginOverrides MatchingAttributes.
func GetPluginOverridesAttributes(vals map[string][]string) *admin.MatchingAttributes {
	overrides := make([]*admin.PluginOverride, 0, len(vals))
	for taskType, pluginIDs := range vals {
		overrides = append(overrides, &admin.PluginOverride{
			TaskType: taskType,
			PluginId: pluginIDs,
		})
	}
	return &admin.MatchingAttributes{
		Target: &admin.MatchingAttributes_PluginOverrides{
			PluginOverrides: &admin.PluginOverrides{
				Overrides: overrides,
			},
		},
	}
}
