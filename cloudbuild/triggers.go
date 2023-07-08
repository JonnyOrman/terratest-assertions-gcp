package cloudbuild

import (
	"context"
	"fmt"
	"testing"

	cloudbuild "cloud.google.com/go/cloudbuild/apiv1"
	"cloud.google.com/go/cloudbuild/apiv1/v2/cloudbuildpb"
	"github.com/stretchr/testify/assert"
)

func VerifyTriggerExists(t *testing.T, projectId string, triggerId string) {
	ctx := context.Background()

	cloudbuildClient, err := cloudbuild.NewClient(ctx)
	if err != nil {
		fmt.Errorf("Failed to create Cloud Build client: %w", err)
	}

	defer cloudbuildClient.Close()

	triggerExists := true

	buildTriggerRequest := cloudbuildpb.GetBuildTriggerRequest{
		TriggerId: triggerId,
		ProjectId: projectId,
	}

	buildTrigger, err := cloudbuildClient.GetBuildTrigger(ctx, &buildTriggerRequest)
	if err != nil {
		fmt.Errorf("Failed to get Cloud Build trigger: %w", err)
		triggerExists = false
	}

	assert.True(t, triggerExists)
}
