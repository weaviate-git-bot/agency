package agency

import (
	"context"
	"fmt"
)

type OperationHandler func(context.Context, Message, *OperationConfig) (Message, error)

// Operation is basic building block.
type Operation struct {
	handler OperationHandler
	config  *OperationConfig
}

func (p *Operation) Config() *OperationConfig {
	return p.config
}

// OperationConfig represents abstract operation configuration.
// It contains fields for all possible modalities but nothing specific to concrete model implementations.
type OperationConfig struct {
	Prompt   string
	Messages []Message
}

// NewOperation allows to create an operation from a function.
func NewOperation(handler OperationHandler) *Operation {
	return &Operation{
		handler: handler,
		config:  &OperationConfig{},
	}
}

// Execute executes operation handler with input message and current configuration.
func (p *Operation) Execute(ctx context.Context, input Message) (Message, error) {
	output, err := p.handler(ctx, input, p.config)
	if err != nil {
		return Message{}, err
	}
	return output, nil
}

func (p *Operation) SetPrompt(prompt string, args ...any) *Operation {
	p.config.Prompt = fmt.Sprintf(prompt, args...)
	return p
}

func (p *Operation) SetMessages(msgs []Message) *Operation {
	p.config.Messages = msgs
	return p
}