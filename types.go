package monero

import (
	"strconv"
	"errors"
)

type XMRAtom      int64
type XMR          float64
type XMRAddress   string
type XMRPaymentID string
type XMRHeight    uint64

type XMRDestination struct {
	Address XMRAddress `json:"address"`
	Amount  XMRAtom    `json:"amount"`
}

var EmptyXMRPaymentID XMRPaymentID = ""

func (f XMR) ToAtom() (t XMRAtom) {
	return XMRAtom(float64(f) * 1e12)
}

func (f XMRAtom) ToXMR() (t XMR) {
	return XMR(float64(f) / 1e12)
}

func (h XMRHeight) String() (s string) {
	return strconv.Itoa(int(h))
}

func StrXMR(f string) (t XMR, err error) {
	n, err := strconv.ParseFloat(f, 64)
	if err != nil {
		return
	}
	return XMR(n), nil
}

func StrXMRAtom(str string) (atom XMRAtom, err error) {
	xmr, err := StrXMR(str)
	if err != nil {
		return
	}
	return xmr.ToAtom(), nil
}

func StrXMRAddress(str string) (t XMRAddress, err error) {
	if len(str) != 106 && len(str) != 95 {
		err = errors.New("Invalid XMR address")
		return
	}
	return XMRAddress(str), nil
}

func StrXMRPaymentID(str string) (t XMRPaymentID, err error) {
	if str == "-" {
		t = ""
		return
	}
	if len(str) != 16 {
		err = errors.New("Invalid XMR payment ID")
		return
	}
	t = XMRPaymentID(str)
	return
}

func StrXMRHeight(str string) (h XMRHeight, err error) {
	var hi uint64
	hi, err = strconv.ParseUint(str, 10, 64)
	if err != nil { return }
	h = XMRHeight(hi)
	return
}
