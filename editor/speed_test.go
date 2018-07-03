package editor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/cirocosta/asciinema-edit/editor"
)

var _ = Describe("Speed", func() {
	Describe("parameter validation", func() {
		var data *cast.Cast

		BeforeEach(func() {
			data = &cast.Cast{
				EventStream: []*cast.Event{
					&cast.Event{},
				},
			}
		})

		Context("with nil cast", func() {
			It("fails", func() {
				err := editor.Speed(nil, 1, 2, 3)
				Expect(err).ToNot(Succeed())
			})
		})

		Context("with an empty event stream", func() {
			var data = &cast.Cast{
				EventStream: []*cast.Event{},
			}

			It("errors", func() {
				err := editor.Speed(data, 1, 1, 2)
				Expect(err).ToNot(Succeed())
			})
		})

		Context("with unusual factors", func() {
			It("fails with factor > 10", func() {
				err := editor.Speed(data, 12, 2, 3)
				Expect(err).ToNot(Succeed())

			})

			It("fails with factor < 0.1", func() {
				err := editor.Speed(data, 0.05, 2, 3)
				Expect(err).ToNot(Succeed())
			})
		})

		Context("with invalid ranges", func() {
			It("fails with `from` == `to`", func() {
				err := editor.Speed(data, 1, 2, 2)
				Expect(err).ToNot(Succeed())
			})

			It("fails with `from` > `to`", func() {
				err := editor.Speed(data, 1, 10, 2)
				Expect(err).ToNot(Succeed())
			})
		})
	})

	Context("with non-empty event stream", func() {
		var (
			data                           *cast.Cast
			event1, event2, event3, event4 *cast.Event
			err                            error
		)

		BeforeEach(func() {
			event1 = &cast.Event{
				Time: 1,
				Data: "event1",
			}
			event2 = &cast.Event{
				Time: 2,
				Data: "event2",
			}
			event3 = &cast.Event{
				Time: 3,
				Data: "event3",
			}
			event4 = &cast.Event{
				Time: 4,
				Data: "event4",
			}

			data = &cast.Cast{
				EventStream: []*cast.Event{
					event1,
					event2,
					event3,
					event4,
				},
			}
		})

		Context("with a slowing down factor and range", func() {
			JustBeforeEach(func() {
				err = editor.Speed(data, 2, 1, 2)
			})

			It("succeeds", func() {
				Expect(err).To(Succeed())
			})

			It("properly updates the timestamps", func() {
				Skip("TODO")
				Expect(event1.Time).To(Equal(float64(1)))
				Expect(event2.Time).To(Equal(float64(4)))
				Expect(event3.Time).To(Equal(float64(5)))
				Expect(event4.Time).To(Equal(float64(6)))
			})
		})
	})
})
