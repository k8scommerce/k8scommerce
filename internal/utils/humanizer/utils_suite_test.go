package humanizer_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHumanizer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Humanizer Suite")
}
