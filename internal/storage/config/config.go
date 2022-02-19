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
	Archive   Archive
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

type Archive struct {
	MaxUploadSize string
	MimeTypes     []string
}

type StorageConfig struct {
	SubDirectory     string
	FileSystem       bool
	FileSystemConfig FileSystemConfig
	AWS              bool
	AWSConfig        AWSConfig
	Azure            bool
	AzureConfig      AzureConfig
	GCP              bool
	GCPConfig        GCPConfig
}

type FileSystemConfig struct {
	BasePath string
}

type AWSConfig struct {
	Region          string
	AccessKeyId     string
	SecretAccessKey string
	S3Bucket        string
}

type AzureConfig struct{}

type GCPConfig struct{}
