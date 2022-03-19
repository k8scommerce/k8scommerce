package encryption_test

import (
	"k8scommerce/internal/encryption"
	"k8scommerce/internal/encryption/config"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Encryption", func() {

	encrypter := encryption.NewEncrypter(&config.EncryptionConfig{
		Secret: `Fg-?8jd3M#hEk-mTMS-9wpfRkD!SEbu5`,
		Token:  `aNmja9ZB4Y+BH#^j`,
	})

	const rawString = `test@k8scommerce.com|1`

	It("should create a encrypt and decrypt a string", func() {
		encText, err := encrypter.Encrypt(rawString)
		Expect(err).To(BeNil())

		decText, err := encrypter.Decrypt(encText)
		Expect(err).To(BeNil())

		Expect(decText).ToNot(BeNil())
		Expect(decText).To(Equal(rawString))

	})
})
