package repository

import (
	"context"
	"io"

	"cloud.google.com/go/pubsub"
	"github.com/aditya37/geospatial-tracking/entity"
)

type Pubsub interface {
	io.Closer
	Subscribe(ctx context.Context, topic, servicename string, Callback func(ctx context.Context, m *pubsub.Message)) error
	Publish(ctx context.Context, param entity.PublishParam) error
}
