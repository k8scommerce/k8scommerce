package storage

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type audioHandler struct {
}

func mustNewAudioHandler(file *File, c UploadConfig) StorageHandler {
	return &audioHandler{}
}

func (h *audioHandler) save() error {
	return status.Error(codes.Internal, "audio handler failed")
}
