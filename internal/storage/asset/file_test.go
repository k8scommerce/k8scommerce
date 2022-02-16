package asset_test

import (
	"errors"
	"fmt"
	"io"
	"k8scommerce/internal/storage/asset"
	"k8scommerce/internal/storage/config"
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ = Describe("Asset", func() {
	defer GinkgoRecover()

	var err error
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
				AwsConfig: config.AwsConfig{
					AwsRegion:          "",
					AwsAccessKeyId:     "",
					AwsSecretAccessKey: "",
					AwsBucket:          "k8scommerce",
				},
			},
		}
	}

	bufferAssetFile := func(assetFile *asset.File) error {
		const BufferSize = 4 * 1024
		f, err := os.Open(assetFile.Name)
		if err != nil {
			return err
		}
		defer f.Close()

		buffer := make([]byte, BufferSize)

		for {
			bytesRead, err := f.Read(buffer)
			if err != nil {
				if err != io.EOF {

					return err
				}
				break
			}
			if err := assetFile.Write(buffer[:bytesRead]); err != nil {
				return status.Error(codes.Internal, err.Error())
			}
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
		return os.RemoveAll("./testfiles/uploads")
	}

	Describe("SaveFile", func() {

		BeforeEach(func() {
			deleteTestUploadDirectory()
		})

		Describe("FileSystem", func() {

			// reset the config
			cfg := getUploadConfig()

			BeforeEach(func() {
				cfg.StorageConfig.FileSystem = true
				cfg.StorageConfig.FileSystemConfig.BasePath = "testfiles"
				cfg.StorageConfig.AWS = false
				cfg.StorageConfig.SubDirectory = "uploads"

				name := "./testfiles/logo.png"
				assetFile, err = asset.MustNewFile(name, *cfg)
				Expect(err).To(BeNil())

				assetFile.Kind = asset.Image
			})

			It("should have all prerequisites", func() {
				Expect(assetFile.Kind).To(Not(BeNil()))
				Expect(assetFile.ContentType).To(Equal(""))

				Expect(assetFile.GetDestinationPath()).To(Not(BeNil()))
				Expect(assetFile.GetDestinationPath()).To(Equal("testfiles/uploads/1/b/b/"))

				Expect(assetFile.GetStorageTransport()).To(Not(BeNil()))
				xType := fmt.Sprintf("%T", assetFile.GetStorageTransport())
				Expect(xType).To(ContainSubstring("transport.fileSystemTransport"))
			})

			It("should save a file locally", func() {
				err = assetFile.Open()
				Expect(err).To(BeNil())

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

				err = assetFile.Open()
				Expect(err).To(BeNil())

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

	})
})
