package cast_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCast(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cast Suite")
}
