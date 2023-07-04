package pubsub

import (
	"context"
	"fmt"
	"testing"

	"cloud.google.com/go/pubsub"
	"github.com/stretchr/testify/assert"
)

func VerifyTopicExists(t *testing.T, projectId string, topicName string) {
	ctx := context.Background()

	pubsubClient, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		fmt.Errorf("Failed to create PubSub client: %w", err)
	}

	topic := pubsubClient.Topic(topicName)

	topicExists, err := topic.Exists(ctx)
	if err != nil {
		fmt.Errorf("Failed to check if PubSub topic exists: %w", err)
	}

	assert.True(t, topicExists)
}
