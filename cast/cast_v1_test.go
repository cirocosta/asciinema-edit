package cast_test

import (
	"bytes"
	"fmt"
	"io"

	"github.com/cirocosta/asciinema-edit/cast"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var streamV1 = `{
  "version": 1,
  "width": 80,
  "height": 24,
  "duration": 1.515658,
  "command": "/bin/zsh",
  "title": "a title",
  "env": {
    "TERM": "xterm-256color",
    "SHELL": "/bin/zsh"
  },
  "stdout": [
    [
      0.248848,
      "\u001b[1;31mHello \u001b[32mWorld!\u001b[0m\n"
    ],
    [
      1.001376,
      "I am \rThis is on the next line."
    ]
  ]
}`

var _ = Describe("CastV1", func() {
	Describe("DecodeV1", func() {
		Context("with non-empty reader", func() {
			FContext("w/ proper stream", func() {
				var (
					reader      io.Reader
					decodedCast *cast.CastV1
					err         error
				)

				BeforeEach(func() {
					reader = bytes.NewBufferString(streamV1)
					decodedCast, err = cast.DecodeV1(reader)
				})

				It("succeeds", func() {
					Expect(err).To(Succeed())
					Expect(decodedCast).ToNot(BeNil())
					fmt.Printf("=------------------\n")
					fmt.Printf("%+v\n", decodedCast)
				})

				It("has header values set", func() {
					Expect(decodedCast.Version).
						To(Equal(uint8(1)))
					Expect(decodedCast.Width).
						To(Equal(uint(80)))
					Expect(decodedCast.Height).
						To(Equal(uint(24)))
				})

				It("captures the events", func() {
					Expect(len(decodedCast.Stdout)).To(Equal(2))
					Expect(decodedCast[0].)).To(Equal(2))
				})
			})
		})
	})
})
