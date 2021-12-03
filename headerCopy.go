package copy_header_value_traefik_plugin

import (
	"context"
	"fmt"
	"net/http"
	"strings"
)

// Config the plugin configuration.
type Config struct {
	From              string `json:"from,omitempty"`
	PairSeparator     string `json:"pairSeparator,omitempty"`
	KeyValueSeparator string `json:"keyValueSeparator,omitempty"`
	Key               string `json:"key,omitempty"`
	To                string `json:"to,omitempty"`
	Prefix            string `json:"prefix,omitempty"`
	Overwrite         bool   `json:"overwrite,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{}
}

type CopyHeaderPlugin struct {
	next   http.Handler
	name   string
	config *Config
}

// New creates a new plugin.
func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.From == "" {
		return nil, fmt.Errorf("from cannot be empty")
	}
	if config.PairSeparator == "" {
		return nil, fmt.Errorf("pairSeparator cannot be empty")
	}
	if config.KeyValueSeparator == "" {
		return nil, fmt.Errorf("keyValueSeparator cannot be empty")
	}
	if config.Key == "" {
		return nil, fmt.Errorf("key cannot be empty")
	}
	if config.To == "" {
		return nil, fmt.Errorf("to cannot be empty")
	}

	return &CopyHeaderPlugin{
		next: next, config: config, name: name,
	}, nil
}

func (copyHeaderPlugin *CopyHeaderPlugin) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	copyHeaderValue(&req.Header, copyHeaderPlugin.config)
	copyHeaderPlugin.next.ServeHTTP(rw, req)
}

func copyHeaderValue(headers *http.Header, config *Config) {
	if headers.Get(config.To) != "" && !config.Overwrite {
		return
	}

	headerValue := headers.Get(config.From)
	if headerValue != "" {
		pairs := strings.Split(headerValue, config.PairSeparator)
		for _, pair := range pairs {
			keyValue := strings.Split(strings.TrimSpace(pair), config.KeyValueSeparator)
			if len(keyValue) == 2 && keyValue[0] == config.Key {
				if config.Prefix != "" {
					keyValue[1] = config.Prefix + keyValue[1]
				}
				headers.Set(config.To, keyValue[1])
			}
		}
	}
}
