package fdkaac

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("AAC enc test", func() {
	var (
		encoder *AacEncoder
		err     error
	)

	BeforeEach(func() {
		encoder = nil
		err = nil
	})

	It("Encoder create and close", func() {
		encoder, err = CreateAccEncoder(&AacEncoderConfig{
			TransMux: TtMp4Adts,
		})
		Expect(err).To(BeNil())
		Expect(encoder.ph).NotTo(BeNil())

		info, err := encoder.GetInfo()
		Expect(err).To(BeNil())
		Expect(info).NotTo(BeNil())
		Expect(info.FrameLength).To(Equal(uint(1024)))
		Expect(len(info.ConfBuf)).To(Equal(2))
		Expect(info.ConfBuf[0]).To(Equal(uint8(18)))
		Expect(info.ConfBuf[1]).To(Equal(uint8(16)))

		err = encoder.Close()
		Expect(err).To(BeNil())
		Expect(encoder.ph).To(BeNil())
	})

	It("Encode", func() {
		encoder, err = CreateAccEncoder(&AacEncoderConfig{
			TransMux: TtMp4Adts,
		})
		output := make([]byte, 1024)

		n, err := encoder.Encode(PCM0, output)
		Expect(err).To(BeNil())
		Expect(n).To(Equal(139))

		n, err = encoder.Flush(output)
		Expect(err).To(BeNil())
		Expect(n).To(Equal(417))
	})
})

