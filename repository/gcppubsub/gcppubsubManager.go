package gcppubsub

import (
	"context"
	"errors"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/repository"
	getenv "github.com/aditya37/get-env"
	"github.com/aditya37/logger"
	"github.com/google/uuid"
)

type gcppubsubManager struct {
	client *pubsub.Client
}

func NewGcpPubsubManager(client *pubsub.Client) repository.Pubsub {
	return &gcppubsubManager{
		client: client,
	}
}

// topic..
func (gm *gcppubsubManager) createTopic(ctx context.Context, topic string) error {
	if _, err := gm.client.CreateTopic(ctx, topic); err != nil {
		return err
	}
	return nil
}

// get topic...
func (gm *gcppubsubManager) getTopic(ctx context.Context, topicname string) (*pubsub.Topic, error) {
	topic := gm.client.Topic(topicname)
	ok, err := topic.Exists(ctx)
	if err != nil {
		return nil, err
	}
	if !ok {
		if err := gm.createTopic(ctx, topicname); err != nil {
			return nil, err
		}
		return nil, errors.New("topic not found")
	}
	return topic, nil
}
func (gm *gcppubsubManager) createSubscription(ctx context.Context, servicename, topicname string) (*pubsub.Subscription, error) {
	topic, err := gm.getTopic(ctx, topicname)
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewUUID()
	if err == nil {
		servicename = servicename + "." + id.String()
	}

	return gm.client.CreateSubscription(
		ctx,
		servicename,
		pubsub.SubscriptionConfig{
			Topic:               topic,
			RetainAckedMessages: getenv.GetBool("PUBSUB_RETAIN_ACKEDMSG", false),
			RetentionDuration: time.Duration(
				getenv.GetInt("PUBSUB_RETENTION_DURATION", 15) * int(time.Minute),
			),
		},
	)
}
func (gm *gcppubsubManager) Subscribe(ctx context.Context, topic, servicename string, Callback func(ctx context.Context, m *pubsub.Message)) error {
	subs, err := gm.createSubscription(ctx, servicename, topic)
	if err != nil {
		return err
	}
	if err := subs.Receive(
		ctx,
		func(c context.Context, m *pubsub.Message) {
			Callback(c, m)
		},
	); err != nil {
		return err
	}
	return nil
}
func (gm *gcppubsubManager) Publish(ctx context.Context, param entity.PublishParam) error {
	topic, err := gm.getTopic(ctx, param.TopicName)
	if err != nil {
		logger.Error(err)
		return err
	}
	resp := topic.Publish(
		ctx,
		&pubsub.Message{
			Data: param.Message,
		},
	)
	if _, err := resp.Get(ctx); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
func (gm *gcppubsubManager) Close() error {
	return gm.client.Close()
}
