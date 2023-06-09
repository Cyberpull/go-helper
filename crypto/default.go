package crypto

import (
	"os"

	_ "cyberpull.com/gotk/env"

	"cyberpull.com/gotk/errors"
)

func GetCipherKey(key ...string) (value string, err error) {
	if len(key) > 0 {
		value = key[0]
		return
	}

	value = os.Getenv("CIPHER_KEY")

	if value == "" {
		err = errors.New(`"CIPHER_KEY" is required`)
	}

	return
}
