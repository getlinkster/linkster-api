package main

import (
	"fmt"

	qrcode "github.com/skip2/go-qrcode"
)

func createQrCode(url string) ([]byte, error) {
	qrCodeString, err := sendGetQrRequest(url)
	if err != nil {
		return nil, err
	}
	fmt.Println("QR code string: ", qrCodeString)

	png, err := qrcode.Encode(qrCodeString, qrcode.Medium, 256)
	if err != nil {
		return nil, err
	}

	return png, nil
}
