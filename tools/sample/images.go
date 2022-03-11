package main

import (
	"database/sql"
	"fmt"
	"io"
	"k8scommerce/internal/events"
	eventsconfig "k8scommerce/internal/events/config"
	"k8scommerce/internal/events/eventkey"
	"k8scommerce/internal/models"
	"k8scommerce/internal/storage/asset"
	storageconfig "k8scommerce/internal/storage/config"
	"k8scommerce/internal/workerpool"
	"net/http"
	"os"
	"path/filepath"

	"github.com/jaswdr/faker"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	totalWorkers int = 20
	wp           workerpool.WorkerPool
)

func init() {
	wp = workerpool.NewWorkerPool(totalWorkers)
	wp.StartWorkers()
}

func createImages() {

	storageConfig := getUploadConfig()
	storageConfig.StorageConfig.CdnURL = os.Getenv("CDN_URL")
	storageConfig.StorageConfig.BasePath = os.Getenv("BASE_PATH")
	storageConfig.StorageConfig.AWS = true
	storageConfig.StorageConfig.AWSConfig.AccessKeyId = os.Getenv("AWS_ACCESS_KEY_ID")
	storageConfig.StorageConfig.AWSConfig.SecretAccessKey = os.Getenv("AWS_SECRET_ACCESS_KEY")
	storageConfig.StorageConfig.AWSConfig.Region = os.Getenv("AWS_DEFAULT_REGION")
	storageConfig.StorageConfig.AWSConfig.S3Bucket = os.Getenv("S3_BUCKET")
	storageConfig.StorageConfig.FileSystem = false
	storageConfig.StorageConfig.GCP = false
	storageConfig.StorageConfig.Azure = false

	eventsConfig := getEventsConfig()
	eventManager := events.NewEventManager(eventsConfig)

	faker := faker.New()
	// get products
	storeId := int64(1)
	res, err := repo.Product().GetAllProducts(storeId, 0, 1000, "")
	if err != nil {
		fmt.Printf("error fetching products: %s", err.Error())
		return
	}
	if res == nil {
		fmt.Println("no results while fetching products")
		return
	}

	var totalTasks int64
	row := repo.GetRawDB().QueryRow("select count(*) as totalTasks from variant")
	row.Scan(&totalTasks)
	doneC := make(chan bool, totalTasks)
	fmt.Println("totalTasks", totalTasks)

	for _, prod := range res.Results {
		prod := prod // https://stackoverflow.com/a/10117080

		wp.AddTask(func() {
			pRes, err := repo.Product().GetProductById(prod.Product.ID)
			if err != nil {
				fmt.Println("ERROR", err.Error())
				doneC <- true
				return
			}

			sortOrder := int64(1)
			for _, variant := range pRes.Variants {
				variant := variant
				tx, err := repo.Begin()
				if err != nil {
					fmt.Printf("failed to start a transaction: %s", err.Error())
				}

				image := faker.LoremFlickr().Image(1000, 1000, []string{}, "", false)

				fmt.Printf("uploading image: %s\n", image.Name())

				baseName := filepath.Base(image.Name())

				file, err := asset.MustNewFile(baseName, *storageConfig)
				if err != nil {
					fmt.Printf("file creation error: %s", err.Error())
				}

				file.Kind = asset.Kind("image")

				if err := bufferAssetFile(image.Name(), file); err != nil {
					fmt.Println(err.Error())
					// return
				}

				modelAsset := &models.Asset{
					StoreID:     storeId,
					ProductID:   variant.ProductID,
					VariantID:   variant.ID,
					Kind:        file.Kind.Int(),
					Name:        file.BaseName,
					DisplayName: sql.NullString{Valid: true, String: file.BaseName},
					ContentType: file.ContentType,
					URL:         file.URL,
					SortOrder:   sql.NullInt64{Valid: true, Int64: sortOrder},
					Sizes:       []byte("{}"),
				}

				sortOrder++

				// save the asset to the database
				if err := repo.Asset().CreateTx(modelAsset, tx); err != nil {
					repo.Rollback(tx)
					fmt.Printf("saving asset to database failed: %s", err.Error())
				}

				// commit the transaction
				if err := repo.Commit(tx); err != nil {
					fmt.Printf("saving asset db commit failed: %s", err.Error())
				}

				if bytes, err := eventkey.CatalogImageUploaded.Marshal(modelAsset); err != nil {
					logx.Infof("%d: marshaling event %s failed: %s", codes.Internal, eventkey.CatalogImageUploaded, err.Error())
				} else {
					// publish event
					err = eventManager.Publish(eventkey.CatalogImageUploaded.AsKey(), bytes)
					if err != nil {
						logx.Infof("%d: publishing event %s failed: %s", codes.Internal, eventkey.CatalogImageUploaded, err.Error())
					}
				}

				// play cleanup
				if err := os.Remove(image.Name()); err != nil {
					fmt.Printf("could not remove image: %s", image.Name())
				}

				// time.Sleep(100 * time.Microsecond)
				doneC <- true
			}
		})

	}

	for i := int64(1); i <= totalTasks; i++ {
		if <-doneC {
			fmt.Printf("finished processing %d images", i)
		}
	}

}

func getContentType(chunk []byte) string {
	if len(chunk) > 1 {
		return http.DetectContentType(chunk)
	}
	return ""
}

func bufferAssetFile(fileWithPath string, assetFile *asset.File) error {
	const BufferSize = 5 * 1024 * 1024

	f, err := os.Open(fileWithPath)
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
			_ = assetFile.Open(contentType)
		}

		if err := assetFile.Write(buffer[:bytesRead], partNumber); err != nil {
			return status.Error(codes.Internal, err.Error())
		}

		partNumber++
	}
	return assetFile.Close()
}

func getEventsConfig() *eventsconfig.EventsConfig {
	return &eventsconfig.EventsConfig{
		Enabled:  true,
		RabbitMQ: true,
		RabbitMQConfig: eventsconfig.RabbitMQConfig{
			Url: os.Getenv("RABBITMQ_URL"),
		},
	}
}

func getUploadConfig() *storageconfig.UploadConfig {
	return &storageconfig.UploadConfig{
		AllowedTypes: storageconfig.AllowedTypes{
			Images: storageconfig.Images{
				MaxUploadSize: "25MB",
				MimeTypes: []string{
					".jpg",
					".png",
					".gif",
					".jpeg",
					".ico",
				},
			},
			Documents: storageconfig.Documents{
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
			Audio: storageconfig.Audio{
				MaxUploadSize: "25MB",
				MimeTypes: []string{
					".mp3",
					".m4a",
					".ogg",
					".wav",
				},
			},
			Video: storageconfig.Video{
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
		StorageConfig: storageconfig.StorageConfig{
			BasePath:   "",
			FileSystem: false,
			FileSystemConfig: storageconfig.FileSystemConfig{
				BasePath: "",
			},
			AWS: true,
			AWSConfig: storageconfig.AWSConfig{
				Region:          "",
				AccessKeyId:     "",
				SecretAccessKey: "",
				S3Bucket:        "",
			},
		},
	}
}
