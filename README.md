# pure
This is real fork of [Zyko0's opencl](https://github.com/Zyko0/go-opencl), big thank. My modification call [pureCL(simplyfy Zyko0's opencl)](https://github.com/opencl-pure/pureCL) as button layer. <br>
This package provide higher level wrapper to OpenCL,
that means it provide GO error handling, handle OpenCL errors to GO errors, 
but still it is near 1:1 wrapper
# goal
- low level wrap of OpenCL 
-  GO error handling
  
-  try to have all functions of OpenCL (so if you have some, give PR)
-  easy to multiplatform (thank [purego](https://github.com/ebitengine/purego))
-  easy find path (custumize path to openclLib shared library)
-  easy to compile, we do not need cgo and not need knowing link to shared library
-  try [purego](https://github.com/ebitengine/purego) and bring opencl on android without complicate link
# not goal
- be faster as cgo version, [purego](https://github.com/ebitengine/purego) is using same mechanism as cgo 
# examples

# example

```go
package main

import (
	"fmt"
	constants "github.com/opencl-pure/constantsCL"
	middle "github.com/opencl-pure/middleCL"
	pure "github.com/opencl-pure/pureCL"
	"log"
)

const (
	dataSize = 32
)

var (
	code = `
        __kernel void set_i(__global float* out){
		size_t i = get_global_id(0);
		out[i] = i;
	}`
)

func main() {
	err := middle.Init(pure.Version2_0)
	if err != nil {
		log.Fatal("err:", err)
	}

	platforms, err := middle.GetPlatforms()
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("platforms", len(platforms))

	name, err := platforms[0].GetName()
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("name:", name)

	version, err := platforms[0].GetVersion()
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("version:", version)

	platformExtensions, err := platforms[0].GetExtensions()
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("platform extensions:", platformExtensions)

	devices, err := platforms[0].GetDevices(constants.CL_DEVICE_TYPE_ALL)
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("devices:", len(devices))

	deviceExtensions, err := devices[0].GetExtensions()
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("device extensions:", deviceExtensions)

	ctx, err := devices[0].CreateContext(nil)
	if err != nil {
		log.Fatal("err:", err)
	}
	defer ctx.Release()
	fmt.Println("context:", ctx)

	queue, err := ctx.CreateCommandQueue(&devices[0])
	if err != nil {
		log.Fatal("err:", err)
	}
	defer queue.Release()
	fmt.Println("queue:", queue)

	program, err := ctx.CreateProgram(fmt.Sprint(code))
	if err != nil {
		log.Fatal("err:", err)
	}
	defer program.Release()
	fmt.Println("program:", program)

	logs, err := program.Build(&devices[0], nil)
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("logs:", logs)

	kernel, err := program.CreateKernel("set_i")
	if err != nil {
		log.Fatal("err:", err)
	}
	defer kernel.Release()
	fmt.Println("kernel:", kernel)

	data := make([]float32, dataSize)
	bufferData := middle.GetBufferData(data)
	fmt.Println("buffer data:", bufferData.DataSize, bufferData.Pointer)
	buffer, err := ctx.CreateBuffer(
		[]pure.MemFlag{
			constants.CL_MEM_WRITE_ONLY,
			constants.CL_MEM_ALLOC_HOST_PTR,
		},
		uint(bufferData.DataSize),
	)
	if err != nil {
		log.Fatal("err:", err)
	}
	defer buffer.Release()
	fmt.Println("buffer:", buffer)

	size, err := buffer.Size()
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("buffer size:", size)
	arg, _ := middle.NewKernelArg(buffer)
	err = kernel.SetArg(0, arg)
	if err != nil {
		log.Fatal("err:", err)
	}

	err = queue.EnqueueNDRangeKernel(kernel, 1, nil, []uint64{dataSize}, nil)
	if err != nil {
		log.Fatal("err:", err)
	}

	_ = queue.Flush()
	_ = queue.Finish()

	err = queue.EnqueueReadBuffer(buffer, true, middle.GetBufferData(data))
	if err != nil {
		log.Fatal("err:", err)
	}
	fmt.Println("data:", data)
}
```
