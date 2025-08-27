package main

import (
	"fmt"

	thaiqrpayment "github.com/iamohmcub/thai-qr-payment/pkg"
)

func main() {
	qr, err := thaiqrpayment.GenerateQR("0123456789", "MerchantName", true, 128)
	if err != nil {
		fmt.Println("Error generating QR code:", err)
		return
	}
	fmt.Println(qr)
}
