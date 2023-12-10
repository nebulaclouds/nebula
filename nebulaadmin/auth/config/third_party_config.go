package config

// This struct encapsulates config options for bootstrapping various Nebula applications with config values
// For example, NebulaClientConfig contains application-specific values to initialize the config required by nebula client
type ThirdPartyConfigOptions struct {
	NebulaClientConfig NebulaClientConfig `json:"nebulaClient"`
}

type NebulaClientConfig struct {
	ClientID    string   `json:"clientId" pflag:",public identifier for the app which handles authorization for a Nebula deployment"`
	RedirectURI string   `json:"redirectUri" pflag:",This is the callback uri registered with the app which handles authorization for a Nebula deployment"`
	Scopes      []string `json:"scopes" pflag:",Recommended scopes for the client to request."`
	Audience    string   `json:"audience" pflag:",Audience to use when initiating OAuth2 authorization requests."`
}

func (o ThirdPartyConfigOptions) IsEmpty() bool {
	return len(o.NebulaClientConfig.ClientID) == 0 && len(o.NebulaClientConfig.RedirectURI) == 0 && len(o.NebulaClientConfig.Scopes) == 0
}
