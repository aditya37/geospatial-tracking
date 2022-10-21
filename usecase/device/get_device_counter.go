package device

import (
	"context"

	"github.com/aditya37/geospatial-tracking/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (du *DeviceUsecase) GetDeviceCounter(ctx context.Context, in *emptypb.Empty) (proto.ResponseGetDeviceCounter, error) {
	count, err := du.deviceManagerRepo.GetDeviceCounter(ctx)
	if err != nil {
		return proto.ResponseGetDeviceCounter{}, err
	}
	countDeviceDetect, err := du.deviceManagerRepo.GetCountDeviceDetect(ctx)
	if err != nil {
		return proto.ResponseGetDeviceCounter{}, err
	}

	return proto.ResponseGetDeviceCounter{
		// TODO: GET Active device
		ActivetedDevice:  count.ActivatedDevice,
		RecordedTracking: count.RecordedTracking,
		DetectDevice:     countDeviceDetect,
	}, nil
}
