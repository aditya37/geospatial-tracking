package device

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/aditya37/geospatial-tracking/entity"
	"github.com/aditya37/geospatial-tracking/proto"
	getenv "github.com/aditya37/get-env"
	"github.com/aditya37/logger"
	rgbqrcode "github.com/aditya37/rgb-qrcode"
)

func (du *DeviceUsecase) DeviceQrCode(ctx context.Context, request *proto.RequestDeviceQrCode) (proto.ResponseDeviceQrCode, error) {
	// validate event type...
	if isValidEventType := du.isValidEventType(request); !isValidEventType {
		return proto.ResponseDeviceQrCode{}, errors.New("unknown event type")
	}

	// validate action type....
	if isValidAction := du.isValidActionType(request); !isValidAction {
		return proto.ResponseDeviceQrCode{}, errors.New("unknown action type")
	}

	// generate device QR with event pairing user device...
	if request.ActionType == proto.ActionType_GENERATE_DEVICE_QR_CODE && request.EventType == proto.EventType_PAIRING_USER_TO_DEVICE {
		resp, err := du.processGenerateQrCodePairingUserDevice(ctx, request)
		if err != nil {
			return proto.ResponseDeviceQrCode{}, err
		}
		return resp, nil
	} else if request.ActionType == proto.ActionType_GET_DEVICE_QR_CODE && request.EventType == proto.EventType_PAIRING_USER_TO_DEVICE {
		// Get device qrcode with event type pairing device...
		resp, err := du.processGetQrCodePairUserDevice(ctx, request)
		if err != nil {
			return proto.ResponseDeviceQrCode{}, err
		}
		return resp, nil
	} else {
		return proto.ResponseDeviceQrCode{}, errors.New("unknown action")
	}
}

// process generate qr code for pairing device....
func (du *DeviceUsecase) processGenerateQrCodePairingUserDevice(ctx context.Context, request *proto.RequestDeviceQrCode) (proto.ResponseDeviceQrCode, error) {
	// validate device id...
	if _, err := du.deviceManagerRepo.GetDeviceByDeviceId(
		ctx,
		request.DeviceId,
	); err != nil {
		logger.Error(err)
		return proto.ResponseDeviceQrCode{}, err
	}

	// logo path
	logo, _ := os.Open(
		getenv.GetString("PATH_LOGO_QR_CODE_PAIR_DEVICE", "./logo/pair_device_logo.png"),
	)
	defer logo.Close()

	// payload...
	payload, _ := json.Marshal(request)

	// generate qrcode
	qrResult, err := du.generateQrCode(
		string(payload),
		logo,
	)
	if err != nil {
		logger.Error(err)
		return proto.ResponseDeviceQrCode{}, err
	}

	// create output qr code to image
	filename := fmt.Sprintf("qr.%s.%s.png", request.DeviceId, request.EventType)
	// path for temp_qr
	pathTempQr := fmt.Sprintf("./qr_code/%s", filename)

	if err := ioutil.WriteFile(
		pathTempQr,
		qrResult.PNG.Bytes(),
		os.FileMode(0644),
	); err != nil {
		logger.Error(err)
		return proto.ResponseDeviceQrCode{}, err
	}

	// open file before upload to firebase...
	fs, _ := os.Open(pathTempQr)
	defer fs.Close()

	// publish to firebase...
	if err := du.fbsStorage.UploadToStorageBucket(ctx, filename, fs); err != nil {
		logger.Error(err)
		return proto.ResponseDeviceQrCode{}, err
	}
	// storage url..
	url, err := du.fbsStorage.GetFileDownloadUrl(ctx, filename)
	if err != nil {
		logger.Error(err)
		return proto.ResponseDeviceQrCode{}, err
	}
	// store to database...
	if err := du.deviceManagerRepo.InsertDeviceQr(
		ctx,
		entity.QRDevice{
			EventType:   int(request.EventType),
			DeviceId:    request.DeviceId,
			Description: request.Description,
			QrFile:      filename,
			Url:         url,
		},
	); err != nil {
		logger.Error(err)
		return proto.ResponseDeviceQrCode{}, err
	}
	defer os.Remove(pathTempQr)

	return proto.ResponseDeviceQrCode{
		QrFile:    filename,
		EventType: request.EventType.String(),
		Url:       url,
	}, nil
}

// processGetQrCodePairUserDevice...
func (du *DeviceUsecase) processGetQrCodePairUserDevice(ctx context.Context, request *proto.RequestDeviceQrCode) (proto.ResponseDeviceQrCode, error) {
	// validate device id...
	if _, err := du.deviceManagerRepo.GetDeviceByDeviceId(
		ctx,
		request.DeviceId,
	); err != nil {
		logger.Error(err)
		return proto.ResponseDeviceQrCode{}, err
	}
	var response proto.ResponseDeviceQrCode
	// redis key...
	key := fmt.Sprintf(
		"cache.get.device.qr.%s.user.%s",
		request.EventType.String(),
		request.DeviceId,
	)
	cache, err := du.cacheManager.Get(key)
	if err != nil {
		qrfile, err := du.deviceManagerRepo.GetDeviceQrCode(
			ctx,
			entity.QRDevice{
				EventType: int(request.EventType),
				DeviceId:  request.DeviceId,
			},
		)
		if err != nil {
			logger.Error(err)
			return proto.ResponseDeviceQrCode{}, err
		}
		resp := proto.ResponseDeviceQrCode{
			QrFile:    qrfile.QrFile,
			Url:       qrfile.Url,
			EventType: request.EventType.String(),
		}
		buf, _ := json.Marshal(resp)
		ttl := 86400 * time.Second
		du.cacheManager.Set(key, buf, ttl) // 1 day
		return resp, nil
	}
	if err := json.Unmarshal([]byte(cache), &response); err != nil {
		logger.Error()
		return proto.ResponseDeviceQrCode{}, err
	}
	return response, nil
}

// generate QR Code...
func (du *DeviceUsecase) generateQrCode(value string, logo *os.File) (rgbqrcode.ResultEncode, error) {
	qr, err := rgbqrcode.New(
		rgbqrcode.GenerateParam{
			LogoPath: logo,
			QrValue:  value,
			QrSize:   256,
		},
	)
	if err != nil {
		logger.Error(err)
		return rgbqrcode.ResultEncode{}, err
	}
	return qr.Encode()
}

// isValidActionType....
func (du *DeviceUsecase) isValidEventType(request *proto.RequestDeviceQrCode) bool {
	if request.EventType < 0 || int(request.EventType) >= len(proto.EventType_name) {
		return false
	}
	return true
}

func (du *DeviceUsecase) isValidActionType(request *proto.RequestDeviceQrCode) bool {
	if request.ActionType < 0 || int(request.ActionType) >= len(proto.ActionType_name) {
		return false
	}
	return true
}
