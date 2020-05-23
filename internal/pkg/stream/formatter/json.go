package formatter

import (
	"context"
	"encoding/json"
	"log"
	"regexp"
	"time"

	"github.com/pedrofaria/draught-log/internal/pkg/stream/types"
)

type provider interface {
	Stream(context.Context, chan<- types.Message) error
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

func (formatter *JSON) Stream(ctx context.Context, sendChan chan<- types.Message) error {
	middleChan := make(chan types.Message)

	go func(p provider) {
		if err := p.Stream(ctx, middleChan); err != nil {
			log.Printf("Failed to stream : %s \n", err.Error())
			return
		}
	}(formatter.provider)

loop:
	for {
		select {
		case msg := <-middleChan:
			sendChan <- formatter.parse(msg)
		case <-ctx.Done():
			break loop
		}
	}

	return nil
}

var rePreJSON = regexp.MustCompile(`^.*?\{`)

func (formatter *JSON) parse(msg types.Message) types.Message {
	msg.Processed = msg.RawLog

	msg.Processed = rePreJSON.ReplaceAllString(msg.Processed, "{")

	var payload map[string]interface{}

	if err := json.Unmarshal([]byte(msg.Processed), &payload); err != nil {
		msg.Message = msg.RawLog
		return msg
	}

	msg.Payload = payload

	if v, ok := payload[formatter.config.MessageField]; !ok {
		msg.Message = msg.Processed
		return msg
	} else {
		msg.Message = v.(string)
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
