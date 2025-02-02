package agents

import (
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/tools"
)

const _defaultMaxIterations = 5

// AgentType is a string type representing the type of agent to create.
type AgentType string

const (
	// ZeroShotReactDescription is an AgentType constant that represents
	// the "zeroShotReactDescription" agent type.
	ZeroShotReactDescription AgentType = "zeroShotReactDescription"
)

// Initialize is a function that creates a new executor with the specified LLM
// model, tools, agent type, and options. It returns an Executor or an error
// if there is any issues during the creation process.
func Initialize(
	llm llms.LLM,
	tools []tools.Tool,
	agentType AgentType,
	opts ...CreationOption,
) (Executor, error) {
	options := executorDefaultOptions()
	for _, opt := range opts {
		opt(&options)
	}

	var agent Agent
	switch agentType {
	case ZeroShotReactDescription:
		agent = NewOneShotAgent(llm, tools, opts...)
	default:
		return Executor{}, ErrUnknownAgentType
	}

	return Executor{
		Agent:         agent,
		Tools:         tools,
		MaxIterations: options.maxIterations,
	}, nil
}
