package unmapify_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestUnmapify(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Unmapify Suite")
}
