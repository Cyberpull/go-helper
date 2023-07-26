package cert

import "cyberpull.com/gotk/v3/errors"

type Options struct {
	CertFile string
	KeyFile  string
}

func (o Options) getKeyPairFiles() (cert string, key string) {
	if ok := o.CertFile != "" && o.KeyFile != ""; ok {
		return o.CertFile, o.KeyFile
	}

	return certFile, keyFile
}

func (o Options) isEnabled() bool {
	cert, key := o.getKeyPairFiles()

	if ok := cert != "" && key != ""; ok {
		return true
	}

	return false
}

func (o Options) validate() (err error) {
	cfile, kfile := o.getKeyPairFiles()

	if cfile == "" || kfile == "" {
		err = errors.New(`"CertFile" and "KeyFile" are required`)
	}

	return
}

// =========================

func getOpts(opts *[]Options) Options {
	if len(*opts) == 0 {
		*opts = append(*opts, Options{})
	}

	return (*opts)[0]
}
