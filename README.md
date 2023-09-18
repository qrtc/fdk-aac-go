[![PkgGoDev](https://pkg.go.dev/badge/github.com/qrtc/fdk-aac-go)](https://pkg.go.dev/github.com/qrtc/fdk-aac-go)

# fdk-aac-go

Go bindings for [fdk-aac](https://github.com/mstorsjo/fdk-aac). A standalone library of the Fraunhofer FDK AAC code from Android.

## Why fdk-aac-go

The purpose of fdk-aac-go is easing the adoption of fdk-aac codec library. Using Go, with just a few lines of code you can implement an application that encode/decode data easy.

##  Is this a new implementation of fdk-aac?

No! We are just exposing the great work done by the research organization of [Fraunhofer IIS](https://www.iis.fraunhofer.de/en/ff/amm/impl.html) as a golang library. All the functionality and implementation still resides in the official fdk-aac project.

# Features supported

- Decode AAC to PCM
- Encode PCM to AAC

# Usage

## Decode AAC frame to PCM

```go
package main

import (
	"fmt"

	fdkaac "github.com/qrtc/fdk-aac-go"
)

func main() {
	decoder, err := fdkaac.CreateAccDecoder(&fdkaac.AacDecoderConfig{
		TransportFmt: fdkaac.TtMp4Adts,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		decoder.Close()
	}()

	inBuf := []byte{
        // AAC frame
    }
	outBuf := make([]byte, 4096)

	n, err := decoder.DecodeFrame(inBuf, outBuf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outBuf[0:n])
}
```

## Enode PCM to AAC

```go
package main

import (
	"fmt"

	fdkaac "github.com/qrtc/fdk-aac-go"
)

func main() {
	encoder, err := fdkaac.CreateAccEncoder(&fdkaac.AacEncoderConfig{
		TransMux:    TtMp4Adts,
		AOT:         AotAacLc,
		SampleRate:  44100,
		MaxChannels: 2,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		encoder.Close()
	}()

	inBuf := []byte{
		// PCM bytes
	}
	outBuf := make([]byte, 4096)

	n, err := encoder.Encode(inBuf, outBuf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(outBuf[0:n])
}

```

# Dependencies

* fdk-aac
