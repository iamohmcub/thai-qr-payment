package main

import (
	"errors"
	"fmt"
	"strings"
)

func formatTagValue(tag, value string) string {
	return fmt.Sprintf("%s%02d%s", tag, len(value), value)
}

func calculateCRC16CCITT(s string) string {
	var poly uint16 = 0x1021
	var crc uint16 = 0xFFFF
	for i := 0; i < len(s); i++ {
		crc ^= uint16(s[i]) << 8
		for range 8 {
			if (crc & 0x8000) != 0 {
				crc = (crc << 1) ^ poly
			} else {
				crc <<= 1
			}
		}
	}
	return fmt.Sprintf("%04X", crc&0xFFFF)
}

func formatMobileNumber(m string) (string, error) {
	m = strings.TrimSpace(m)
	if !reMobile10.MatchString(m) {
		return "", errors.New("mobile must be 10 digits")
	}
	if m[0] != '0' {
		return "", errors.New("mobile must start with 0")
	}
	return "0066" + m[1:], nil
}

func isValidCitizenID(id string) bool {
	if !reCID13.MatchString(id) {
		return false
	}
	sum := 0
	for i := range 12 {
		d := int(id[i] - '0')
		w := 13 - i
		sum += d * w
	}
	check := (11 - (sum % 11)) % 10
	return check == int(id[12]-'0')
}

func identifyIDType(id string) IDType {
	switch {
	case reMobile10.MatchString(id):
		return IDMobile
	case reCID13.MatchString(id):
		return IDCitizenID
	default:
		return IDUnknown
	}
}
