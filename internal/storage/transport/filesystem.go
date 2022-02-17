package transport

import (
	"k8scommerce/internal/storage/config"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MustNewFileSystemTransport(cfg config.FileSystemConfig) (Transport, error) {
	t := &fileSystemTransport{}
	err := t.getSession()
	return t, err
}

type fileSystemTransport struct {
	fileHandle *os.File
}

func (t *fileSystemTransport) Open(destinationPath, fileName, contentType string) error {
	if err := os.MkdirAll(destinationPath, os.ModePerm); err != nil {
		return status.Errorf(codes.Internal, "cannot create directory path %s", destinationPath)
	}

	f, err := os.Create(destinationPath + fileName)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot open %s for writing", destinationPath+fileName)
	}

	t.fileHandle = f
	return nil
}

func (t *fileSystemTransport) StreamPut(buffer []byte, partNumber int) error {
	n, err := t.fileHandle.Write(buffer)
	if err != nil {
		return status.Errorf(codes.Internal, "cannot write to file handle: %s", err.Error())
	}
	if len(buffer) != n {
		return status.Errorf(codes.Internal, "byte written not same length as bytes given")
	}

	return nil
}

func (t *fileSystemTransport) Close() error {
	err := t.fileHandle.Close()
	if err != nil {
		return status.Errorf(codes.Internal, "cannot close file handler: %s", err.Error())
	}
	return nil
}

func (t *fileSystemTransport) getSession() error {
	return nil
}
