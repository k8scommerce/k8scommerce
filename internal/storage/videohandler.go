package storage

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type videoHandler struct {
}

func mustNewVideoHandler(file *File, c UploadConfig) StorageHandler {
	return &videoHandler{}
}

func (h *videoHandler) save() error {
	return status.Error(codes.Internal, "video handler failed")
}
