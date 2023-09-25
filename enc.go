package fdkaac

/*
#cgo pkg-config: fdk-aac
#include <fdk-aac/aacenc_lib.h>

AACENC_ERROR aacEncEncodeWrapped(const HANDLE_AACENCODER hAacEncoder,
			void* in, int inLen, int sampleBitDepth,
			void* out, int outLen, int* numOutBytes) {
	AACENC_ERROR err;
	AACENC_BufDesc inBuf = { 0 }, outBuf = { 0 };
	AACENC_InArgs inArgs = { 0 };
	AACENC_OutArgs outArgs = { 0 };
	int inIdentifier = IN_AUDIO_DATA;
	int inElemSize = sampleBitDepth / 8;
	int outIdentifier = OUT_BITSTREAM_DATA;
	int outElemSize = 1;

	inArgs.numInSamples = in ? inLen / inElemSize : -1;
	inBuf.numBufs = 1;
	inBuf.bufs = &in;
	inBuf.bufferIdentifiers = &inIdentifier;
	inBuf.bufSizes = &inLen;
	inBuf.bufElSizes = &inElemSize;
	outBuf.numBufs = 1;
	outBuf.bufs = &out;
	outBuf.bufferIdentifiers = &outIdentifier;
	outBuf.bufSizes = &outLen;
	outBuf.bufElSizes = &outElemSize;

	err = aacEncEncode(hAacEncoder, &inBuf, &outBuf, &inArgs, &outArgs);
	*numOutBytes = outArgs.numOutBytes;
	return err;
}
*/
import "C"
import (
	"errors"
	"unsafe"
)

var encErrors = [...]error{
	C.AACENC_OK:                    nil,
	C.AACENC_INVALID_HANDLE:        errors.New("handle passed to function call was invalid"),
	C.AACENC_MEMORY_ERROR:          errors.New("memory allocation failed"),
	C.AACENC_UNSUPPORTED_PARAMETER: errors.New("parameter not available"),
	C.AACENC_INVALID_CONFIG:        errors.New("configuration not provided"),
	C.AACENC_INIT_ERROR:            errors.New("general initialization error"),
	C.AACENC_INIT_AAC_ERROR:        errors.New("AAC library initialization error"),
	C.AACENC_INIT_SBR_ERROR:        errors.New("SBR library initialization error"),
	C.AACENC_INIT_TP_ERROR:         errors.New("transport library initialization error"),
	C.AACENC_INIT_META_ERROR:       errors.New("meta data library initialization error"),
	C.AACENC_INIT_MPS_ERROR:        errors.New("MPS library initialization error"),
	C.AACENC_ENCODE_ERROR:          errors.New("the encoding process was interrupted by an unexpected error"),
	C.AACENC_ENCODE_EOF:            errors.New("end of file reached"),
}

// Encoder End Of File.
var EncEOF = encErrors[C.AACENC_ENCODE_EOF]

// Bitrate Mode
type BitrateMode int

const (
	BitrateModeConstant BitrateMode = iota
	BitrateModeVeryLow
	BitrateModeLow
	BitrateModeMedium
	BitrateModeHigh
	BitrateModeVeryHigh
)

// Signaling Mode
type SignalingMode int

const (
	SignalingModeImplicitCompatible SignalingMode = iota
	SignalingModeExplicitCompatible
	SignalingModeExplicitHierarchical
)

// Meta Data Mode
type MetaDataMode int

const (
	MetaDataModeNone MetaDataMode = iota
	MetaDataModeDynamicRangeInfoOnly
	MetaDataModeDynamicRangeInfoAndAncillaryData
	MetaDataModeNoneAncillaryDataOnly
)

// AAC Encoder Config
type AacEncoderConfig struct {
	// Number of channels to be allocated.
	MaxChannels int
	// Sample bitdepth.
	SampleBitDepth int
	// Audio object type.
	AOT AudioObjectType
	// Total encoder bitrate.
	Bitrate int
	// Bitrate mode.
	BitrateMode BitrateMode
	// Audio input data sampling rate.
	SampleRate int
	// Configure SBR independently of the chosen Audio Object Type.
	SbrMode SbrMode
	// Core encoder (AAC) audio frame length in samples.
	GranuleLength int
	// Set explicit channel mode. Channel mode must match with number of input channels.
	ChannelMode ChannelMode
	// Input audio data channel ordering scheme.
	ChannelOrder ChannelOrder
	// Controls activation of downsampled SBR.
	SbrRatio int
	// Controls the use of the afterburner feature.
	IsAfterBurner bool
	// Core encoder audio bandwidth.
	Bandwith int
	// Peak bitrate configuration parameter to adjust maximum bits per audio frame.
	PeakBitrate int
	// Transport type to be used.
	TransMux TransportType
	// Frame count period for sending in-band configuration buffers within LATM/LOAS transport layer.
	HeaderPeriod int
	// Signaling mode of the extension AOT.
	SignalingMode SignalingMode
	// Number of sub frames in a transport frame for LOAS/LATM or ADTS (default 1).
	TransportSubFrames int
	// AudioMuxVersion to be used for LATM.
	AudioMuxVersion int
	// Configure protection in transport layer.
	IsProtection bool
	// Constant ancillary data bitrate in bits/second.
	AncillaryBitrate int
	// Configure Meta Data.
	MetaDataMode MetaDataMode
}

