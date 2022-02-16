package storage

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StorageHandler interface {
	save() error
}

func SaveFile(file *File, c UploadConfig) error {
	handler, err := mustNewFileHandler(file, c)
	if err != nil {
		return err
	}

	if err := handler.save(); err != nil {
		return err
	}

	return nil
}

func mustNewFileHandler(file *File, config UploadConfig) (*fileHandler, error) {
	fileHandler := &fileHandler{
		file:   file,
		config: config,
	}

	err := fileHandler.setStoragetHandler()
	if err != nil {
		return nil, err
	}
	return fileHandler, nil
}

type fileHandler struct {
	file           *File
	config         UploadConfig
	storageHandler StorageHandler
}

func (h *fileHandler) setStoragetHandler() error {
	switch h.file.AssetKind {
	case AssetKind_Image:
		h.storageHandler = mustNewImageHandler(h.file, h.config)
		return nil
	case AssetKind_Document:
		h.storageHandler = mustNewDocumentHandler(h.file, h.config)
		return nil
	case AssetKind_Audio:
		h.storageHandler = mustNewAudioHandler(h.file, h.config)
		return nil
	case AssetKind_Video:
		h.storageHandler = mustNewVideoHandler(h.file, h.config)
		return nil
	}
	return status.Error(codes.Internal, "unknown storage handler")
}

func (h *fileHandler) save() error {
	return nil
}
