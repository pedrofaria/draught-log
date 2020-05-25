package config

import (
	"context"
	"errors"
	"strings"

	"github.com/pedrofaria/draught-log/internal/pkg/config"
	"github.com/pedrofaria/draught-log/internal/pkg/stream"
	"github.com/pedrofaria/draught-log/internal/pkg/stream/formatter"
	"github.com/pedrofaria/draught-log/internal/pkg/stream/provider"
	"github.com/pedrofaria/draught-log/internal/pkg/stream/types"
)

type cfgProvider interface {
	Stream(context.Context, chan<- types.Message) error
}

func RegisterProvidersFromConfig(manager *stream.Manager, cfg *config.Config) error {
	for _, resource := range cfg.Resources {
		var p cfgProvider

		switch strings.ToLower(resource.Provider) {
		case "docker":
			p = provider.NewDockerProvider(resource.Name, resource.ProviderID)
		case "file":
			p = provider.NewFileProvider(resource.Name, resource.ProviderID)
		default:
			return errors.New("unknown provider")
		}

		var f cfgProvider

		switch strings.ToLower(resource.Formatter.Type) {
		case "json":
			f = formatter.NewJSON(p, formatter.Config{
				MessageField:          resource.Formatter.MessageField,
				LevelField:            resource.Formatter.LevelField,
				TimestampField:        resource.Formatter.TimestampField,
				TimestampFormat:       resource.Formatter.TimestampFormat,
				PreFilterRegex:        resource.Formatter.PreFilterRegex,
				PreFilterRegexReplace: resource.Formatter.PreFilterRegexReplace,
			})
		default:
			return errors.New("unknown formatter type")
		}

		manager.RegisterProvider(f)
	}

	return nil
}
