package storage

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type documentHandler struct {
}

func mustNewDocumentHandler(file *File, c UploadConfig) StorageHandler {
	return &documentHandler{}
}

func (h *documentHandler) save() error {
	return status.Error(codes.Internal, "document handler failed")
}
