package email

import (
	"crypto/tls"
	"time"

	"github.com/k8scommerce/k8scommerce/services/consumer/email/internal/config"

	mail "github.com/xhit/go-simple-mail/v2"
)

func NewEmailClient(c config.EmailConfig) EmailClient {
	return &emailClient{
		config: c,
	}
}

type EmailClient interface {
	Connect() *emailClient
	GetSMTPClient() *mail.SMTPClient
}

type emailClient struct {
	config     config.EmailConfig
	smtpClient *mail.SMTPClient
}

func (c *emailClient) Connect() *emailClient {
	client := mail.NewSMTPClient()

	// SMTP Server
	client.Host = c.config.Host
	client.Port = c.config.Port
	client.Username = c.config.Username
	client.Password = c.config.Password

	if c.config.Encryption == "SSL/TLS" {
		client.Encryption = mail.EncryptionSSLTLS
	} else if c.config.Encryption == "STARTTLS" {
		client.Encryption = mail.EncryptionSTARTTLS
	}

	// Since v2.3.0 you can specified authentication type:
	// - PLAIN (default)
	// - LOGIN
	// - CRAM-MD5
	// - None
	// client.Authentication = mail.AuthPlain

	// Variable to keep alive connection
	client.KeepAlive = true

	// Timeout for connect to SMTP Server
	// client.ConnectTimeout = 10 * time.Second

	// Timeout for send the data and wait respond
	client.SendTimeout = 10 * time.Second

	// Set TLSConfig to provide custom TLS configuration. For example,
	// to skip TLS verification (useful for testing):
	client.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// SMTP client
	smtpClient, err := client.Connect()
	c.smtpClient = smtpClient

	if err != nil {
		panic(err)
	}

	return c
}

func (c *emailClient) GetSMTPClient() *mail.SMTPClient {
	return c.smtpClient
}

// 62335d557954222b95161cbd/
//aHR0cHM6Ly9ibGF6ZXR2LmNvbS9yZXNldC1wdy8zbk1OVnh3YmpLNmo1MVNx/5fc867892e7e854400611a65Bbc1db7e1