// EncInfo provides some info about the encoder configuration.
type EncInfo struct {
	// Maximum number of encoder bitstream bytes within one frame.
	// Size depends on maximum number of supported channels in encoder instance.
	MaxOutBufBytes uint
	// Maximum number of ancillary data bytes which can be
	// inserted into bitstream within one frame.
	MaxAncBytes uint
	// Internal input buffer fill level in samples per channel.
	InBufFillLevel uint
	// Number of input channels expected in encoding process.
	InputChannels uint
	// Amount of input audio samples consumed each frame per channel,
	// depending on audio object type configuration.
	FrameLength uint
	// Codec delay in PCM samples/channel.
	NDelay uint
	// Codec delay in PCM samples/channel.
	NDelayCore uint
	// Configuration buffer in binary format as an AudioSpecificConfig or
	// StreamMuxConfig according to the selected transport type.
	ConfBuf []byte
	// Number of valid bytes in confBuf.
	ConfSize uint
}

// AAC Encoder
type AacEncoder struct {
	// private handler
	ph C.HANDLE_AACENCODER
	// config
	AacEncoderConfig
	// info
	info *EncInfo
}

// Encode
func (enc *AacEncoder) Encode(in, out []byte) (n int, err error) {
	var inPtr, outPtr unsafe.Pointer
	if len(in) > 0 {
		inPtr = unsafe.Pointer(&in[0])
	}
	if len(out) > 0 {
		outPtr = unsafe.Pointer(&out[0])
	}
	errNo := C.aacEncEncodeWrapped(enc.ph,
		inPtr, C.int(len(in)), C.int(enc.SampleBitDepth),
		outPtr, C.int(len(out)), (*C.int)(unsafe.Pointer(&n)))

	return n, encErrors[errNo]
}

// Flush
func (enc *AacEncoder) Flush(out []byte) (n int, err error) {
	validBytes := 0
	for {
		validBytes, err = enc.Encode(nil, out[n:])
		n += validBytes
		if err == EncEOF {
			return n, nil
		} else if err != nil {
			return n, err
		}
	}
}

func (enc *AacEncoder) GetInfo() (*EncInfo, error) {
	info := C.AACENC_InfoStruct{}
	if errNo := C.aacEncInfo(enc.ph, &info); errNo != C.AACENC_OK {
		return nil, encErrors[errNo]
	}

	return &EncInfo{
		MaxOutBufBytes: uint(info.maxOutBufBytes),
		MaxAncBytes:    uint(info.maxAncBytes),
		InBufFillLevel: uint(info.inBufFillLevel),
		InputChannels:  uint(info.inputChannels),
		FrameLength:    uint(info.frameLength),
		NDelay:         uint(info.nDelay),
		NDelayCore:     uint(info.nDelay),
		ConfBuf:        C.GoBytes(unsafe.Pointer(&info.confBuf[0]), C.int(info.confSize)),
		ConfSize:       uint(info.confSize),
	}, nil
}

// Close
func (enc *AacEncoder) Close() error {
	return encErrors[C.aacEncClose(&enc.ph)]
}

