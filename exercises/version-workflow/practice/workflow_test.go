package loanprocess

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.temporal.io/sdk/worker"
)

func TestReplayWorkflowHistoryFromFile(t *testing.T) {

	// Create the Workflow Replayer
	replayer := worker.NewWorkflowReplayer()

	// Register the Workflow Definition with the Replayer
	replayer.RegisterWorkflow(LoanProcessingWorkflow)

	// Replay the Event History from the JSON file you downloaded
	err := replayer.ReplayWorkflowHistoryFromJSONFile(nil, "history_for_original_execution.json")

	// Use assert.NoError to verify that replaying the history
	//      does not return an error
	assert.NoError(t, err)
}
