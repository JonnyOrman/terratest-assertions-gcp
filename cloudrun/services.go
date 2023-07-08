package cloudrun

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/run/apiv2/runpb"
	"github.com/stretchr/testify/assert"
	run "google.golang.org/genproto/googleapis/cloud/run/v2"
)

func VerifyServiceExists(t *testing.T, projectId string, serviceName string) {
	ctx := context.Background()

	servicesClient, err := run.NewServicesClient(ctx)
	if err != nil {
		fmt.Errorf("Failed to create Services client: %w", err)
	}

	defer servicesClient.Close()

	serviceExists := true

	serviceRequest := &runpb.GetServiceRequest{
		Name: serviceName,
	}

	service, err := servicesClient.GetService(ctx, serviceRequest)
	if err != nil {
		fmt.Errorf("Failed to get Cloud Run service: %w", err)
		serviceExists = false
	}

	assert.True(t, serviceExists)
}
