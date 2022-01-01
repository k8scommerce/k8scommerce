package logic_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCartLogic(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cart Logic Suite")
}
