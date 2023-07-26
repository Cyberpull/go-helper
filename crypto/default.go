package crypto

import (
	"os"

	_ "cyberpull.com/gotk/v2/env"

	"cyberpull.com/gotk/v2/errors"
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
