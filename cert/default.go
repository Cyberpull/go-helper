package cert

import (
	"crypto/tls"
	"os"

	_ "cyberpull.com/gotk/v3/env"

	"cyberpull.com/gotk/v3/errors"
)

var (
	certFile string
	keyFile  string
)

func init() {
	certFile = os.Getenv("CERT_CRT_FILE")
	keyFile = os.Getenv("CERT_KEY_FILE")
}

func Validate(opts ...Options) (err error) {
	return getOpts(&opts).validate()
}

func IsEnabled(opts ...Options) bool {
	return getOpts(&opts).isEnabled()
}

func GetTLSConfig(opts ...Options) (config *tls.Config, err error) {
	tmpConfig := &tls.Config{}

	if IsEnabled(opts...) {
		config.Certificates, err = GetCertificates(opts...)

		if err != nil {
			return
		}

		err = SanitizeTlsConfig(config)

		if err != nil {
			return
		}

		config = tmpConfig
	}

	return
}

func SanitizeTlsConfig(config *tls.Config) (err error) {
	if config == nil {
		err = errors.New("No config found.")
		return
	}

	return
}

func GetCertificates(opts ...Options) (value []tls.Certificate, err error) {
	opt := getOpts(&opts)

	value = make([]tls.Certificate, 0)

	if err = opt.validate(); err != nil {
		return
	}

	cfile, kfile := opt.getKeyPairFiles()
	cert, err := tls.LoadX509KeyPair(cfile, kfile)

	if err != nil {
		return
	}

	value = append(value, cert)

	return
}
