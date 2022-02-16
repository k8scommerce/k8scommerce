package asset

import (
	"bytes"
	"k8scommerce/internal/storage/config"
	"k8scommerce/internal/storage/transport"
	"k8scommerce/internal/utils"
	"net/http"
	"path/filepath"
	"regexp"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Kind int

const (
	Image    Kind = 0
	Document Kind = 1
	Audio    Kind = 2
	Video    Kind = 3
)

type File struct {
	Buffer      *bytes.Buffer
	Name        string
	Kind        Kind
	ContentType string

	cfg              config.UploadConfig
	storageTransport transport.Transport
	destinationPath  string
	baseName         string
}

func MustNewFile(name string, cfg config.UploadConfig) (*File, error) {
	f := &File{
		cfg:    cfg,
		Name:   name,
		Buffer: &bytes.Buffer{},
	}

	if err := f.setStoragetTransport(); err != nil {
		return nil, err
	}

	// set the folder path, i.e. /base/1/b/c/
	f.setDestinationPath()

	return f, nil
}

func (f *File) Open() error {
	// open the stream
	if err := f.storageTransport.Open(f.destinationPath, f.baseName); err != nil {
		return err
	}
	return nil
}

func (f *File) Write(chunk []byte) error {
	_, err := f.Buffer.Write(chunk)
	if err != nil {
		return err
	}

	// set the mime type
	// this will only occur on the first chunk
	if f.ContentType == "" {
		f.setContentType()
	}

	// stream the content to the destination
	if err := f.storageTransport.StreamPut(chunk); err != nil {
		return err
	}

	return nil
}

func (f *File) Close() error {
	return nil
}

func (f *File) Config() config.UploadConfig {
	return f.cfg
}

func (f *File) GetDestinationPath() string {
	return f.destinationPath
}

func (f *File) GetStorageTransport() transport.Transport {
	return f.storageTransport
}

func (f *File) setContentType() error {
	if f.Buffer.Len() > 1 {
		fileType := http.DetectContentType(f.Buffer.Bytes())
		f.ContentType = fileType
		return nil
	}
	return status.Error(codes.Internal, "mime type cannot be detected. file buffer length is zero")
}

func (f *File) setStoragetTransport() (err error) {
	if f.cfg.StorageConfig.FileSystem {
		f.storageTransport, err = transport.MustNewFileSystemTransport(f.cfg.StorageConfig.FileSystemConfig)
		f.cfg.StorageConfig.SubDirectory = f.cfg.StorageConfig.FileSystemConfig.BasePath + "/" + f.cfg.StorageConfig.SubDirectory
		return err
	} else if f.cfg.StorageConfig.AWS {
		f.storageTransport, err = transport.MustNewAwsTransport(f.cfg.StorageConfig.AwsConfig)
		return err
	} else if f.cfg.StorageConfig.Gcp {
		f.storageTransport, err = transport.MustNewGcpTransport(f.cfg.StorageConfig.GcpConfig)
		return err
	} else if f.cfg.StorageConfig.Azure {
		f.storageTransport, err = transport.MustNewAzureTransport(f.cfg.StorageConfig.AzureConfig)
		return err
	}
	return status.Error(codes.Internal, "unknown storage transport")
}

// create a subdirectory structure for the file
// md5 the file: 1bb87d41d15fe27b500a4bfcde01bb0e
// returned directory would be this:
// 1/b/b
func (f *File) setDestinationPath() {
	// get the raw filename
	f.baseName = filepath.Base(f.Name)

	// turn it into an md 5. this ensures that
	// if we have a filename with one letter or number we
	// don't have an issue creating the directory structure
	md5 := utils.StringToMD5(f.baseName)

	r := []rune(md5)
	base := string(r[0:1])
	sub := string(r[1])
	subplus := string(r[2])

	sections := []string{
		f.cfg.StorageConfig.SubDirectory,
		base,
		sub,
		subplus,
	}

	for i := range sections {
		sections[i] = strings.TrimSpace(sections[i])
	}

	pattern := regexp.MustCompile(`\/+`)
	f.destinationPath = pattern.ReplaceAllString(strings.Join(sections, "/")+"/", "/")
}

// func (f *File) setStoragetHandler() error {
// 	switch h.File.Kind {
// 	case asset.Image:
// 		h.StorageHandler = handler.NewImageHandler(h.File)
// 		return nil
// 	case asset.Document:
// 		h.StorageHandler = handler.NewDocumentHandler(h.File)
// 		return nil
// 	case asset.Audio:
// 		h.StorageHandler = handler.NewAudioHandler(h.File)
// 		return nil
// 	case asset.Video:
// 		h.StorageHandler = handler.NewVideoHandler(h.File)
// 		return nil
// 	}

// 	return status.Error(codes.Internal, "unknown storage handler")
// }
