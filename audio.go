package fdkaac

// The FileFormat is file format.
type FileFormat int

const (
	// Unknown format
	FfUnknown FileFormat = -1
	// No container, bit stream data conveyed "as is"
	FfRaw FileFormat = 0
	// 3GPP file format
	Ff3gpp FileFormat = 3
	// MPEG-4 File format
	FfMp4f FileFormat = 4
	// Proprietary raw packet file
	FfRawPackets FileFormat = 5
)

// The TransportType is transport type.
type TransportType int

const (
	// Unknown format
	TtUnknown TransportType = -1
	// As is access units (packet based since there is obviously no sync layer)
	TtMp4Raw TransportType = 0
	// ADIF bitstream format
	TtMp4Adif TransportType = 1
	// ADTS bitstream format
	TtMp4Adts TransportType = 2
	// Audio Mux Elements with muxConfigPresent = 1
	TtMp4LatmMcp1 TransportType = 6
	// Audio Mux Elements with muxConfigPresent = 0, out of band StreamMuxConfig
	TtMp4LatmMcp0 TransportType = 7
	// Audio Sync Stream
	TtMp4Loas TransportType = 10
	// Digital Radio Mondial (DRM30/DRM+) bitstream format
	TtDrm TransportType = 12
)

// TtIsPacket check transport type is support packet.
func TtIsPacket(tt TransportType) bool {
	return tt == TtMp4Raw || tt == TtDrm || tt == TtMp4LatmMcp0 || tt == TtMp4LatmMcp1
}

// The AudioObjectType is audio object type
type AudioObjectType int

const (
	// None
	AotNone AudioObjectType = -1
	// Null Object
	AotNullObject AudioObjectType = 0
	// Main profile
	AotAacMain AudioObjectType = 1
	// Low Complexity object
	AotAacLc AudioObjectType = 2
	// Scalable Sampling Rate
	AotAacSsr AudioObjectType = 3
	// Long Term Prediction
	AotAacLtp AudioObjectType = 4
	// Spectral Band Replication
	AotSbr AudioObjectType = 5
	// Scalable
	AotAacScal AudioObjectType = 6
	// TwinVQ
	AotTwinVq AudioObjectType = 7
	// Code-Excited Linear Prediction
	AotCelp AudioObjectType = 8
	// Harmonic Vector Excitation Coding
	AotHvxc AudioObjectType = 9
	// Reserved
	AotRsvd10 AudioObjectType = 10
	// Reserved
	AotRsvd11 AudioObjectType = 11
	// TTSI Object
	AotTtsi AudioObjectType = 12
	// Main Synthetic object
	AotMainSynth AudioObjectType = 13
	// Wavetable Synthesis object
	AotWavTabSynth AudioObjectType = 14
	// General MIDI object
	AotGenMidi AudioObjectType = 15
	// Algorithmic Synthesis and Audio FX object
	AotAlgSynthAudFx AudioObjectType = 16
	// Error Resilient(ER) AAC Low Complexity
	AotErAacLc AudioObjectType = 17
	// Reserved
	AotRsvd18 AudioObjectType = 18
	// Error Resilient(ER) AAC LTP object
	AotErAacLtp AudioObjectType = 19
	// Error Resilient(ER) AAC Scalable object
	AotErAacScal AudioObjectType = 20
	// Error Resilient(ER) TwinVQ object
	AotErTwinVq AudioObjectType = 21
	// Error Resilient(ER) BSAC object
	AotErBsac AudioObjectType = 22
	// Error Resilient(ER) AAC LowDelay object
	AotErAacLd AudioObjectType = 23
	// Error Resilient(ER) CELP object
	AotErCelp AudioObjectType = 24
	// Error Resilient(ER) HVXC object
	AotErHvxc AudioObjectType = 25
	// Error Resilient(ER) HILN object
	AotErHiln AudioObjectType = 26
	// Error Resilient(ER) Parametric object
	AotErPara AudioObjectType = 27
	// Might become SSC
	AotRsvd28 AudioObjectType = 28
	// PS, Parametric Stereo (includes SBR)
	AotPs AudioObjectType = 29
	// MPEG Surround
	AotMpegs AudioObjectType = 30

	// Signal AOT uses more than 5 bits
	AotEscape AudioObjectType = 31

	// MPEG-Layer1 in mp4
	AotMp3OnMp4L1 AudioObjectType = 32
	// MPEG-Layer2 in mp4
	AotMp3OnMp4L2 AudioObjectType = 33
	// MPEG-Layer3 in mp4
	AotMp3OnMp4L3 AudioObjectType = 34
	// Might become DST (Direct Stream Transfer)
	AotRsvd35 AudioObjectType = 35
	// Might become ALS (Audio Lossless Coding)
	AotRsvd36 AudioObjectType = 36
	// AAC + SLS
	AotAacSls AudioObjectType = 37
	// Scalable To Lossless
	AotSls AudioObjectType = 38
	// AAC Enhanced Low Delay
	AotErAacEld AudioObjectType = 39
	// Unified Speech and Audio Coding
	AotUsac AudioObjectType = 42
	// Spatial Audio Object Coding
	AotSaoc AudioObjectType = 43
	// Low Delay MPEG Surround
	AotLdMpegs AudioObjectType = 44

	// Pseudo AOTs
	// Virtual AOT MP2 Low Complexity profile
	AotMp2AacLc AudioObjectType = 129
	// Virtual AOT MP2 Low Complexity Profile with SBR
	AotMp3Sbr AudioObjectType = 132
	// Virtual AOT for DRM (ER-AAC-SCAL without SBR)
	AotDrmAac AudioObjectType = 143
	// Virtual AOT for DRM (ER-AAC-SCAL with SBR)
	AotDrmSbr AudioObjectType = 144
	// Virtual AOT for DRM (ER-AAC-SCAL with SBR and MPEG-PS)
	AotDrmMpegPs AudioObjectType = 145
	// Virtual AOT for DRM Surround (ER-AAC-SCAL (+SBR) +MPS)
	AotDrmSurround AudioObjectType = 146
	// Virtual AOT for DRM with USAC
	AotDrmUsac AudioObjectType = 147
)

