package asset

import (
	"path/filepath"
	"regexp"
	"strings"

	"github.com/k8scommerce/k8scommerce/internal/storage/config"
	"github.com/k8scommerce/k8scommerce/internal/storage/transport"
	"github.com/k8scommerce/k8scommerce/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type File struct {
	Name        string
	BaseName    string
	Kind        Kind
	ContentType string
	URL         string

	cfg              config.UploadConfig
	storageTransport transport.Transport
	destinationPath  string
}

func MustNewFile(name string, cfg config.UploadConfig) (*File, error) {
	f := &File{
		cfg:  cfg,
		Name: name,
	}

	if err := f.checkForSingleTransportEnabled(); err != nil {
		return nil, err
	}

	if err := f.setStoragetTransport(); err != nil {
		return nil, err
	}

	// set the folder path, i.e. /base/1/b/c/
	f.setDestinationPath()

	// set the asset public URL
	f.setURL()

	return f, nil
}

func (f *File) Open(contentType string) error {
	f.ContentType = contentType

	// open the stream
	if err := f.storageTransport.Open(f.destinationPath, f.BaseName, f.ContentType); err != nil {
		return err
	}
	return nil
}

// Write is a passthru function
// it sends the raw []byte chunks to the configured
// storage transport destination
func (f *File) Write(chunk []byte, partNumber int) error {
	// stream the content to the destination
	if chunk != nil {
		if err := f.storageTransport.StreamPut(chunk, partNumber); err != nil {
			return err
		}
	}
	return nil
}

// close the transport
func (f *File) Close() error {
	return f.storageTransport.Close()
}

// getter for UploadConfig
func (f *File) Config() config.UploadConfig {
	return f.cfg
}

// getter for destination path
func (f *File) GetDestinationPath() string {
	return f.destinationPath
}

// returns a storage object
// that adheres to the transport.Transport interface
func (f *File) GetStorageTransport() transport.Transport {
	return f.storageTransport
}

func (f *File) checkForSingleTransportEnabled() error {
	enabled := []string{}
	cnt := 0
	if f.cfg.StorageConfig.FileSystem {
		cnt++
		enabled = append(enabled, "FileSystem")
	}
	if f.cfg.StorageConfig.AWS {
		cnt++
		enabled = append(enabled, "AWS")
	}
	if f.cfg.StorageConfig.GCP {
		cnt++
		enabled = append(enabled, "GCP")
	}
	if f.cfg.StorageConfig.Azure {
		cnt++
		enabled = append(enabled, "Azure")
	}

	if cnt > 1 {
		return status.Errorf(codes.Internal, "only one transport can be enabled at a time. Enabled: %s", strings.Join(enabled, ", "))
	}
	return nil
}

// determine which transport we are configured to use
func (f *File) setStoragetTransport() (err error) {
	if f.cfg.StorageConfig.FileSystem {
		f.storageTransport, err = transport.MustNewFileSystemTransport(f.cfg.StorageConfig.FileSystemConfig)
		f.cfg.StorageConfig.BasePath = f.cfg.StorageConfig.FileSystemConfig.BasePath + "/" + f.cfg.StorageConfig.BasePath

		return err
	} else if f.cfg.StorageConfig.AWS {
		f.storageTransport, err = transport.MustNewAwsTransport(f.cfg.StorageConfig.AWSConfig)
		return err
	} else if f.cfg.StorageConfig.GCP {
		f.storageTransport, err = transport.MustNewGcpTransport(f.cfg.StorageConfig.GCPConfig)
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
	f.BaseName = filepath.Base(f.Name)

	// turn it into an md 5. this ensures that
	// if we have a filename with one letter or number we
	// don't have an issue creating the directory structure
	md5 := utils.StringToMD5(f.BaseName)

	r := []rune(md5)
	base := string(r[0:1])
	sub := string(r[1])
	subplus := string(r[2])

	sections := []string{
		f.cfg.StorageConfig.BasePath,
		base,
		sub,
		subplus,
	}

	for i := range sections {
		if val := strings.TrimSpace(sections[i]); val != "" {
			sections[i] = val
		}
	}

	f.destinationPath = "/" + strings.Join(sections, "/") + "/"
	if f.destinationPath[:1] == "/" {
		f.destinationPath = f.destinationPath[1:len(f.destinationPath)]
	}
	pattern := regexp.MustCompile(`\/+`)
	f.destinationPath = pattern.ReplaceAllString(f.destinationPath, "/")
}

func (f *File) setURL() {
	cdnURL := f.cfg.StorageConfig.CdnURL
	if cdnURL[len(cdnURL)-1:] == "/" {
		cdnURL = cdnURL[0 : len(cdnURL)-1]
	}
	f.URL = cdnURL + f.destinationPath + f.Name
}
