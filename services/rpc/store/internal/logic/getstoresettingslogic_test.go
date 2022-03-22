package logic_test

import (
	"encoding/json"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"k8scommerce/internal/convert"
	"k8scommerce/internal/models"
	"k8scommerce/services/rpc/store/pb/store"
	"k8scommerce/services/rpc/store/storeclient"
)

var _ = Describe("GetStoreSettingsLogic", func() {
	getStoreConfig := func() *storeclient.StoreConfig {
		return &store.StoreConfig{
			Currency: &store.Currency{
				DefaultCurrency: "USD",
				SupportedCurrencies: []string{
					"USD",
					"CAN",
				},
			},
			Locale: &store.Locale{
				DefaultLocale: "en-US",
				SupportedLocales: []string{
					"en-US",
					"en-CA",
				},
			},
			Contact: &store.Contact{
				Phone: &store.Phone{
					CustomerSupport: "800-555-1212",
					Corportate:      "800-555-1212",
					Custom: map[string]string{
						"Sales": "800-555-1212 Ext 1234",
					},
				},
				Addresses: []*store.Address{
					{
						Name:          "Warehouse 1",
						Street:        "123 Any Street",
						AptSuite:      "100",
						City:          "Big Springs",
						StateProvince: "VA",
						PostalCode:    "12345",
						Country:       "US",
						IsDefault:     true,
					},
				},
			},
			Emails: &store.Emails{
				Default: &store.Email{
					Name:  "K8sCommerce",
					Email: "test@k8scommerce.com",
				},
				CustomerSupport: &store.Email{
					Name:  "K8sCommerce",
					Email: "support@k8scommerce.com",
				},
				CustomerCompletedOrder: &store.Email{
					Name:  "K8sCommerce",
					Email: "support@k8scommerce.com",
				},
				CustomerConfirmationEmail: &store.Email{
					Name:  "K8sCommerce",
					Email: "support@k8scommerce.com",
				},
				CustomerNewAccount: &store.Email{
					Name:  "K8sCommerce",
					Email: "support@k8scommerce.com",
				},
				CustomerNote: &store.Email{
					Name:  "K8sCommerce",
					Email: "support@k8scommerce.com",
				},
				CustomerOnHoldOrder: &store.Email{
					Name:  "K8sCommerce",
					Email: "support@k8scommerce.com",
				},
				CustomerPasswordChanged: &store.Email{
					Name:  "K8sCommerce",
					Email: "no-reply@k8scommerce.com",
				},
				CustomerProcessingOrder: &store.Email{
					Name:  "K8sCommerce",
					Email: "support@k8scommerce.com",
				},
				CustomerRefundedOrder: &store.Email{
					Name:  "K8sCommerce",
					Email: "support@k8scommerce.com",
				},
				CustomerResetPassword: &store.Email{
					Name:  "K8sCommerce",
					Email: "no-reply@k8scommerce.com",
				},
				CustomerSale: &store.Email{
					Name:  "K8sCommerce",
					Email: "support@k8scommerce.com",
				},
				AdminCancelledOrder: &store.Email{
					Name:  "K8sCommerce",
					Email: "no-reply@k8scommerce.com",
				},
				AdminFailedOrder: &store.Email{
					Name:  "K8sCommerce",
					Email: "no-reply@k8scommerce.com",
				},
				AdminNewOrder: &store.Email{
					Name:  "K8sCommerce",
					Email: "no-reply@k8scommerce.com",
				},
			},
		}
	}

	repo.GetRawDB().MustExec("TRUNCATE store_setting RESTART IDENTITY CASCADE;")

	Describe("Config", func() {

		It("should write the store settings to the db", func() {
			config, err := json.Marshal(getStoreConfig())
			Expect(err).To(BeNil())
			// convert settings
			setting := &models.StoreSetting{
				StoreID: 1,
				Config:  config,
			}

			err = repo.StoreSetting().Create(setting)
			Expect(err).To(BeNil())

			// ModelStoreSettingToProtoStoreSetting
		})

		It("should get and unmarshal store settings", func() {
			setting, err := repo.StoreSetting().GetStoreSettingById(1)
			Expect(err).To(BeNil())
			// fmt.Println(string(setting.Config))

			// storeConfig := &store.StoreConfig{}
			// err = json.Unmarshal(setting.Config, storeConfig)
			// Expect(err).To(BeNil())
			// fmt.Println(storeConfig)

			s := &store.StoreSetting{}
			convert.ModelStoreSettingToProtoStoreSetting(setting, s)
			fmt.Println(s)
		})
	})
})
