package middleCL

import (
	constants "github.com/opencl-pure/constantsCL"
	pure "github.com/opencl-pure/pureCL"
	"strings"
)

type Device struct {
	D pure.Device
}

func (d *Device) getInfo(name pure.DeviceInfo) (string, error) {
	size := pure.Size(0)
	st := pure.GetDeviceInfo(d.D, name, pure.Size(0), nil, &size)
	if st != constants.CL_SUCCESS {
		return "", pure.StatusToErr(st)
	}
	info := make([]byte, size)
	st = pure.GetDeviceInfo(d.D, name, size, info, nil)
	if st != constants.CL_SUCCESS {
		return "", pure.StatusToErr(st)
	}
	return string(info), nil
}

func (d *Device) GetExtensions() ([]pure.Extension, error) {
	extensions, err := d.getInfo(constants.CL_DEVICE_EXTENSIONS)
	if err != nil {
		return nil, err
	}
	return strings.Split(extensions, " "), nil
}
