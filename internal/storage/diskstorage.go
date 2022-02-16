package storage

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/google/uuid"
)

// DiskStorage stores image on disk, and its info on memory
type DiskStorage struct {
	mutex       sync.RWMutex
	imageFolder string
}

// NewDiskStorage returns a new DiskStorage
func NewDiskStorage(imageFolder string) *DiskStorage {
	return &DiskStorage{
		imageFolder: imageFolder,
	}
}

// Save adds a new image to a laptop
func (store *DiskStorage) Save(
	laptopID string,
	imageType string,
	imageData bytes.Buffer,
) (string, error) {
	imageID, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("cannot generate image id: %w", err)
	}

	imagePath := fmt.Sprintf("%s/%s%s", store.imageFolder, imageID, imageType)

	file, err := os.Create(imagePath)
	if err != nil {
		return "", fmt.Errorf("cannot create image file: %w", err)
	}

	_, err = imageData.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("cannot write image to file: %w", err)
	}

	store.mutex.Lock()
	defer store.mutex.Unlock()

	// store.images[imageID.String()] = &ImageInfo{
	// 	LaptopID: laptopID,
	// 	Type:     imageType,
	// 	Path:     imagePath,
	// }

	return imageID.String(), nil
}
