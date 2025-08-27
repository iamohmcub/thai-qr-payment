package thaiqrpayment

import (
	"errors"
	"strconv"
	"strings"

	qrcode "github.com/skip2/go-qrcode"
)

func GenerateQRString(id, amount string, dynamic bool) (string, error) {
	id = strings.TrimSpace(id)

	payload := formatTagValue("00", "01")

	if dynamic {
		payload += formatTagValue("01", "12")
	} else {
		payload += formatTagValue("01", "11")
	}

	aid := formatTagValue("00", "A000000677010111")

	var acct string
	switch identifyIDType(id) {
	case IDMobile:
		norm, err := formatMobileNumber(id)
		if err != nil {
			return "", err
		}
		acct = formatTagValue("01", norm)
	case IDCitizenID:
		if !isValidCitizenID(id) {
			return "", errors.New("invalid citizen id checksum")
		}
		acct = formatTagValue("02", id)
	default:
		return "", errors.New("unsupported id: must be 10-digit mobile or 13-digit citizen id")
	}

	payload += formatTagValue("29", aid+acct)
	payload += formatTagValue("53", "764")

	amount = strings.TrimSpace(amount)
	if amount != "" {
		if _, err := strconv.ParseFloat(amount, 64); err != nil {
			return "", errors.New("amount must be numeric (e.g., 50 or 50.00)")
		}
		payload += formatTagValue("54", amount)
	}

	payload += formatTagValue("58", "TH")

	tmp := payload + "6304"
	return tmp + calculateCRC16CCITT(tmp), nil
}

func GenerateQR(id, amount string, dynamic bool, sizePX int) ([]byte, error) {
	pl, err := GenerateQRString(id, amount, dynamic)
	if err != nil {
		return nil, err
	}
	return qrcode.Encode(pl, qrcode.Medium, sizePX)
}
