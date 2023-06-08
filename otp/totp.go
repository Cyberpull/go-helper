package otp

import (
	"fmt"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

type TOTP interface {
	Generate(ttime time.Time, digits Digits, seconds ...uint) (string, error)
	Validate(code string, digits Digits, seconds ...uint) (bool, error)
	ToURL() string
	QRCode() QRCode
}

type pTOTP struct {
	opts *Options
}

func (t *pTOTP) Generate(ttime time.Time, digits Digits, seconds ...uint) (string, error) {
	if len(seconds) == 0 {
		seconds = append(seconds, 30)
	}

	return totp.GenerateCodeCustom(t.opts.Secret, ttime, totp.ValidateOpts{
		Period:    seconds[0],
		Skew:      1,
		Digits:    otp.Digits(digits),
		Algorithm: otp.AlgorithmSHA1,
	})
}

func (t *pTOTP) Validate(code string, digits Digits, seconds ...uint) (bool, error) {
	if len(seconds) == 0 {
		seconds = append(seconds, 30)
	}

	return totp.ValidateCustom(
		code,
		t.opts.Secret,
		time.Now().UTC(),
		totp.ValidateOpts{
			Period:    seconds[0],
			Skew:      1,
			Digits:    otp.Digits(digits),
			Algorithm: otp.AlgorithmSHA1,
		},
	)
}

func (t *pTOTP) ToURL() string {
	return fmt.Sprintf(
		"otpauth://%s/%s:%s?secret=%s&issuer=%s",
		"totp",
		t.opts.Issuer,
		t.opts.Account,
		t.opts.Secret,
		t.opts.Issuer,
	)
}

func (t *pTOTP) QRCode() QRCode {
	return NewQR(t.ToURL())
}

/*********************************************/

func NewTOTP(opts Options) TOTP {
	return &pTOTP{
		opts: &opts,
	}
}
