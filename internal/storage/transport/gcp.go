package transport

import "k8scommerce/internal/storage/config"

func MustNewGcpTransport(cfg config.GcpConfig) (Transport, error) {
	t := &gcpTransport{}
	err := t.getSession()
	return t, err
}

type gcpTransport struct {
}

func (t *gcpTransport) Open(destinationPath, fileName string) error {
	return nil
}

func (t *gcpTransport) StreamPut(buffer []byte) error {
	return nil
}

func (t *gcpTransport) Close() error {
	return nil
}

func (t *gcpTransport) getSession() error {
	return nil
}
