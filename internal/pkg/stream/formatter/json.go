package formatter

import (
	"context"
	"encoding/json"
	"regexp"
	"strings"
	"time"

	"github.com/pedrofaria/draught-log/internal/pkg/stream/types"
)

type provider interface {
	Stream(context.Context, chan<- types.Message, types.Metadata) error
}

type JSON struct {
	provider provider
	config   Config
}

func NewJSON(p provider, c Config) *JSON {
	return &JSON{
		provider: p,
		config:   c,
	}
}

func (formatter *JSON) Stream(ctx context.Context, sendChan chan<- types.Message, metadata types.Metadata) error {
	middleChan := make(chan types.Message)
	errChan := make(chan error)

	go func(p provider) {
		if err := p.Stream(ctx, middleChan, metadata); err != nil {
			errChan <- err
			return
		}
	}(formatter.provider)

loop:
	for {
		select {
		case err := <-errChan:
			return err
		case msg := <-middleChan:
			line := formatter.parse(msg)
			if strings.Trim(line.Message, " ") == "" {
				continue
			}
			sendChan <- line
		case <-ctx.Done():
			break loop
		}
	}

	return nil
}
func (formatter *JSON) parse(msg types.Message) types.Message {
	msg.Processed = msg.RawLog

	msg.Processed = strings.Trim(msg.Processed, "\r\n \u0000\u0001\u0006")

	if formatter.config.PreFilterRegex != "" {
		preFilterRe := regexp.MustCompile(formatter.config.PreFilterRegex)
		msg.Processed = preFilterRe.ReplaceAllString(msg.Processed, formatter.config.PreFilterRegexReplace)
	}

	var payload map[string]interface{}

	if err := json.Unmarshal([]byte(msg.Processed), &payload); err != nil {
		msg.Message = msg.Processed
		return msg
	}

	msg.Payload = payload

	if v, ok := payload[formatter.config.MessageField]; !ok {
		msg.Message = strings.Trim(msg.Processed, "\r\n ")
		return msg
	} else {
		msg.Message = strings.Trim(v.(string), "\r\n ")
	}

	if v, ok := payload[formatter.config.LevelField]; ok {
		msg.Level = v.(string)
	}

	if v, ok := payload[formatter.config.TimestampField]; ok {
		if t, err := time.Parse(formatter.config.TimestampFormat, v.(string)); err == nil {
			msg.Timestamp = t
		}
	}

	return msg
}
