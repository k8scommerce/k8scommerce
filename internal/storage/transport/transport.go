package transport

type Transport interface {
	Open(destinationPath, fileName, contentType string) error
	StreamPut(buffer []byte, partNumber int) error
	Close() error
}
