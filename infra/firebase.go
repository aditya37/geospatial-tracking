package infra

import (
	"context"
	"sync"

	gofbs "firebase.google.com/go/v4"
	"github.com/aditya37/geofence-service/util"
	"google.golang.org/api/option"
)

var (
	firebaseInstance  *gofbs.App = nil
	firebaseSingleton sync.Once
	firebaseErr       error
)

type FirebaseConfig struct {
	StorageBucketName string
	ProjectId         string
	PathCredFile      string
}

func NewFirebaseClient(ctx context.Context, config FirebaseConfig) error {
	firebaseSingleton.Do(func() {
		client, err := gofbs.NewApp(
			ctx,
			&gofbs.Config{
				StorageBucket: config.StorageBucketName,
				ProjectID:     config.ProjectId,
			},
			option.WithCredentialsFile(config.PathCredFile),
			option.WithServiceAccountFile(config.PathCredFile),
		)
		if err != nil {
			util.Logger().Error(err)
			firebaseErr = err
		}
		firebaseInstance = client
	})
	if firebaseErr != nil {
		return firebaseErr
	}
	return nil
}

func GetFirebaseInstance() *gofbs.App {
	return firebaseInstance
}
