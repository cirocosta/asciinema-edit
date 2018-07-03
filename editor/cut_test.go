package editor_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/cirocosta/asciinema-edit/editor"
)

var (
	event1 = &cast.Event{
		Time: 1,
		Data: "event1",
	}
	event1_2 = &cast.Event{
		Time: 1.2,
		Data: "event1_2",
	}
	event1_6 = &cast.Event{
		Time: 1.6,
		Data: "event1_6",
	}
	event2 = &cast.Event{
		Time: 2,
		Data: "event2",
	}
)

var _ = Describe("Cut", func() {
	Context("with nil cast", func() {
		It("fails", func() {
			err := editor.Cut(nil, 1, 2)
			Expect(err).ToNot(Succeed())
		})
	})

	Context("with an empty event stream", func() {
		var data = &cast.Cast{
			EventStream: []*cast.Event{},
		}

		It("errors", func() {
			err := editor.Cut(data, 1, 2)
			Expect(err).ToNot(Succeed())
		})
	})

	Context("with non-empty event stream", func() {
		var (
			err  error
			data *cast.Cast
		)

		BeforeEach(func() {
			data = &cast.Cast{
				EventStream: []*cast.Event{
					event1,
					event1_2,
					event1_6,
					event2,
				},
			}
		})

		It("fails if `from` > `to`", func() {
			err = editor.Cut(data, 3, 2)
			Expect(err).ToNot(Succeed())
		})
	})
})
