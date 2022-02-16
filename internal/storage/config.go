package storage

type UploadConfig struct {
	AllowedTypes AllowedTypes
	Storage      StorageConfig
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
	FileSystem       bool
	FileSystemConfig FileSystemConfig
	AWS              bool
	AwsConfig        AwsConfig
}

type FileSystemConfig struct {
	SavePath string
}

type AwsConfig struct {
	AwsRegion          string
	AwsAccessKeyId     string
	AwsSecretAccessKey string
	BucketName         string
	SavePath           string
}
