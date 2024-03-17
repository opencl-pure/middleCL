package middleCL

import (
	pure "github.com/opencl-pure/pureCL"
	"image"
)

func GetBufferData[T pure.BufferType](data []T) *pure.BufferData {
	return pure.GetBufferData(data)
}

func GetImageBufferData(img image.RGBA) *pure.ImageData {
	return pure.GetImageBufferData(img)
}

func Init(v pure.Version, paths ...string) error {
	return pure.Init(v, paths...)
}
