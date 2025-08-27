package thaiqrpayment

import "regexp"

type IDType int

const (
	IDUnknown IDType = iota
	IDMobile
	IDCitizenID
)

var (
	reMobile10 = regexp.MustCompile(`^\d{10}$`)
	reCID13    = regexp.MustCompile(`^\d{13}$`)
)