var PCM0 = []byte{
	0x32, 0x04, 0xf4, 0xfd, 0x0d, 0x05, 0xaa, 0xfe, 0xfc, 0x05, 0x8b, 0xff, 0x0e, 0x07, 0x9f, 0x00, 0x44, 0x08, 0xdc, 0x01, 0x88, 0x09, 0x21, 0x03, 0xaa, 0x0a, 0x3f, 0x04, 0x75, 0x0b, 0x04, 0x05, 0xc5, 0x0b, 0x54, 0x05, 0xa0, 0x0b, 0x39, 0x05, 0x3e, 0x0b, 0xe6, 0x04, 0xf1, 0x0a, 0xa6, 0x04, 0x06, 0x0b, 0xbf, 0x04, 0x9f, 0x0b, 0x57, 0x05, 0xa7, 0x0c, 0x5f, 0x06, 0xd6, 0x0d, 0xa0, 0x07, 0xdc, 0x0e, 0xcc, 0x08, 0x7f, 0x0f, 0x9f, 0x09, 0xb5, 0x0f, 0xfc, 0x09, 0xa1, 0x0f, 0xed, 0x09, 0x7b, 0x0f, 0x9f, 0x09, 0x72, 0x0f, 0x4c, 0x09, 0x9c, 0x0f, 0x29, 0x09, 0xf5, 0x0f, 0x53, 0x09, 0x70, 0x10, 0xd0, 0x09, 0x05, 0x11, 0x91, 0x0a, 0xb7, 0x11, 0x78, 0x0b, 0x8c, 0x12, 0x63, 0x0c, 0x7b, 0x13, 0x2e, 0x0d, 0x6c, 0x14, 0xc4, 0x0d, 0x32, 0x15, 0x1d, 0x0e, 0xa9, 0x15, 0x43, 0x0e,
	0xc4, 0x15, 0x4e, 0x0e, 0x98, 0x15, 0x56, 0x0e, 0x54, 0x15, 0x6e, 0x0e, 0x2c, 0x15, 0x99, 0x0e, 0x3d, 0x15, 0xcc, 0x0e, 0x81, 0x15, 0xf7, 0x0e, 0xd5, 0x15, 0x0e, 0x0f, 0x0b, 0x16, 0x0f, 0x0f, 0x04, 0x16, 0x00, 0x0f, 0xbe, 0x15, 0xe9, 0x0e, 0x53, 0x15, 0xce, 0x0e, 0xe6, 0x14, 0xab, 0x0e, 0x93, 0x14, 0x7a, 0x0e, 0x5e, 0x14, 0x36, 0x0e, 0x36, 0x14, 0xde, 0x0d, 0x02, 0x14, 0x7d, 0x0d, 0xb3, 0x13, 0x1d, 0x0d, 0x4d, 0x13, 0xc5, 0x0c, 0xec, 0x12, 0x7a, 0x0c, 0xad, 0x12, 0x3b, 0x0c, 0xa2, 0x12, 0x0a, 0x0c, 0xc0, 0x12, 0xe2, 0x0b, 0xe7, 0x12, 0xc2, 0x0b, 0xed, 0x12, 0xa0, 0x0b, 0xb2, 0x12, 0x6f, 0x0b, 0x38, 0x12, 0x25, 0x0b, 0xa0, 0x11, 0xbc, 0x0a, 0x19, 0x11, 0x3e, 0x0a, 0xca, 0x10, 0xc3, 0x09, 0xbc, 0x10, 0x65, 0x09, 0xd4, 0x10, 0x38, 0x09, 0xe6, 0x10, 0x3a, 0x09,
	0xca, 0x10, 0x5a, 0x09, 0x77, 0x10, 0x7b, 0x09, 0x07, 0x10, 0x86, 0x09, 0xa5, 0x0f, 0x70, 0x09, 0x72, 0x0f, 0x3d, 0x09, 0x6b, 0x0f, 0xf6, 0x08, 0x67, 0x0f, 0x9f, 0x08, 0x2b, 0x0f, 0x34, 0x08, 0x8c, 0x0e, 0xaa, 0x07, 0x87, 0x0d, 0xf9, 0x06, 0x47, 0x0c, 0x21, 0x06, 0x0e, 0x0b, 0x31, 0x05, 0x16, 0x0a, 0x40, 0x04, 0x75, 0x09, 0x60, 0x03, 0x17, 0x09, 0x9f, 0x02, 0xd3, 0x08, 0x03, 0x02, 0x8c, 0x08, 0x8f, 0x01, 0x3e, 0x08, 0x4b, 0x01, 0x02, 0x08, 0x39, 0x01, 0xf6, 0x07, 0x56, 0x01, 0x20, 0x08, 0x8e, 0x01, 0x63, 0x08, 0xbe, 0x01, 0x8a, 0x08, 0xc3, 0x01, 0x67, 0x08, 0x86, 0x01, 0xef, 0x07, 0x0f, 0x01, 0x49, 0x07, 0x85, 0x00, 0xb7, 0x06, 0x1d, 0x00, 0x77, 0x06, 0x04, 0x00, 0x9e, 0x06, 0x48, 0x00, 0x0c, 0x07, 0xcc, 0x00, 0x78, 0x07, 0x57, 0x01, 0x98, 0x07, 0xa7, 0x01,
	0x3f, 0x07, 0x90, 0x01, 0x74, 0x06, 0x0a, 0x01, 0x6c, 0x05, 0x35, 0x00, 0x6d, 0x04, 0x48, 0xff, 0xaf, 0x03, 0x7d, 0xfe, 0x4a, 0x03, 0xfd, 0xfd, 0x30, 0x03, 0xd2, 0xfd, 0x3e, 0x03, 0xeb, 0xfd, 0x4b, 0x03, 0x21, 0xfe, 0x3d, 0x03, 0x4a, 0xfe, 0x0a, 0x03, 0x42, 0xfe, 0xb7, 0x02, 0xf9, 0xfd, 0x51, 0x02, 0x75, 0xfd, 0xe5, 0x01, 0xd0, 0xfc, 0x7a, 0x01, 0x2b, 0xfc, 0x14, 0x01, 0xa6, 0xfb, 0xb5, 0x00, 0x50, 0xfb, 0x5e, 0x00, 0x2a, 0xfb, 0x11, 0x00, 0x25, 0xfb, 0xd0, 0xff, 0x2d, 0xfb, 0x9c, 0xff, 0x32, 0xfb, 0x76, 0xff, 0x29, 0xfb, 0x5a, 0xff, 0x15, 0xfb, 0x44, 0xff, 0x02, 0xfb, 0x32, 0xff, 0xfd, 0xfa, 0x27, 0xff, 0x1a, 0xfb, 0x2b, 0xff, 0x63, 0xfb, 0x47, 0xff, 0xdd, 0xfb, 0x7f, 0xff, 0x7b, 0xfc, 0xc8, 0xff, 0x1f, 0xfd, 0x0d, 0x00, 0xa2, 0xfd, 0x31, 0x00, 0xe1, 0xfd,
	0x1d, 0x00, 0xca, 0xfd, 0xce, 0xff, 0x69, 0xfd, 0x54, 0xff, 0xe3, 0xfc, 0xd2, 0xfe, 0x69, 0xfc, 0x6a, 0xfe, 0x24, 0xfc, 0x32, 0xfe, 0x1f, 0xfc, 0x2a, 0xfe, 0x4d, 0xfc, 0x46, 0xfe, 0x8c, 0xfc, 0x6f, 0xfe, 0xbd, 0xfc, 0x94, 0xfe, 0xd2, 0xfc, 0xae, 0xfe, 0xd0, 0xfc, 0xb7, 0xfe, 0xc6, 0xfc, 0xa9, 0xfe, 0xbc, 0xfc, 0x7d, 0xfe, 0xb1, 0xfc, 0x2e, 0xfe, 0x9d, 0xfc, 0xc9, 0xfd, 0x80, 0xfc, 0x6d, 0xfd, 0x6e, 0xfc, 0x49, 0xfd, 0x8f, 0xfc, 0x8b, 0xfd, 0x0d, 0xfd, 0x45, 0xfe, 0xfc, 0xfd, 0x5f, 0xff, 0x47, 0xff, 0xa0, 0x00, 0xb8, 0x00, 0xbc, 0x01, 0x06, 0x02, 0x78, 0x02, 0xf8, 0x02, 0xbd, 0x02, 0x79, 0x03, 0x9c, 0x02, 0x9e, 0x03, 0x40, 0x02, 0x8e, 0x03, 0xda, 0x01, 0x73, 0x03, 0x8b, 0x01, 0x61, 0x03, 0x5f, 0x01, 0x59, 0x03, 0x52, 0x01, 0x52, 0x03, 0x5c, 0x01, 0x4a, 0x03,
	0x78, 0x01, 0x4a, 0x03, 0xa0, 0x01, 0x63, 0x03, 0xcc, 0x01, 0x98, 0x03, 0xea, 0x01, 0xdd, 0x03, 0xe4, 0x01, 0x10, 0x04, 0xaa, 0x01, 0x0f, 0x04, 0x3d, 0x01, 0xc8, 0x03, 0xb5, 0x00, 0x4b, 0x03, 0x34, 0x00, 0xbf, 0x02, 0xe0, 0xff, 0x5a, 0x02, 0xcb, 0xff, 0x43, 0x02, 0xf1, 0xff, 0x89, 0x02, 0x3b, 0x00, 0x18, 0x03, 0x8a, 0x00, 0xcb, 0x03, 0xc7, 0x00, 0x78, 0x04, 0xe6, 0x00, 0x00, 0x05, 0xeb, 0x00, 0x54, 0x05, 0xe0, 0x00, 0x77, 0x05, 0xcd, 0x00, 0x77, 0x05, 0xb5, 0x00, 0x65, 0x05, 0x97, 0x00, 0x54, 0x05, 0x6d, 0x00, 0x4f, 0x05, 0x35, 0x00, 0x58, 0x05, 0xec, 0xff, 0x65, 0x05, 0x96, 0xff, 0x64, 0x05, 0x34, 0xff, 0x3e, 0x05, 0xc9, 0xfe, 0xe1, 0x04, 0x56, 0xfe, 0x4f, 0x04, 0xdc, 0xfd, 0x99, 0x03, 0x5f, 0xfd, 0xe1, 0x02, 0xea, 0xfc, 0x4d, 0x02, 0x89, 0xfc, 0xfc, 0x01,
	0x4b, 0xfc, 0xf4, 0x01, 0x38, 0xfc, 0x24, 0x02, 0x44, 0xfc, 0x64, 0x02, 0x51, 0xfc, 0x83, 0x02, 0x30, 0xfc, 0x53, 0x02, 0xb8, 0xfb, 0xc0, 0x01, 0xd8, 0xfa, 0xd5, 0x00, 0xa9, 0xf9, 0xc1, 0xff, 0x6f, 0xf8, 0xc7, 0xfe, 0x7a, 0xf7, 0x29, 0xfe, 0x0a, 0xf7, 0x0c, 0xfe, 0x29, 0xf7, 0x69, 0xfe, 0xa3, 0xf7, 0x0c, 0xff, 0x1f, 0xf8, 0xa9, 0xff, 0x47, 0xf8, 0xfa, 0xff, 0xf4, 0xf7, 0xdd, 0xff, 0x3e, 0xf7, 0x5f, 0xff, 0x6f, 0xf6, 0xb8, 0xfe, 0xd7, 0xf5, 0x27, 0xfe, 0xa4, 0xf5, 0xdc, 0xfd, 0xcb, 0xf5, 0xe0, 0xfd, 0x14, 0xf6, 0x19, 0xfe, 0x3f, 0xf6, 0x5c, 0xfe, 0x29, 0xf6, 0x7e, 0xfe, 0xd8, 0xf5, 0x67, 0xfe, 0x6f, 0xf5, 0x11, 0xfe, 0x0e, 0xf5, 0x81, 0xfd, 0xb1, 0xf4, 0xbd, 0xfc, 0x33, 0xf4, 0xc9, 0xfb, 0x5f, 0xf3, 0xae, 0xfa, 0x1e, 0xf2, 0x7a, 0xf9, 0x89, 0xf0, 0x47, 0xf8,
	0xe8, 0xee, 0x30, 0xf7, 0x8f, 0xed, 0x48, 0xf6, 0xb6, 0xec, 0x91, 0xf5, 0x5a, 0xec, 0x00, 0xf5, 0x45, 0xec, 0x88, 0xf4, 0x2c, 0xec, 0x23, 0xf4, 0xda, 0xeb, 0xd8, 0xf3, 0x4b, 0xeb, 0xb2, 0xf3, 0xaa, 0xea, 0xb6, 0xf3, 0x2d, 0xea, 0xd7, 0xf3, 0xf6, 0xe9, 0xf7, 0xf3, 0xfc, 0xe9, 0xf0, 0xf3, 0x0a, 0xea, 0xab, 0xf3, 0xe5, 0xe9, 0x26, 0xf3, 0x6a, 0xe9, 0x78, 0xf2, 0xa5, 0xe8, 0xc7, 0xf1, 0xc8, 0xe7, 0x2f, 0xf1, 0x12, 0xe7, 0xc0, 0xf0, 0xaf, 0xe6, 0x73, 0xf0, 0xa1, 0xe6, 0x38, 0xf0, 0xc1, 0xe6, 0xff, 0xef, 0xd6, 0xe6, 0xbd, 0xef, 0xb1, 0xe6, 0x6c, 0xef, 0x47, 0xe6, 0x0a, 0xef, 0xad, 0xe5, 0x96, 0xee, 0x0d, 0xe5, 0x0f, 0xee, 0x89, 0xe4, 0x7c, 0xed, 0x2b, 0xe4, 0xe6, 0xec, 0xdc, 0xe3, 0x5e, 0xec, 0x7c, 0xe3, 0xef, 0xeb, 0xf1, 0xe2, 0x9a, 0xeb, 0x41, 0xe2, 0x59, 0xeb,
	0x93, 0xe1, 0x22, 0xeb, 0x18, 0xe1, 0xf1, 0xea, 0xf4, 0xe0, 0xca, 0xea, 0x26, 0xe1, 0xba, 0xea, 0x83, 0xe1, 0xc9, 0xea, 0xca, 0xe1, 0xf4, 0xea, 0xc9, 0xe1, 0x2a, 0xeb, 0x76, 0xe1, 0x56, 0xeb, 0xfb, 0xe0, 0x67, 0xeb, 0xa2, 0xe0, 0x63, 0xeb, 0xae, 0xe0, 0x61, 0xeb, 0x41, 0xe1, 0x87, 0xeb, 0x48, 0xe2, 0xf8, 0xeb, 0x8e, 0xe3, 0xc1, 0xec, 0xd5, 0xe4, 0xdb, 0xed, 0xf6, 0xe5, 0x23, 0xef, 0xe6, 0xe6, 0x6a, 0xf0, 0xad, 0xe7, 0x7b, 0xf1, 0x4f, 0xe8, 0x28, 0xf2, 0xc0, 0xe8, 0x5b, 0xf2, 0xeb, 0xe8, 0x1a, 0xf2, 0xca, 0xe8, 0x93, 0xf1, 0x7d, 0xe8, 0x11, 0xf1, 0x4b, 0xe8, 0xed, 0xf0, 0x88, 0xe8, 0x6d, 0xf1, 0x6c, 0xe9, 0xa2, 0xf2, 0xe7, 0xea, 0x5f, 0xf4, 0xa5, 0xec, 0x40, 0xf6, 0x29, 0xee, 0xce, 0xf7, 0x06, 0xef, 0xb0, 0xf8, 0x18, 0xef, 0xd0, 0xf8, 0x92, 0xee, 0x66, 0xf8,
	0xea, 0xed, 0xdf, 0xf7, 0xa1, 0xed, 0xae, 0xf7, 0x0e, 0xee, 0x24, 0xf8, 0x3d, 0xef, 0x53, 0xf9, 0xf8, 0xf0, 0x10, 0xfb, 0xeb, 0xf2, 0x0e, 0xfd, 0xcc, 0xf4, 0x02, 0xff, 0x76, 0xf6, 0xb9, 0x00, 0xed, 0xf7, 0x21, 0x02, 0x44, 0xf9, 0x47, 0x03, 0x8c, 0xfa, 0x3f, 0x04, 0xc7, 0xfb, 0x22, 0x05, 0xea, 0xfc, 0xfe, 0x05, 0xec, 0xfd, 0xe0, 0x06, 0xd0, 0xfe, 0xcc, 0x07, 0xab, 0xff, 0xc8, 0x08, 0x9a, 0x00, 0xd6, 0x09, 0xb4, 0x01, 0xf9, 0x0a, 0x02, 0x03, 0x2e, 0x0c, 0x75, 0x04, 0x72, 0x0d, 0xf0, 0x05, 0xbc, 0x0e, 0x4e, 0x07, 0xff, 0x0f, 0x73, 0x08, 0x2d, 0x11, 0x4b, 0x09, 0x36, 0x12, 0xd9, 0x09, 0x0d, 0x13, 0x2a, 0x0a, 0xae, 0x13, 0x5b, 0x0a, 0x1c, 0x14, 0x82, 0x0a, 0x61, 0x14, 0xae, 0x0a, 0x86, 0x14, 0xdb, 0x0a, 0x8d, 0x14, 0xfb, 0x0a, 0x77, 0x14, 0xf8, 0x0a, 0x44, 0x14,
	0xc8, 0x0a, 0xfa, 0x13, 0x75, 0x0a, 0xa8, 0x13, 0x18, 0x0a, 0x61, 0x13, 0xd0, 0x09, 0x2e, 0x13, 0xb1, 0x09, 0x0e, 0x13, 0xbb, 0x09, 0xed, 0x12, 0xdc, 0x09, 0xb6, 0x12, 0xf7, 0x09, 0x5b, 0x12, 0xf5, 0x09, 0xe6, 0x11, 0xd4, 0x09, 0x70, 0x11, 0xa5, 0x09, 0x1c, 0x11, 0x82, 0x09, 0x05, 0x11, 0x8b, 0x09, 0x31, 0x11, 0xcf, 0x09, 0x94, 0x11, 0x55, 0x0a, 0x1d, 0x12, 0x13, 0x0b, 0xb9, 0x12, 0xf5, 0x0b, 0x61, 0x13, 0xdb, 0x0c, 0x0d, 0x14, 0x9f, 0x0d, 0xae, 0x14, 0x14, 0x0e, 0x24, 0x15, 0x15, 0x0e, 0x48, 0x15, 0x8e, 0x0d, 0xf3, 0x14, 0x84, 0x0c, 0x16, 0x14, 0x16, 0x0b, 0xbe, 0x12, 0x74, 0x09, 0x16, 0x11, 0xd4, 0x07, 0x54, 0x0f, 0x60, 0x06, 0xab, 0x0d, 0x35, 0x05, 0x41, 0x0c, 0x64, 0x04, 0x31, 0x0b, 0xf7, 0x03, 0x89, 0x0a, 0xf5, 0x03, 0x50, 0x0a, 0x54, 0x04, 0x81, 0x0a,
	0x01, 0x05, 0x03, 0x0b, 0xd5, 0x05, 0xac, 0x0b, 0xa0, 0x06, 0x45, 0x0c, 0x3b, 0x07, 0x9f, 0x0c, 0x96, 0x07, 0xa7, 0x0c, 0xb6, 0x07, 0x66, 0x0c, 0xae, 0x07, 0xfb, 0x0b, 0x8b, 0x07, 0x82, 0x0b, 0x43, 0x07, 0x02, 0x0b, 0xba, 0x06, 0x63, 0x0a, 0xd0, 0x05, 0x83, 0x09, 0x7d, 0x04, 0x49, 0x08, 0xe1, 0x02, 0xc0, 0x06, 0x3e, 0x01, 0x1a, 0x05, 0xe7, 0xff, 0xa4, 0x03, 0x1c, 0xff, 0xa2, 0x02, 0xf4, 0xfe, 0x3a, 0x02, 0x56, 0xff, 0x66, 0x02, 0x08, 0x00, 0xff, 0x02, 0xc7, 0x00, 0xc7, 0x03, 0x61, 0x01, 0x80, 0x04, 0xb6, 0x01, 0xf5, 0x04, 0xb9, 0x01, 0xfc, 0x04, 0x67, 0x01, 0x7e, 0x04, 0xc3, 0x00, 0x76, 0x03, 0xd7, 0xff, 0xfe, 0x01, 0xc0, 0xfe, 0x52, 0x00, 0xad, 0xfd, 0xc7, 0xfe, 0xdc, 0xfc, 0xb4, 0xfd, 0x88, 0xfc, 0x52, 0xfd, 0xd2, 0xfc, 0xaa, 0xfd, 0xb7, 0xfd, 0x8f, 0xfe,
	0x0a, 0xff, 0xb3, 0xff, 0x89, 0x00, 0xc6, 0x00, 0xee, 0x01, 0x98, 0x01, 0x09, 0x03, 0x23, 0x02, 0xce, 0x03, 0x8a, 0x02, 0x51, 0x04, 0xf9, 0x02, 0xb9, 0x04, 0x90, 0x03, 0x2b, 0x05, 0x4d, 0x04, 0xb5, 0x05, 0x0e, 0x05, 0x41, 0x06, 0x9d, 0x05, 0xa3, 0x06, 0xcb, 0x05, 0xaa, 0x06, 0x83, 0x05, 0x3b, 0x06, 0xd2, 0x04, 0x68, 0x05, 0xe9, 0x03, 0x73, 0x04, 0x07, 0x03, 0xaf, 0x03, 0x68, 0x02, 0x63, 0x03, 0x30, 0x02, 0xa9, 0x03, 0x5f, 0x02, 0x60, 0x04, 0xd8, 0x02, 0x40, 0x05, 0x6a, 0x03, 0x01, 0x06, 0xeb, 0x03, 0x7e, 0x06, 0x45, 0x04, 0xc8, 0x06, 0x7b, 0x04, 0x13, 0x07, 0xa4, 0x04, 0x92, 0x07, 0xd9, 0x04, 0x57, 0x08, 0x2b, 0x05, 0x47, 0x09, 0x9b, 0x05, 0x2c, 0x0a, 0x20, 0x06, 0xd8, 0x0a, 0xb1, 0x06, 0x4a, 0x0b, 0x4d, 0x07, 0xab, 0x0b, 0xf8, 0x07, 0x33, 0x0c, 0xae, 0x08,
	0xfd, 0x0c, 0x5e, 0x09, 0xef, 0x0d, 0xe7, 0x09, 0xb7, 0x0e, 0x27, 0x0a, 0xfa, 0x0e, 0x0c, 0x0a, 0x87, 0x0e, 0x9e, 0x09, 0x76, 0x0d, 0x01, 0x09, 0x27, 0x0c, 0x64, 0x08, 0x13, 0x0b, 0xee, 0x07, 0x95, 0x0a, 0xae, 0x07, 0xbe, 0x0a, 0x9c, 0x07, 0x55, 0x0b, 0xa2, 0x07, 0xfa, 0x0b, 0xb3, 0x07, 0x5e, 0x0c, 0xd0, 0x07, 0x66, 0x0c, 0x00, 0x08, 0x36, 0x0c, 0x49, 0x08, 0x0d, 0x0c, 0xa1, 0x08, 0x1e, 0x0c, 0xed, 0x08, 0x6e, 0x0c, 0x0a, 0x09, 0xd3, 0x0c, 0xe0, 0x08, 0x08, 0x0d, 0x6e, 0x08, 0xdb, 0x0c, 0xca, 0x07, 0x48, 0x0c, 0x19, 0x07, 0x7d, 0x0b, 0x7d, 0x06, 0xc5, 0x0a, 0x0f, 0x06, 0x68, 0x0a, 0xde, 0x05, 0x84, 0x0a, 0xeb, 0x05, 0x0c, 0x0b, 0x38, 0x06, 0xcf, 0x0b, 0xbf, 0x06, 0x96, 0x0c, 0x7a, 0x07, 0x48, 0x0d, 0x5d, 0x08, 0xf6, 0x0d, 0x63, 0x09, 0xd1, 0x0e, 0x90, 0x0a,
	0x10, 0x10, 0xf5, 0x0b, 0xca, 0x11, 0xa3, 0x0d, 0xdf, 0x13, 0x9a, 0x0f, 0xfc, 0x15, 0xb6, 0x11, 0xb5, 0x17, 0xad, 0x13, 0xba, 0x18, 0x25, 0x15, 0xf3, 0x18, 0xdb, 0x15, 0x96, 0x18, 0xc8, 0x15, 0x0f, 0x18, 0x33, 0x15, 0xd3, 0x17, 0x96, 0x14, 0x2f, 0x18, 0x72, 0x14, 0x29, 0x19, 0x12, 0x15, 0x89, 0x1a, 0x6c, 0x16, 0xf8, 0x1b, 0x29, 0x18, 0x28, 0x1d, 0xcb, 0x19, 0xf8, 0x1d, 0xeb, 0x1a, 0x79, 0x1e, 0x64, 0x1b, 0xd4, 0x1e, 0x5c, 0x1b, 0x34, 0x1f, 0x27, 0x1b, 0xac, 0x1f, 0x1d, 0x1b, 0x2f, 0x20, 0x6e, 0x1b, 0xa2, 0x20, 0x13, 0x1c, 0xed, 0x20, 0xd5, 0x1c, 0x02, 0x21, 0x6f, 0x1d, 0xe3, 0x20, 0xac, 0x1d, 0x9b, 0x20, 0x7a, 0x1d, 0x33, 0x20, 0xf1, 0x1c, 0xb6, 0x1f, 0x42, 0x1c, 0x31, 0x1f, 0xa5, 0x1b, 0xb6, 0x1e, 0x45, 0x1b, 0x5b, 0x1e, 0x31, 0x1b, 0x2e, 0x1e, 0x60, 0x1b,
	0x2a, 0x1e, 0xae, 0x1b, 0x36, 0x1e, 0xef, 0x1b, 0x2b, 0x1e, 0xf7, 0x1b, 0xe6, 0x1d, 0xae, 0x1b, 0x55, 0x1d, 0x0c, 0x1b, 0x81, 0x1c, 0x1f, 0x1a, 0x83, 0x1b, 0x02, 0x19, 0x79, 0x1a, 0xd2, 0x17, 0x7c, 0x19, 0xad, 0x16, 0x9a, 0x18, 0xa8, 0x15, 0xd4, 0x17, 0xcf, 0x14, 0x28, 0x17, 0x22, 0x14, 0x8e, 0x16, 0x92, 0x13, 0xf3, 0x15, 0xfd, 0x12, 0x3d, 0x15, 0x39, 0x12, 0x50, 0x14, 0x23, 0x11, 0x17, 0x13, 0xb0, 0x0f, 0x9b, 0x11, 0xff, 0x0d, 0x05, 0x10, 0x51, 0x0c, 0x97, 0x0e, 0xf5, 0x0a, 0x94, 0x0d, 0x2d, 0x0a, 0x21, 0x0d, 0x0b, 0x0a, 0x39, 0x0d, 0x71, 0x0a, 0xa9, 0x0d, 0x18, 0x0b, 0x2c, 0x0e, 0xb0, 0x0b, 0x80, 0x0e, 0xfa, 0x0b, 0x7b, 0x0e, 0xde, 0x0b, 0x0c, 0x0e, 0x5d, 0x0b, 0x37, 0x0d, 0x87, 0x0a, 0x05, 0x0c, 0x69, 0x09, 0x80, 0x0a, 0x09, 0x08, 0xbc, 0x08, 0x6f, 0x06,
	0xde, 0x06, 0xb2, 0x04, 0x1e, 0x05, 0xf8, 0x02, 0xb8, 0x03, 0x75, 0x01, 0xd2, 0x02, 0x4f, 0x00, 0x6a, 0x02, 0x8e, 0xff, 0x4f, 0x02, 0x16, 0xff, 0x32, 0x02, 0xb0, 0xfe, 0xc2, 0x01, 0x1f, 0xfe, 0xcf, 0x00, 0x38, 0xfd, 0x5b, 0xff, 0xef, 0xfb, 0x95, 0xfd, 0x56, 0xfa, 0xc5, 0xfb, 0x9a, 0xf8, 0x32, 0xfa, 0xf1, 0xf6, 0x0d, 0xf9, 0x90, 0xf5, 0x6a, 0xf8, 0xa7, 0xf4, 0x43, 0xf8, 0x51, 0xf4, 0x7b, 0xf8, 0x90, 0xf4, 0xed, 0xf8, 0x45, 0xf5, 0x6e, 0xf9, 0x32, 0xf6, 0xd8, 0xf9, 0x0e, 0xf7, 0x12, 0xfa, 0x9c, 0xf7, 0x16, 0xfa, 0xbd, 0xf7, 0xf0, 0xf9, 0x82, 0xf7, 0xbe, 0xf9, 0x1e, 0xf7, 0x9c, 0xf9, 0xce, 0xf6, 0x9f, 0xf9, 0xc2, 0xf6, 0xcd, 0xf9, 0x09, 0xf7, 0x1f, 0xfa, 0x8e, 0xf7, 0x86, 0xfa, 0x2b, 0xf8, 0xf2, 0xfa, 0xb7, 0xf8, 0x57, 0xfb, 0x16, 0xf9, 0xaf, 0xfb, 0x43, 0xf9,
	0xf3, 0xfb, 0x45, 0xf9, 0x24, 0xfc, 0x31, 0xf9, 0x4b, 0xfc, 0x24, 0xf9, 0x7b, 0xfc, 0x3a, 0xf9, 0xce, 0xfc, 0x8d, 0xf9, 0x55, 0xfd, 0x28, 0xfa, 0x11, 0xfe, 0x04, 0xfb, 0xea, 0xfe, 0xfe, 0xfb, 0xb6, 0xff, 0xe8, 0xfc, 0x48, 0x00, 0x94, 0xfd, 0x88, 0x00, 0xe9, 0xfd, 0x7d, 0x00, 0xf1, 0xfd, 0x47, 0x00, 0xd1, 0xfd, 0x1a, 0x00, 0xbe, 0xfd, 0x1e, 0x00, 0xe8, 0xfd, 0x69, 0x00, 0x68, 0xfe, 0xf8, 0x00, 0x3f, 0xff, 0xbb, 0x01, 0x58, 0x00, 0xa0, 0x02, 0x95, 0x01, 0x98, 0x03, 0xd6, 0x02, 0x99, 0x04, 0x02, 0x04, 0x98, 0x05, 0x04, 0x05, 0x84, 0x06, 0xcf, 0x05, 0x45, 0x07, 0x5f, 0x06, 0xc5, 0x07, 0xbb, 0x06, 0xf3, 0x07, 0xea, 0x06, 0xca, 0x07, 0xf1, 0x06, 0x4f, 0x07, 0xc6, 0x06, 0x8e, 0x06, 0x54, 0x06, 0x9b, 0x05, 0x88, 0x05, 0x91, 0x04, 0x67, 0x04, 0x95, 0x03, 0x1c, 0x03,
	0xd0, 0x02, 0xf5, 0x01, 0x65, 0x02, 0x46, 0x01, 0x61, 0x02, 0x4b, 0x01, 0xb8, 0x02, 0x01, 0x02, 0x44, 0x03, 0x28, 0x03, 0xd7, 0x03, 0x59, 0x04, 0x4a, 0x04, 0x34, 0x05, 0x8b, 0x04, 0x8d, 0x05, 0x9b, 0x04, 0x72, 0x05, 0x84, 0x04, 0x20, 0x05, 0x4c, 0x04, 0xd8, 0x04, 0xee, 0x03, 0xba, 0x04, 0x63, 0x03, 0xb9, 0x04, 0xa9, 0x02, 0xa7, 0x04, 0xd2, 0x01, 0x5a, 0x04, 0xf9, 0x00, 0xc0, 0x03, 0x3b, 0x00, 0xf1, 0x02, 0xa0, 0xff, 0x18, 0x02, 0x21, 0xff, 0x5f, 0x01, 0xa7, 0xfe, 0xda, 0x00, 0x2a, 0xfe, 0x84, 0x00, 0xbc, 0xfd, 0x57, 0x00, 0x8f, 0xfd, 0x58, 0x00, 0xda, 0xfd, 0xa5, 0x00, 0xc2, 0xfe, 0x64, 0x01, 0x3c, 0x00, 0xaa, 0x02, 0x0b, 0x02, 0x60, 0x04, 0xcc, 0x03, 0x39, 0x06, 0x1d, 0x05, 0xc4, 0x07, 0xb3, 0x05, 0x92, 0x08, 0x6f, 0x05, 0x5f, 0x08, 0x5c, 0x04, 0x2a, 0x07,
	0xa5, 0x02, 0x3a, 0x05, 0x86, 0x00, 0xfd, 0x02, 0x46, 0xfe, 0xe7, 0x00, 0x28, 0xfc, 0x45, 0xff, 0x68, 0xfa, 0x2b, 0xfe, 0x2d, 0xf9, 0x7c, 0xfd, 0x7a, 0xf8, 0xfd, 0xfc, 0x2a, 0xf8, 0x79, 0xfc, 0x01, 0xf8, 0xda, 0xfb, 0xbd, 0xf7, 0x2d, 0xfb, 0x37, 0xf7, 0x91, 0xfa, 0x70, 0xf6, 0x22, 0xfa, 0x90, 0xf5, 0xe4, 0xf9, 0xcf, 0xf4, 0xbd, 0xf9, 0x60, 0xf4, 0x8b, 0xf9, 0x57, 0xf4, 0x3b, 0xf9, 0xac, 0xf4, 0xe1, 0xf8, 0x43, 0xf5, 0xae, 0xf8, 0xf7, 0xf5, 0xd9, 0xf8, 0xa8, 0xf6, 0x74, 0xf9, 0x36, 0xf7, 0x54, 0xfa, 0x7a, 0xf7, 0x0f, 0xfb, 0x47, 0xf7, 0x28, 0xfb, 0x71, 0xf6, 0x40, 0xfa, 0xe5, 0xf4, 0x46, 0xf8, 0xbb, 0xf2, 0x89, 0xf5, 0x3a, 0xf0, 0x97, 0xf2, 0xcc, 0xed, 0x0a, 0xf0, 0xe2, 0xeb, 0x53, 0xee, 0xc3, 0xea, 0x8f, 0xed, 0x7b, 0xea, 0x91, 0xed, 0xc8, 0xea, 0xf2, 0xed,
	0x31, 0xeb, 0x3a, 0xee, 0x2e, 0xeb, 0x03, 0xee, 0x52, 0xea, 0x08, 0xed, 0x75, 0xe8, 0x34, 0xeb, 0xc2, 0xe5, 0xa5, 0xe8, 0xb0, 0xe2, 0xad, 0xe5, 0xd7, 0xdf, 0xc2, 0xe2, 0xc4, 0xdd, 0x67, 0xe0, 0xc6, 0xdc, 0x05, 0xdf, 0xd6, 0xdc, 0xc0, 0xde, 0xa9, 0xdd, 0x6e, 0xdf, 0xc9, 0xde, 0x9c, 0xe0, 0xca, 0xdf, 0xbc, 0xe1, 0x66, 0xe0, 0x55, 0xe2, 0x90, 0xe0, 0x33, 0xe2, 0x59, 0xe0, 0x67, 0xe1, 0xd7, 0xdf, 0x33, 0xe0, 0x14, 0xdf, 0xdd, 0xde, 0x08, 0xde, 0x8f, 0xdd, 0xb1, 0xdc, 0x50, 0xdc, 0x2c, 0xdb, 0x18, 0xdb, 0xb9, 0xd9, 0xeb, 0xd9, 0xad, 0xd8, 0xe8, 0xd8, 0x50, 0xd8, 0x42, 0xd8, 0xb4, 0xd8, 0x24, 0xd8, 0xab, 0xd9, 0x98, 0xd8, 0xda, 0xda, 0x79, 0xd9, 0xe2, 0xdb, 0x82, 0xda, 0x88, 0xdc, 0x6e, 0xdb, 0xd0, 0xdc, 0x13, 0xdc, 0xed, 0xdc, 0x72, 0xdc, 0x1d, 0xdd, 0xa8, 0xdc,
	0x8d, 0xdd, 0xdd, 0xdc, 0x3e, 0xde, 0x2c, 0xdd, 0x14, 0xdf, 0x9a, 0xdd, 0xea, 0xdf, 0x20, 0xde, 0xb2, 0xe0, 0xba, 0xde, 0x7c, 0xe1, 0x70, 0xdf, 0x66, 0xe2, 0x51, 0xe0, 0x87, 0xe3, 0x63, 0xe1, 0xdb, 0xe4, 0x96, 0xe2, 0x39, 0xe6, 0xc1, 0xe3, 0x69, 0xe7, 0xac, 0xe4, 0x3c, 0xe8, 0x2e, 0xe5, 0xa4, 0xe8, 0x3e, 0xe5, 0xc0, 0xe8, 0x06, 0xe5, 0xce, 0xe8, 0xcf, 0xe4, 0x12, 0xe9, 0xea, 0xe4, 0xbc, 0xe9, 0x8c, 0xe5, 0xcf, 0xea, 0xb5, 0xe6, 0x23, 0xec, 0x2d, 0xe8, 0x73, 0xed, 0x99, 0xe9, 0x79, 0xee, 0xa1, 0xea, 0x0a, 0xef, 0x17, 0xeb, 0x29, 0xef, 0x07, 0xeb, 0x02, 0xef, 0xb3, 0xea, 0xd5, 0xee, 0x6f, 0xea, 0xd3, 0xee, 0x7a, 0xea, 0x0f, 0xef, 0xe0, 0xea, 0x70, 0xef, 0x77, 0xeb, 0xc9, 0xef, 0xfa, 0xeb, 0xfa, 0xef, 0x38, 0xec, 0x01, 0xf0, 0x29, 0xec, 0x04, 0xf0, 0xf8, 0xeb,
	0x34, 0xf0, 0xe5, 0xeb, 0xac, 0xf0, 0x1c, 0xec, 0x54, 0xf1, 0x91, 0xec, 0xe6, 0xf1, 0x02, 0xed, 0x0e, 0xf2, 0x1c, 0xed, 0x95, 0xf1, 0xa5, 0xec, 0x8e, 0xf0, 0xa5, 0xeb, 0x4c, 0xef, 0x66, 0xea, 0x47, 0xee, 0x55, 0xe9, 0xe1, 0xed, 0xcf, 0xe8, 0x3f, 0xee, 0xf8, 0xe8, 0x37, 0xef, 0xb4, 0xe9, 0x74, 0xf0, 0xc2, 0xea, 0x9e, 0xf1, 0xdf, 0xeb, 0x85, 0xf2, 0xe8, 0xec, 0x28, 0xf3, 0xd7, 0xed, 0xa5, 0xf3, 0xb4, 0xee, 0x14, 0xf4, 0x75, 0xef, 0x73, 0xf4, 0xfc, 0xef, 0xa4, 0xf4, 0x21, 0xf0, 0x86, 0xf4, 0xcf, 0xef, 0x12, 0xf4, 0x21, 0xef, 0x6a, 0xf3, 0x5c, 0xee, 0xce, 0xf2, 0xd7, 0xed, 0x80, 0xf2, 0xd2, 0xed, 0xa5, 0xf2, 0x55, 0xee, 0x38, 0xf3, 0x34, 0xef, 0x14, 0xf4, 0x26, 0xf0, 0x0a, 0xf5, 0xf0, 0xf0, 0xfb, 0xf5, 0x82, 0xf1, 0xe3, 0xf6, 0xff, 0xf1, 0xd1, 0xf7, 0x9b, 0xf2,
	0xd1, 0xf8, 0x80, 0xf3, 0xdd, 0xf9, 0xa6, 0xf4, 0xd6, 0xfa, 0xd8, 0xf5, 0x89, 0xfb, 0xc0, 0xf6, 0xc9, 0xfb, 0x14, 0xf7, 0x7e, 0xfb, 0xb1, 0xf6, 0xb9, 0xfa, 0xb2, 0xf5, 0xa9, 0xf9, 0x63, 0xf4, 0x97, 0xf8, 0x29, 0xf3, 0xc6, 0xf7, 0x58, 0xf2, 0x66, 0xf7, 0x1e, 0xf2, 0x7a, 0xf7, 0x70, 0xf2, 0xe1, 0xf7, 0x14, 0xf3, 0x5f, 0xf8, 0xbf, 0xf3, 0xb7, 0xf8, 0x2d, 0xf4, 0xc8, 0xf8, 0x43, 0xf4, 0x99, 0xf8, 0x15, 0xf4, 0x5b, 0xf8, 0xd7, 0xf3, 0x4d, 0xf8, 0xca, 0xf3, 0xa2, 0xf8, 0x1b, 0xf4, 0x6b, 0xf9, 0xd6, 0xf4, 0x8c, 0xfa, 0xe0, 0xf5, 0xd2, 0xfb, 0x0f, 0xf7, 0x03, 0xfd, 0x34, 0xf8, 0xfa, 0xfd, 0x31, 0xf9, 0xae, 0xfe, 0xf3, 0xf9, 0x29, 0xff, 0x70, 0xfa, 0x7b, 0xff, 0xa2, 0xfa, 0xa9, 0xff, 0x85, 0xfa, 0xaf, 0xff, 0x1e, 0xfa, 0x85, 0xff, 0x89, 0xf9, 0x2e, 0xff, 0xed, 0xf8,
	0xbb, 0xfe, 0x78, 0xf8, 0x48, 0xfe, 0x41, 0xf8, 0xee, 0xfd, 0x40, 0xf8, 0xb8, 0xfd, 0x4b, 0xf8, 0x98, 0xfd, 0x28, 0xf8, 0x6d, 0xfd, 0xae, 0xf7, 0x14, 0xfd, 0xd8, 0xf6, 0x75, 0xfc, 0xc7, 0xf5, 0x92, 0xfb, 0xba, 0xf4, 0x90, 0xfa, 0xeb, 0xf3, 0xae, 0xf9, 0x84, 0xf3, 0x39, 0xf9, 0x94, 0xf3, 0x6f, 0xf9, 0x15, 0xf4, 0x66, 0xfa, 0xf5, 0xf4, 0xf8, 0xfb, 0x13, 0xf6, 0xc2, 0xfd, 0x42, 0xf7, 0x3d, 0xff, 0x3e, 0xf8, 0xe4, 0xff, 0xb5, 0xf8, 0x6c, 0xff, 0x5f, 0xf8, 0xda, 0xfd, 0x19, 0xf7, 0x88, 0xfb, 0xfe, 0xf4, 0xfa, 0xf8, 0x63, 0xf2, 0xab, 0xf6, 0xb8, 0xef, 0xe1, 0xf4, 0x63, 0xed, 0x9e, 0xf3, 0x9d, 0xeb, 0xb7, 0xf2, 0x68, 0xea, 0x00, 0xf2, 0xa2, 0xe9, 0x71, 0xf1, 0x29, 0xe9, 0x24, 0xf1, 0xef, 0xe8, 0x41, 0xf1, 0xf9, 0xe8, 0xd1, 0xf1, 0x51, 0xe9, 0xa6, 0xf2, 0xe2, 0xe9,
	0x60, 0xf3, 0x76, 0xea, 0x9d, 0xf3, 0xbe, 0xea, 0x2b, 0xf3, 0x7e, 0xea, 0x29, 0xf2, 0xac, 0xe9, 0xf9, 0xf0, 0x81, 0xe8, 0x0c, 0xf0, 0x5f, 0xe7, 0xae, 0xef, 0xa4, 0xe6, 0xd5, 0xef, 0x7b, 0xe6, 0x32, 0xf0, 0xc6, 0xe6, 0x59, 0xf0, 0x27, 0xe7, 0xff, 0xef, 0x37, 0xe7, 0x24, 0xef, 0xb5, 0xe6, 0x17, 0xee, 0xb0, 0xe5, 0x4e, 0xed, 0x86, 0xe4, 0x2e, 0xed, 0xbc, 0xe3, 0xe4, 0xed, 0xc4, 0xe3, 0x56, 0xef, 0xd3, 0xe4, 0x36, 0xf1, 0xc1, 0xe6, 0x2d, 0xf3, 0x22, 0xe9, 0xf9, 0xf4, 0x6e, 0xeb, 0x83, 0xf6, 0x3a, 0xed, 0xd2, 0xf7, 0x5f, 0xee, 0xfe, 0xf8, 0xfe, 0xee, 0x19, 0xfa, 0x6a, 0xef, 0x28, 0xfb, 0xfb, 0xef, 0x28, 0xfc, 0xe7, 0xf0, 0x19, 0xfd, 0x31, 0xf2, 0x08, 0xfe, 0xb4, 0xf3, 0x0e, 0xff, 0x3f, 0xf5, 0x44, 0x00, 0xaf, 0xf6, 0xad, 0x01, 0xfc, 0xf7, 0x2f, 0x03, 0x2e, 0xf9,
	0x8e, 0x04, 0x44, 0xfa, 0x81, 0x05, 0x26, 0xfb, 0xd1, 0x05, 0xa8, 0xfb, 0x75, 0x05, 0xa8, 0xfb, 0xa0, 0x04, 0x2b, 0xfb, 0xb4, 0x03, 0x6f, 0xfa, 0x22, 0x03, 0xdc, 0xf9, 0x3c, 0x03, 0xdb, 0xf9, 0x1a, 0x04, 0xa3, 0xfa, 0x92, 0x05, 0x21, 0xfc, 0x57, 0x07, 0x04, 0xfe, 0x18, 0x09, 0xe3, 0xff, 0xa4, 0x0a, 0x74, 0x01, 0xf1, 0x0b, 0xa1, 0x02, 0x07, 0x0d, 0x87, 0x03, 0xeb, 0x0d, 0x4a, 0x04, 0x89, 0x0e, 0xf5, 0x04, 0xbc, 0x0e, 0x6a, 0x05, 0x6e, 0x0e, 0x79, 0x05, 0xb5, 0x0d, 0x10, 0x05, 0xe2, 0x0c, 0x5a, 0x04, 0x6e, 0x0c, 0xc7, 0x03, 0xc9, 0x0c, 0xe4, 0x03, 0x2e, 0x0e, 0x19, 0x05, 0x80, 0x10, 0x78, 0x07, 0x5a, 0x13, 0xab, 0x0a, 0x37, 0x16, 0x18, 0x0e, 0xa3, 0x18, 0x1a, 0x11, 0x64, 0x1a, 0x41, 0x13, 0x7c, 0x1b, 0x76, 0x14, 0x1a, 0x1c, 0xee, 0x14, 0x77, 0x1c, 0x09, 0x15,
	0xc2, 0x1c, 0x25, 0x15, 0x16, 0x1d, 0x83, 0x15, 0x86, 0x1d, 0x38, 0x16, 0x1e, 0x1e, 0x38, 0x17, 0xe7, 0x1e, 0x65, 0x18, 0xda, 0x1f, 0x95, 0x19, 0xd7, 0x20, 0x97, 0x1a, 0xa7, 0x21, 0x3e, 0x1b, 0x0d, 0x22, 0x66, 0x1b, 0xdf, 0x21, 0x02, 0x1b, 0x22, 0x21, 0x2e, 0x1a, 0x10, 0x20, 0x2b, 0x19, 0x07, 0x1f, 0x51, 0x18, 0x69, 0x1e, 0xf4, 0x17, 0x7a, 0x1e, 0x44, 0x18, 0x47, 0x1f, 0x43, 0x19, 0xa4, 0x20, 0xbc, 0x1a, 0x3f, 0x22, 0x5c, 0x1c, 0xb9, 0x23, 0xca, 0x1d, 0xbf, 0x24, 0xc4, 0x1e, 0x22, 0x25, 0x29, 0x1f, 0xd9, 0x24, 0xff, 0x1e, 0x00, 0x24, 0x64, 0x1e, 0xc4, 0x22, 0x81, 0x1d, 0x5f, 0x21, 0x78, 0x1c, 0x03, 0x20, 0x66, 0x1b, 0xd7, 0x1e, 0x61, 0x1a, 0xf5, 0x1d, 0x80, 0x19, 0x6b, 0x1d, 0xdd, 0x18, 0x3c, 0x1d, 0x90, 0x18, 0x58, 0x1d, 0xa2, 0x18, 0x9e, 0x1d, 0xfc, 0x18,
	0xdb, 0x1d, 0x67, 0x19, 0xcf, 0x1d, 0x94, 0x19, 0x48, 0x1d, 0x3d, 0x19, 0x35, 0x1c, 0x45, 0x18, 0xbb, 0x1a, 0xca, 0x16, 0x27, 0x19, 0x25, 0x15, 0xe1, 0x17, 0xc5, 0x13, 0x39, 0x17, 0x0b, 0x13, 0x4d, 0x17, 0x1a, 0x13, 0xf5, 0x17, 0xcd, 0x13, 0xd3, 0x18, 0xc5, 0x14, 0x74, 0x19, 0x90, 0x15, 0x80, 0x19, 0xd0, 0x15, 0xcf, 0x18, 0x5b, 0x15, 0x78, 0x17, 0x3a, 0x14, 0xb7, 0x15, 0xa0, 0x12, 0xde, 0x13, 0xd5, 0x10, 0x37, 0x12, 0x23, 0x0f, 0xfc, 0x10, 0xd0, 0x0d, 0x4b, 0x10, 0x10, 0x0d, 0x25, 0x10, 0xfb, 0x0c, 0x6f, 0x10, 0x7e, 0x0d, 0xec, 0x10, 0x53, 0x0e, 0x4c, 0x11, 0x0e, 0x0f, 0x3e, 0x11, 0x37, 0x0f, 0x8c, 0x10, 0x81, 0x0e, 0x34, 0x0f, 0xe5, 0x0c, 0x69, 0x0d, 0xb3, 0x0a, 0x85, 0x0b, 0x71, 0x08, 0xe6, 0x09, 0xa8, 0x06, 0xcf, 0x08, 0xa9, 0x05, 0x4e, 0x08, 0x76, 0x05,
	0x46, 0x08, 0xc5, 0x05, 0x81, 0x08, 0x2f, 0x06, 0xc8, 0x08, 0x5b, 0x06, 0xf2, 0x08, 0x26, 0x06, 0xe8, 0x08, 0x9b, 0x05, 0x97, 0x08, 0xdf, 0x04, 0xef, 0x07, 0x0a, 0x04, 0xdc, 0x06, 0x19, 0x03, 0x59, 0x05, 0xec, 0x01, 0x77, 0x03, 0x66, 0x00, 0x64, 0x01, 0x86, 0xfe, 0x63, 0xff, 0x74, 0xfc, 0xb5, 0xfd, 0x7c, 0xfa, 0x85, 0xfc, 0xef, 0xf8, 0xdc, 0xfb, 0x00, 0xf8, 0xa0, 0xfb, 0xb1, 0xf7, 0xa3, 0xfb, 0xd3, 0xf7, 0xb6, 0xfb, 0x1b, 0xf8, 0xb6, 0xfb, 0x42, 0xf8, 0x95, 0xfb, 0x21, 0xf8, 0x5a, 0xfb, 0xbd, 0xf7, 0x14, 0xfb, 0x39, 0xf7, 0xd7, 0xfa, 0xc7, 0xf6, 0xb3, 0xfa, 0x87, 0xf6, 0xaa, 0xfa, 0x7a, 0xf6, 0xb2, 0xfa, 0x81, 0xf6, 0xb3, 0xfa, 0x6f, 0xf6, 0x93, 0xfa, 0x20, 0xf6, 0x37, 0xfa, 0x8a, 0xf5, 0x94, 0xf9, 0xbc, 0xf4, 0xb2, 0xf8, 0xd6, 0xf3, 0xa8, 0xf7, 0xf7, 0xf2,
	0x9b, 0xf6, 0x2b, 0xf2, 0xb3, 0xf5, 0x73, 0xf1, 0x11, 0xf5, 0xcf, 0xf0, 0xce, 0xf4, 0x53, 0xf0, 0xf8, 0xf4, 0x28, 0xf0, 0x8c, 0xf5, 0x80, 0xf0, 0x7a, 0xf6, 0x79, 0xf1, 0xa0, 0xf7, 0x01, 0xf3, 0xcf, 0xf8, 0xcb, 0xf4, 0xd4, 0xf9, 0x66, 0xf6, 0x80, 0xfa, 0x63, 0xf7, 0xb3, 0xfa, 0x87, 0xf7, 0x68, 0xfa, 0xdd, 0xf6, 0xb1, 0xf9, 0xb9, 0xf5, 0xb5, 0xf8, 0x89, 0xf4, 0xa9, 0xf7, 0xb1, 0xf3, 0xc9, 0xf6, 0x63, 0xf3, 0x4a, 0xf6, 0x9d, 0xf3, 0x53, 0xf6, 0x38, 0xf4, 0xef, 0xf6, 0x06, 0xf5, 0x03, 0xf8, 0xe6, 0xf5, 0x56, 0xf9, 0xca, 0xf6, 0x99, 0xfa, 0xb0, 0xf7, 0x87, 0xfb, 0x90, 0xf8, 0xfa, 0xfb, 0x5a, 0xf9, 0xfe, 0xfb, 0xf9, 0xf9, 0xc2, 0xfb, 0x5d, 0xfa, 0x88, 0xfb, 0x8d, 0xfa, 0x7e, 0xfb, 0x9f, 0xfa, 0xb3, 0xfb, 0xb4, 0xfa, 0x10, 0xfc, 0xe9, 0xfa, 0x74, 0xfc, 0x50, 0xfb,
	0xcb, 0xfc, 0xe9, 0xfb, 0x19, 0xfd, 0xaa, 0xfc, 0x74, 0xfd, 0x80, 0xfd, 0xef, 0xfd, 0x56, 0xfe, 0x80, 0xfe, 0x16, 0xff, 0x05, 0xff, 0xac, 0xff, 0x50, 0xff, 0x0d, 0x00, 0x4d, 0xff, 0x3f, 0x00, 0x17, 0xff, 0x60, 0x00, 0xf8, 0xfe, 0xa2, 0x00, 0x47, 0xff, 0x37, 0x01, 0x3f, 0x00, 0x43, 0x02, 0xdb, 0x01, 0xc8, 0x03, 0xcf, 0x03, 0x9d, 0x05, 0xa7, 0x05, 0x7a, 0x07, 0xfb, 0x06, 0x0f, 0x09, 0x93, 0x07, 0x1b, 0x0a, 0x7e, 0x07, 0x81, 0x0a, 0xfd, 0x06, 0x4e, 0x0a, 0x61, 0x06, 0xb4, 0x09, 0xe3, 0x05, 0xf4, 0x08, 0x94, 0x05, 0x4d, 0x08, 0x61, 0x05, 0xe8, 0x07, 0x30, 0x05, 0xd7, 0x07, 0xf7, 0x04, 0x0f, 0x08, 0xc5, 0x04, 0x78, 0x08, 0xbf, 0x04, 0xf6, 0x08, 0x04, 0x05, 0x78, 0x09, 0x99, 0x05, 0xf3, 0x09, 0x59, 0x06, 0x63, 0x0a, 0x04, 0x07, 0xc1, 0x0a, 0x57, 0x07, 0xfc, 0x0a,
}
