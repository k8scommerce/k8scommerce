package repos_test

import (
	"database/sql"
	"fmt"
	"k8scommerce/internal/models"
	"k8scommerce/internal/repos"
	"log"
	"os"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Asset", func() {
	defer GinkgoRecover()

	var err error

	err = godotenv.Load("../../.env")
	Expect(err).To(BeNil())
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	getPostgresConfig := func() *repos.PostgresConfig {
		return &repos.PostgresConfig{
			DataSourceName: os.Getenv("POSTGRES_DSN"),
		}
	}

	repo := repos.NewRepo(getPostgresConfig())

	asset := &models.Asset{
		StoreID:     1,
		ProductID:   1,
		VariantID:   1,
		Name:        "Pizigani_1367_Chart_10MB.jpeg",
		DisplayName: sql.NullString{String: "Pizigani 1367 Chart 10MB", Valid: true},
		Kind:        "image",
		ContentType: "image/jpeg",
		URL:         "https://k8scommerce.s3.us-west-1.amazonaws.com/uploads/d/b/e/Pizigani_1367_Chart_10MB.jpeg",
		SortOrder:   sql.NullInt64{Int64: 100, Valid: true},
		Sizes:       []byte("[]"),
	}

	Describe("Create", func() {

		BeforeEach(func() {
			_, err := repo.GetRawDB().Exec("delete from asset where name = $1", asset.Name)
			Expect(err).To(BeNil())
		})

		It("should create an asset using", func() {
			err = repo.Asset().Create(asset)
			Expect(err).To(BeNil())
		})
	})

	FDescribe("GetById", func() {
		BeforeEach(func() {
			_, err := repo.GetRawDB().Exec("delete from asset where name = $1", asset.Name)
			Expect(err).To(BeNil())
			err = repo.Asset().Create(asset)
			Expect(err).To(BeNil())
		})

		It("should get an asset by id", func() {
			fmt.Println("ASSET ID: ", asset.ID)
			_, err := repo.Asset().GetAssetById(asset.ID)
			Expect(err).To(BeNil())
			// Expect(res).ToNot(BeNil())
		})
	})

})
