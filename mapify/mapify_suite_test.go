package mapify_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestMapify(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Mapify Suite")
}
