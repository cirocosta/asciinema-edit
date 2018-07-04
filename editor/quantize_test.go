package editor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/cirocosta/asciinema-edit/editor"
)

var _ = Describe("Quantize", func() {
	Describe("parameter validation", func() {
		var data *cast.Cast

		BeforeEach(func() {
			data = &cast.Cast{
				EventStream: []*cast.Event{
					{},
				},
			}
		})

		Context("with nil cast", func() {
			It("fails", func() {
				err := editor.Quantize(nil, nil)
				Expect(err).ToNot(Succeed())
			})
		})

		Context("with an empty event stream", func() {
			JustBeforeEach(func() {
				data = &cast.Cast{
					EventStream: []*cast.Event{},
				}
			})

			It("fails", func() {
				err := editor.Quantize(data, nil)
				Expect(err).ToNot(Succeed())
			})
		})

		Context("with a nil range list", func() {
			It("fails", func() {
				err := editor.Quantize(data, nil)
				Expect(err).ToNot(Succeed())
			})
		})

		Context("with an empty range list", func() {
			It("fails", func() {
				err := editor.Quantize(data, []editor.QuantizeRange{})
				Expect(err).ToNot(Succeed())
			})
		})
	})
})
