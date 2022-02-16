package logic_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestCartLogic(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CartLogic Suite")
}
