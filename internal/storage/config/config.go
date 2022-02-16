package config

type UploadConfig struct {
	AllowedTypes  AllowedTypes
	StorageConfig StorageConfig
}

type AllowedTypes struct {
	Images    Images
	Documents Documents
	Audio     Audio
	Video     Video
}

type Images struct {
	MaxUploadSize string
	MimeTypes     []string
}

type Documents struct {
	MaxUploadSize string
	MimeTypes     []string
}

type Audio struct {
	MaxUploadSize string
	MimeTypes     []string
}

type Video struct {
	MaxUploadSize string
	MimeTypes     []string
}

type StorageConfig struct {
	SubDirectory     string
	FileSystem       bool
	FileSystemConfig FileSystemConfig
	AWS              bool
	AwsConfig        AwsConfig
	Azure            bool
	AzureConfig      AzureConfig
	Gcp              bool
	GcpConfig        GcpConfig
}

type FileSystemConfig struct {
	BasePath string
}

type AwsConfig struct {
	AwsRegion          string
	AwsAccessKeyId     string
	AwsSecretAccessKey string
	AwsBucket          string
}

type AzureConfig struct{}

type GcpConfig struct{}
