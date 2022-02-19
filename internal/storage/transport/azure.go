package transport

import "k8scommerce/internal/storage/config"

func MustNewAzureTransport(cfg config.AzureConfig) (Transport, error) {
	t := &azureTransport{}
	err := t.getSession()
	return t, err
}

type azureTransport struct {
}

func (t *azureTransport) Open(destinationPath, fileName, contentType string) error {
	return nil
}

func (t *azureTransport) StreamPut(buffer []byte, partNumber int) error {
	return nil
}

func (t *azureTransport) Close() error {
	return nil
}

func (t *azureTransport) getSession() error {
	return nil
}
