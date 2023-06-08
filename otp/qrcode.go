package otp

import (
	"encoding/base64"
	"fmt"

	"github.com/skip2/go-qrcode"
)

type QRCode interface {
	Encode(size int, level ...qrcode.RecoveryLevel) ([]byte, error)
	ToDataURL(size int, level ...qrcode.RecoveryLevel) (value string, err error)
}

type pQRCode struct {
	data string
}

func (q *pQRCode) Encode(size int, level ...qrcode.RecoveryLevel) ([]byte, error) {
	if len(level) == 0 {
		level = append(level, qrcode.Medium)
	}

	return qrcode.Encode(q.data, level[0], size)
}

func (q *pQRCode) ToDataURL(size int, level ...qrcode.RecoveryLevel) (value string, err error) {
	data, err := q.Encode(size, level...)

	if err != nil {
		return
	}

	dataStr := base64.StdEncoding.EncodeToString(data)
	value = fmt.Sprintf("data:image/png;base64,%s", dataStr)

	return
}

/**********************************************/

func NewQR(data string) QRCode {
	return &pQRCode{
		data: data,
	}
}
