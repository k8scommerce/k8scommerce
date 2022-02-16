package buildsql_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestBuildSql(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "BuildSql Suite")
}