// Create AAC Encoder
func CreateAccEncoder(config *AacEncoderConfig) (enc *AacEncoder, err error) {
	var errNo C.AACENC_ERROR = C.AACENC_OK
	enc = &AacEncoder{
		AacEncoderConfig: *populateEncConfig(config),
	}

	if errNo = C.aacEncOpen(&enc.ph, 0, C.uint(enc.MaxChannels)); errNo != C.AACENC_OK {
		return nil, encErrors[errNo]
	}

	defer func() {
		if errNo != C.AACENC_OK {
			C.aacEncClose(&enc.ph)
		}
	}()

	if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_AOT,
		C.uint(enc.AOT)); errNo != C.AACENC_OK {
		return nil, encErrors[errNo]
	}
	if enc.BitrateMode != BitrateModeConstant {
		if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_BITRATEMODE,
			C.uint(enc.BitrateMode)); errNo != C.AACENC_OK {
			return nil, encErrors[errNo]
		}
	}
	if enc.BitrateMode == BitrateModeConstant {
		if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_BITRATE,
			C.uint(enc.Bitrate)); errNo != C.AACENC_OK {
			return nil, encErrors[errNo]
		}
	}
	if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_SAMPLERATE,
		C.uint(enc.SampleRate)); errNo != C.AACENC_OK {
		return nil, encErrors[errNo]
	}
	if enc.SbrMode != SbrModeDefault {
		if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_SBR_MODE,
			C.uint(enc.SbrMode-1)); errNo != C.AACENC_OK {
			return nil, encErrors[errNo]
		}
	}
	if enc.GranuleLength != 0 {
		if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_GRANULE_LENGTH,
			C.uint(enc.GranuleLength)); errNo != C.AACENC_OK {
			return nil, encErrors[errNo]
		}
	}
	if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_CHANNELMODE,
		C.uint(enc.ChannelMode)); errNo != C.AACENC_OK {
		return nil, encErrors[errNo]
	}
	if enc.ChannelOrder != ChannelOrderMpeg {
		if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_CHANNELORDER,
			C.uint(enc.ChannelOrder)); errNo != C.AACENC_OK {
			return nil, encErrors[errNo]
		}
	}
	if enc.SbrRatio != 0 {
		if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_SBR_RATIO,
			C.uint(enc.SbrRatio)); errNo != C.AACENC_OK {
			return nil, encErrors[errNo]
		}
	}
	if enc.IsAfterBurner {
		if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_AFTERBURNER,
			C.uint(1)); errNo != C.AACENC_OK {
			return nil, encErrors[errNo]
		}
	}
	if enc.Bandwith > 0 {
		if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_BANDWIDTH,
			C.uint(enc.Bandwith)); errNo != C.AACENC_OK {
			return nil, encErrors[errNo]
		}
	}
	if enc.PeakBitrate > 0 {
		if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_PEAK_BITRATE,
			C.uint(enc.PeakBitrate)); errNo != C.AACENC_OK {
			return nil, encErrors[errNo]
		}
	}
	if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_TRANSMUX,
		C.uint(enc.TransMux)); errNo != C.AACENC_OK {
		return nil, encErrors[errNo]
	}
	if enc.HeaderPeriod > 0 {
		if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_HEADER_PERIOD,
			C.uint(enc.HeaderPeriod)); errNo != C.AACENC_OK {
			return nil, encErrors[errNo]
		}
	}
	if enc.SignalingMode != SignalingModeImplicitCompatible {
		if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_SIGNALING_MODE,
			C.uint(enc.SignalingMode)); errNo != C.AACENC_OK {
			return nil, encErrors[errNo]
		}
	}
	if enc.TransportSubFrames > 0 {
		if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_TPSUBFRAMES,
			C.uint(enc.TransportSubFrames)); errNo != C.AACENC_OK {
			return nil, encErrors[errNo]
		}
	}
	if enc.AudioMuxVersion > 0 {
		if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_AUDIOMUXVER,
			C.uint(enc.AudioMuxVersion)); errNo != C.AACENC_OK {
			return nil, encErrors[errNo]
		}
	}
	if enc.IsProtection {
		if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_PROTECTION,
			C.uint(1)); errNo != C.AACENC_OK {
			return nil, encErrors[errNo]
		}
	}
	if enc.AncillaryBitrate > 0 {
		if errNo = C.aacEncoder_SetParam(enc.ph, C.AACENC_ANCILLARY_BITRATE,
			C.uint(enc.AncillaryBitrate)); errNo != C.AACENC_OK {
			return nil, encErrors[errNo]
		}
	}
	if enc.MetaDataMode != MetaDataModeNone {
		panic("TODO. support metadata mode")
	}

	if errNo = C.aacEncEncode(enc.ph, nil, nil, nil, nil); errNo != C.AACENC_OK {
		return nil, encErrors[errNo]
	}

	if enc.info, err = enc.GetInfo(); err != nil {
		return nil, err
	}

	return enc, encErrors[errNo]
}

func populateEncConfig(c *AacEncoderConfig) *AacEncoderConfig {
	if c == nil {
		c = &AacEncoderConfig{}
	}
	if c.MaxChannels == 0 {
		c.MaxChannels = defaultMaxChannels
	}
	if c.SampleBitDepth == 0 {
		c.SampleBitDepth = defaultSampleBitdepth
	}
	if c.AOT == 0 {
		c.AOT = defaultAOT
	}
	if c.ChannelMode == 0 {
		if c.MaxChannels <= 7 {
			c.ChannelMode = ChannelMode(c.MaxChannels)
		}
	}
	if c.SampleRate == 0 {
		c.SampleRate = defaultSamplerate
	}
	if c.Bitrate == 0 {
		c.Bitrate = defaultBitrate
	}

	return c
}
