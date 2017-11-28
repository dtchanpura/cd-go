package config

// Configuration for whole configuration file
type Configuration struct {
	ServeConfig    Serve     `yaml:"serve" json:"serve"`
	ProjectConfigs []Project `yaml:"projects,omitempty" json:"projects,omitempty"`
}

// Serve for server (listener) related configurations
type Serve struct {
	Host string `yaml:"host" json:"host"`
	Port int    `yaml:"port" json:"port"`
}

// Project for all projects configurations
type Project struct {
	Name       string        `yaml:"name" json:"name"`
	UUID       string        `yaml:"uuid" json:"uuid"`
	WorkDir    string        `yaml:"work_dir" json:"work_dir"`
	RemotePath string        `yaml:"remote_path" json:"remote_path"` // For downloading
	PostHook   string        `yaml:"post_hook,omitempty" json:"post_hook,omitempty"`
	PreHook    string        `yaml:"pre_hook,omitempty" json:"pre_hook,omitempty"`
	ErrorHook  string        `yaml:"error_hook,omitempty" json:"error_hook,omitempty"`
	Secret     string        `yaml:"secret" json:"secret"`
	Tokens     []TokenDetail `yaml:"tokens" json:"tokens"`
}

// TokenDetail is for allowing multiple ips to access same
// repository with different tokens
type TokenDetail struct {
	Token              string `yaml:"token" json:"token"`
	WhitelistedNetwork string `yaml:"whitelist_net" json:"whitelistnet"` // CIDR notation
	Name               string `yaml:"name" json:"name"`
}
