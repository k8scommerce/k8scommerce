package humanizer_test

import (
	"math"

	"github.com/k8scommerce/k8scommerce/internal/utils/humanizer"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Helpers", func() {
	defer GinkgoRecover()

	// expectNilErr := func(err error) {
	// 	Expect(err).To(BeNil())
	// }

	// expectErr := func(err error) {
	// 	Expect(err).To(Not(BeNil()))
	// }

	Describe("Size", func() {

		It("should convert from int64 to human", func() {
			sizes := map[int64]string{
				999:           "999 B",
				1000:          "1.0 kB",
				1023:          "1.0 kB",
				1024:          "1.0 kB",
				1000000:       "1.0 MB",
				987654321:     "987.7 MB",
				1e+9:          "1.0 GB",
				math.MaxInt64: "9.2 EB",
			}

			for intSize, testValue := range sizes {
				human := humanizer.SizeToHuman(intSize)
				Expect(human).To(Equal(testValue))
			}
		})

		It("should convert from human to int64", func() {
			sizes := map[int64]string{
				1:                "1.0 B",
				100:              "100.0 B",
				1000:             "1 kB",
				5000:             "5 kB",
				10000:            "10 kB",
				1000000:          "1 MB",
				5000000:          "5 MB",
				10000000:         "10 MB",
				20000000:         "20MB",
				25000000:         "25 MB",
				100000000:        "100 MB",
				500000000:        "500 MB",
				1000000000:       "1 GB",
				5000000000:       "5 GB",
				10000000000:      "10 GB",
				100000000000:     "100 GB",
				1000000000000:    "1 TB",
				1000000000000000: "1 PB",
			}

			for testValue, strSize := range sizes {
				asBytes := humanizer.HumanToSize(strSize)
				Expect(asBytes).To(Equal(testValue))
			}
		})
	})
})
