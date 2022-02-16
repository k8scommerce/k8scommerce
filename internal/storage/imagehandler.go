package storage

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type imageHandler struct {
}

func mustNewImageHandler(file *File, c UploadConfig) StorageHandler {
	return &imageHandler{}
}

func (h *imageHandler) save() error {
	return status.Error(codes.Internal, "image handler failed")
}
