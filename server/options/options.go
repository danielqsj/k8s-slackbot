package options

import (
	_ "net/http/pprof"

	"github.com/spf13/pflag"
)

// SlackBotServerConfig configures and runs a Kubernetes Slack Bot server
type SlackBotServerConfig struct {
	KubeConfigFile string
	HealthzPort    int
	ConfigMap      string
	BotToken       string
	DebugEnable    bool
}

// NewSlackBotServerConfig returns default config
func NewSlackBotServerConfig() *SlackBotServerConfig {
	return &SlackBotServerConfig{
		KubeConfigFile: "",
		HealthzPort:    8081,
		ConfigMap:      "",
		BotToken:       "",
		DebugEnable:    false,
	}
}

// AddFlags adds flags for a specific ProxyServer to the specified FlagSet
func (c *SlackBotServerConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&c.KubeConfigFile, "kubecfg-file", c.KubeConfigFile, "Location of kubecfg file for access to kubernetes master service; --kube-master-url overrides the URL part of this; if neither this nor --kube-master-url are provided, defaults to service account tokens")
	fs.IntVar(&c.HealthzPort, "healthz-port", c.HealthzPort, "Port on which to serve an App Monitor HTTP readiness probe.")
	fs.StringVar(&c.ConfigMap, "configmap", c.ConfigMap, "Name of the ConfigMap that contains the custom configuration to use")
	fs.StringVar(&c.BotToken, "bot-token", c.BotToken, "Token of slack bot to use")
	fs.BoolVar(&c.DebugEnable, "debug", c.DebugEnable, "Whether enable debug log")
}
