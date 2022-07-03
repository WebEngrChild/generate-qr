package qrgen

import (
	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

// GenQRCode generates QR code for the given URL
func GenQRCode(url string, width, height int) (barcode.Barcode, error) {
	qrCode, err := qr.Encode(url, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	return barcode.Scale(qrCode, width, height)
}