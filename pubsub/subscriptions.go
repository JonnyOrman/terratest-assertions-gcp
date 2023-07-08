package pubsub

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/pubsub"
	"github.com/stretchr/testify/assert"
)

func VerifySubscriptionExists(t *testing.T, projectId string, subscriptionName string) {
	ctx := context.Background()

	pubsubClient, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		fmt.Errorf("Failed to create PubSub client: %w", err)
	}

	defer pubsubClient.Close()

	subscription := pubsubClient.Subscription(subscriptionName)

	subscriptionExists, err := subscription.Exists(ctx)
	if err != nil {
		fmt.Errorf("Failed to check if PubSub subscription exists: %w", err)
	}

	assert.True(t, subscriptionExists)
}
