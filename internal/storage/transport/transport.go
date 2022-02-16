package transport

type Transport interface {
	Open(destinationPath, fileName string) error
	StreamPut(buffer []byte) error
	Close() error
}
