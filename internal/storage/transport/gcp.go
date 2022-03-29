package transport

import "github.com/k8scommerce/k8scommerce/internal/storage/config"

func MustNewGcpTransport(cfg config.GCPConfig) (Transport, error) {
	t := &gcpTransport{}
	err := t.getSession()
	return t, err
}

type gcpTransport struct {
}

func (t *gcpTransport) Open(destinationPath, fileName, contentType string) error {
	return nil
}

func (t *gcpTransport) StreamPut(buffer []byte, partNumber int) error {
	return nil
}

func (t *gcpTransport) Close() error {
	return nil
}

func (t *gcpTransport) getSession() error {
	return nil
}