// AotCanDoPs check a audio object type is support parameter stereo.
func AotCanDoPs(aot AudioObjectType) bool {
	return aot == AotAacLc || aot == AotSbr ||
		aot == AotPs || aot == AotErBsac || aot == AotDrmAac
}

// IsUSAC check a audio object type is USAC
func IsUSAC(aot AudioObjectType) bool {
	return aot == AotUsac
}

// IsLowDelay check a audio object type is low delay
func IsLowDelay(aot AudioObjectType) bool {
	return aot == AotErAacLd || aot == AotErAacEld
}

// Channel Mode ( 1-7 equals MPEG channel configurations, others are arbitrary).
type ChannelMode int

const (
	ModeInvalid ChannelMode = -1
	ModeUnknown ChannelMode = 0
	// C
	Mode_1 ChannelMode = 1
	// L+R
	Mode_2 ChannelMode = 2
	// C, L+R
	Mode_1_2 ChannelMode = 3
	// C, L+R, Rear
	Mode_1_2_1 ChannelMode = 4
	// C, L+R, LS+RS
	Mode_1_2_2 ChannelMode = 5
	// C, L+R, LS+RS, LFE
	Mode_1_2_2_1 ChannelMode = 6
	// C, LC+RC, L+R, LS+RS, LFE
	Mode_1_2_2_2_1 ChannelMode = 7

	// C, L+R, LS+RS, Crear, LFE
	Mode_6_1 ChannelMode = 11
	// C, L+R, LS+RS, Lrear+Rrear, LFE
	Mode_7_1_Back ChannelMode = 12
	// C, L+R, LS+RS, LFE, Ltop+Rtop
	Mode_7_1_Top_Front ChannelMode = 14

	// C, L+R, LS+RS, Lrear+Rrear, LFE
	Mode_7_1_Rear_Surround ChannelMode = 33
	// C, LC+RC, L+R, LS+RS, LFE
	Mode_7_1_Front_Center ChannelMode = 34

	// 212 configuration, used in ELDv2
	Mode_212 ChannelMode = 128
)

// Speaker description tags.
// segmentation:
// - Bit 0-3: Horizontal postion (0: none, 1: front, 2: side, 3: back, 4: lfe)
// - Bit 4-7: Vertical position (0: normal, 1: top, 2: bottom)
type AudioChannelType uint8

const (
	ActNone AudioChannelType = 0x00
	// Front speaker position (at normal height)
	ActFront AudioChannelType = 0x01
	// Side speaker position (at normal height)
	ActSide AudioChannelType = 0x02
	// Back speaker position (at normal height)
	ActBack AudioChannelType = 0x03
	// Low frequency effect speaker postion (front)
	ActLfe AudioChannelType = 0x04
	// Top speaker area (for combination with speaker positions)
	ActTop AudioChannelType = 0x10
	// Top front speaker = (ACT_FRONT|ACT_TOP)
	ActFrontTop AudioChannelType = 0x11
	// Top side speaker  = (ACT_SIDE |ACT_TOP)
	ActSideTop AudioChannelType = 0x12
	// Top back speaker  = (ACT_BACK |ACT_TOP)
	ActBackTop AudioChannelType = 0x13
	// Bottom speaker area (for combination with speaker positions)
	ActBottom AudioChannelType = 0x20
	// Bottom front speaker = (ACT_FRONT|ACT_BOTTOM)
	ActFrontBottom AudioChannelType = 0x21
	// Bottom side speaker  = (ACT_SIDE |ACT_BOTTOM)
	ActSideBottom AudioChannelType = 0x22
	// Bottom back speaker  = (ACT_BACK |ACT_BOTTOM)
	ActBackBottom AudioChannelType = 0x23
)

// SBR Mode
type SbrMode int

const (
	SbrModeDefault SbrMode = iota
	SbrModeDisable
	SbrModeEnable
)

// Channel Order
type ChannelOrder int

const (
	ChannelOrderMpeg ChannelOrder = iota
	ChannelOrderWav
)

const (
	defaultMaxChannels    = 2
	defaultSampleBitdepth = 16
	defaultAOT            = AotAacLc
	defaultSamplerate     = 44100
	defaultBitrate        = 48000
)
