package logic_test

import (
	"errors"
	"fmt"
	"io"
	"k8scommerce/internal/storage/asset"
	"k8scommerce/internal/storage/config"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ = Describe("UploadAssetLogic", func() {
	defer GinkgoRecover()

	var err error
	var testFilesPath = "../../../../../internal/storage/asset/testfiles"

	err = godotenv.Load("../../../../../.env")
	Expect(err).To(BeNil())
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var assetFile *asset.File

	getUploadConfig := func() *config.UploadConfig {
		return &config.UploadConfig{
			AllowedTypes: config.AllowedTypes{
				Images: config.Images{
					MaxUploadSize: "25MB",
					MimeTypes: []string{
						".jpg",
						".png",
						".gif",
						".jpeg",
						".ico",
					},
				},
				Documents: config.Documents{
					MaxUploadSize: "25MB",
					MimeTypes: []string{
						".pdf",
						".doc",
						".ppt",
						".odt",
						".xls",
						".psd",
					},
				},
				Audio: config.Audio{
					MaxUploadSize: "25MB",
					MimeTypes: []string{
						".mp3",
						".m4a",
						".ogg",
						".wav",
					},
				},
				Video: config.Video{
					MaxUploadSize: "25MB",
					MimeTypes: []string{
						".mp4",
						".mov",
						".avi",
						".mpg",
						".ogv",
						".3gp",
						".3g2",
					},
				},
			},
			StorageConfig: config.StorageConfig{
				SubDirectory: "",
				FileSystem:   true,
				FileSystemConfig: config.FileSystemConfig{
					BasePath: "",
				},
				AWS: false,
				AWSConfig: config.AWSConfig{
					Region:          "",
					AccessKeyId:     "",
					SecretAccessKey: "",
					S3Bucket:        "",
				},
			},
		}
	}

	getContentType := func(chunk []byte) string {
		if len(chunk) > 1 {
			return http.DetectContentType(chunk)
		}
		return ""
	}

	bufferAssetFile := func(assetFile *asset.File) error {
		const BufferSize = 5 * 1024 * 1024
		f, err := os.Open(assetFile.Name)
		if err != nil {
			return err
		}
		defer f.Close()

		buffer := make([]byte, BufferSize)

		partNumber := 1
		for {
			bytesRead, err := f.Read(buffer)
			if err != nil {
				if err != io.EOF {

					return err
				}
				break
			}

			if partNumber == 1 {
				contentType := getContentType(buffer)
				Expect(contentType).To(Not(BeNil()))
				err = assetFile.Open(contentType)
				Expect(err).To(BeNil())
			}

			if err := assetFile.Write(buffer[:bytesRead], partNumber); err != nil {
				return status.Error(codes.Internal, err.Error())
			}

			partNumber++
		}
		return nil
	}

	fileExists := func(path string) bool {
		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			return false
		}
		return true
	}

	fileSize := func(path string) int64 {
		f, err := os.Stat(path)
		if err != nil {
			return 0
		}
		return f.Size()
	}

	deleteTestUploadDirectory := func() error {
		return os.RemoveAll(testFilesPath + "/uploads")
	}

	Describe("SaveFile", func() {

		BeforeEach(func() {
			deleteTestUploadDirectory()
		})

		Describe("FileSystem", func() {

			// reset the config
			cfg := getUploadConfig()

			BeforeEach(func() {
				cfg.StorageConfig.SubDirectory = "uploads"

				cfg.StorageConfig.FileSystem = true
				cfg.StorageConfig.FileSystemConfig.BasePath = testFilesPath

				cfg.StorageConfig.AWS = false
				cfg.StorageConfig.GCP = false
				cfg.StorageConfig.Azure = false

				name := testFilesPath + "logo.png"
				assetFile, err = asset.MustNewFile(name, *cfg)
				Expect(err).To(BeNil())

				assetFile.Kind = asset.Image
			})

			It("should have all prerequisites", func() {
				Expect(assetFile.Kind).To(Not(BeNil()))
				Expect(assetFile.ContentType).To(Equal(""))

				Expect(assetFile.GetDestinationPath()).To(Not(BeNil()))
				Expect(assetFile.GetDestinationPath()).To(Equal(testFilesPath + "/uploads/1/b/b/"))

				Expect(assetFile.GetStorageTransport()).To(Not(BeNil()))
				xType := fmt.Sprintf("%T", assetFile.GetStorageTransport())
				Expect(xType).To(ContainSubstring("transport.fileSystemTransport"))
			})

			It("should save a file locally", func() {
				err = bufferAssetFile(assetFile)
				Expect(err).To(BeNil())

				err = assetFile.Close()
				Expect(err).To(BeNil())

				exists := fileExists(assetFile.Name)
				Expect(exists).To(BeTrue())

				size := fileSize(assetFile.Name)
				Expect(size).To(Not(Equal(0)))
			})

			It("should save an svg", func() {
				name := "./testfiles/logo.svg"
				assetFile, err = asset.MustNewFile(name, *cfg)
				Expect(err).To(BeNil())

				assetFile.Kind = asset.Image

				err = bufferAssetFile(assetFile)
				Expect(err).To(BeNil())

				err = assetFile.Close()
				Expect(err).To(BeNil())

				exists := fileExists(assetFile.Name)
				Expect(exists).To(BeTrue())

				size := fileSize(assetFile.Name)
				Expect(size).To(Not(Equal(0)))
			})
		})

		// Describe("Aws S3", func() {

		// 	// reset the config
		// 	cfg := getUploadConfig()

		// 	BeforeEach(func() {
		// 		cfg.StorageConfig.SubDirectory = "uploads"

		// 		cfg.StorageConfig.AWS = true
		// 		cfg.StorageConfig.AWSConfig.AccessKeyId = os.Getenv("AWS_ACCESS_KEY_ID")
		// 		cfg.StorageConfig.AWSConfig.SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
		// 		cfg.StorageConfig.AWSConfig.Region = os.Getenv("AWS_DEFAULT_REGION")
		// 		cfg.StorageConfig.AWSConfig.S3Bucket = os.Getenv("S3_BUCKET")

		// 		cfg.StorageConfig.FileSystem = false
		// 		cfg.StorageConfig.GCP = false
		// 		cfg.StorageConfig.Azure = false

		// 		name := "./testfiles/Pizigani_1367_Chart_10MB.jpeg"
		// 		assetFile, err = asset.MustNewFile(name, *cfg)
		// 		Expect(err).To(BeNil())

		// 		assetFile.Kind = asset.Image
		// 	})

		// 	It("should have all prerequisites", func() {
		// 		Expect(assetFile.Kind).To(Not(BeNil()))
		// 		Expect(assetFile.ContentType).To(Equal(""))

		// 		Expect(assetFile.GetDestinationPath()).To(Not(BeNil()))
		// 		Expect(assetFile.GetDestinationPath()).To(Equal("uploads/1/b/b/"))

		// 		Expect(assetFile.GetStorageTransport()).To(Not(BeNil()))
		// 		xType := fmt.Sprintf("%T", assetFile.GetStorageTransport())
		// 		Expect(xType).To(ContainSubstring("transport.awsTransport"))
		// 	})

		// 	It("should save a file on S3", func() {
		// 		err = bufferAssetFile(assetFile)
		// 		Expect(err).To(BeNil())

		// 		err = assetFile.Close()
		// 		Expect(err).To(BeNil())
		// 	})
		// })

	})
})
