package transformer_test

import (
	"io/ioutil"
	"os"
	"path"

	"github.com/cirocosta/asciinema-edit/cast"
	"github.com/cirocosta/asciinema-edit/commands/transformer"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type DummyTransformation struct{}

func (t *DummyTransformation) Transform(c *cast.Cast) (err error) {
	return
}

var _ = Describe("Transformer", func() {
	Describe("New", func() {
		Context("with nil transform", func() {
			It("fails", func() {
				_, err := transformer.New(nil, "", "")
				Expect(err).ToNot(Succeed())
			})
		})

		Context("with transformation", func() {
			var (
				transformation transformer.Transformation
				tempDir        string
				err            error
			)

			BeforeEach(func() {
				transformation = &DummyTransformation{}
				tempDir, err = ioutil.TempDir("", "")
				Expect(err).To(Succeed())
			})

			AfterEach(func() {
				os.RemoveAll(tempDir)
			})

			Context("having empty input and output", func() {
				It("succeeds", func() {
					_, err = transformer.New(transformation, "", "")
					Expect(err).To(Succeed())
				})
			})

			Context("with input specified", func() {
				var input string

				It("fails if it doesnt exist", func() {
					input = path.Join(tempDir, "inexistent")

					_, err = transformer.New(transformation, input, "")
					Expect(err).ToNot(Succeed())
				})

				It("fails if is a directory", func() {
					input = tempDir
					_, err = transformer.New(transformation, input, "")
					Expect(err).ToNot(Succeed())
				})

				It("succeeds if file exists", func() {
					inputFile, err := ioutil.TempFile(tempDir, "")
					Expect(err).To(Succeed())

					input = inputFile.Name()

					_, err = transformer.New(transformation,
						inputFile.Name(), "")
					Expect(err).To(Succeed())
				})
			})

			Context("with output specified", func() {
				var output string

				It("fails if directory doesnt exist", func() {
					output = "/inexistent/directory/file.txt"

					_, err = transformer.New(transformation, "", output)
					Expect(err).ToNot(Succeed())
				})

				It("creates file if it doesnt exist in existing directory", func() {
					output = path.Join(tempDir, "output-file")

					_, err = transformer.New(transformation, "", output)
					Expect(err).To(Succeed())

					_, err = os.Stat(output)
					Expect(err).To(Succeed())
				})

				It("succeeds if file exists", func() {
					outputFile, err := ioutil.TempFile(tempDir, "")
					Expect(err).To(Succeed())

					output = outputFile.Name()

					_, err = transformer.New(transformation, "", output)
					Expect(err).To(Succeed())
				})
			})
		})
	})
})
