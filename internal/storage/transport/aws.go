package transport

import (
	"k8scommerce/internal/storage/config"
)

func MustNewAwsTransport(cfg config.AwsConfig) (Transport, error) {
	t := &awsTransport{
		cfg: cfg,
	}
	err := t.getSession()
	return t, err
}

type awsTransport struct {
	cfg config.AwsConfig
}

func (t *awsTransport) Open(destinationPath, fileName string) error {
	return nil
}

func (t *awsTransport) StreamPut(buffer []byte) error {
	// reader, writer := io.Pipe()

	// go func() {
	// 	gw := gzip.NewWriter(writer)
	// 	io.Copy(gw, file)
	// 	file.Close()
	// 	gw.Close()
	// 	writer.Close()
	// }()

	// // Initialize a session in us-west-2 that the SDK will use to load
	// // credentials from the shared credentials file ~/.aws/credentials.
	// sess, err := session.NewSession(&aws.Config{
	// 	Region: aws.String(t.cfg.AwsRegion)},
	// )

	// uploader := s3manager.NewUploader(sess)

	// result, err := uploader.Upload(&s3manager.UploadInput{
	// 	Body:   reader,
	// 	Bucket: aws.String(t.cfg.AwsBucket),
	// 	Key:    aws.String(t.cfg.AwsAccessKeyId),
	// })
	// if err != nil {
	// 	log.Fatalln("Failed to upload", err)
	// }

	// log.Println("Successfully uploaded to", result.Location)
	return nil
}

func (t *awsTransport) Close() error {
	return nil
}

func (t *awsTransport) getSession() error {
	return nil
}
