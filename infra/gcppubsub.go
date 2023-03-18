package infra

import (
	"context"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/aditya37/logger"
	"google.golang.org/api/option"
)

var (
	gcppubsubInstance  *pubsub.Client = nil
	gcppubsubSingleton sync.Once
)

func NewGcpPubsubInstance(ctx context.Context, projectid string, opts ...option.ClientOption) {
	gcppubsubSingleton.Do(func() {
		client, err := pubsub.NewClient(ctx, projectid, opts...)
		if err != nil {
			logger.Error(err)
			return
		}
		gcppubsubInstance = client
	})
}

func GetGcpPubsubInstance() *pubsub.Client {
	return gcppubsubInstance
}
