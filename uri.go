package monero

import (
	"errors"
)

type XMRURI struct {
	Amount        XMRAtom      `json:"amount"`
	Address       XMRAddress   `json:"address"`
	RecipientName string       `json:"recipient_name"`
	TxDescription string       `json:"tx_description"`
}

func (m *Monero) MakeURI(amount XMRAtom, desc string) (uri string, id XMRPaymentID, h XMRHeight, err error) {
	var params XMRURI
	var result struct { Uri string `json:"uri"` }
	
	params.Amount = amount
	params.Address, id, h, err = m.MakeIntegratedAddress()
	if err != nil { return }
	params.RecipientName = m.RecipientName
	params.TxDescription = desc
	
	err = m.RPCQuery("/json_rpc", "POST", "make_uri", params, &result)
	if err != nil { return }
	
	if result.Uri == "" {
		err = errors.New("No URI in response.")
		return
	}
	
	uri = result.Uri
	return
}

func (m *Monero) ParseURI(uri string) (uriS XMRURI, err error) {
	var params struct { Uri string `json:"uri"` }
	var result struct { Uri XMRURI `json:"uri"` }
	
	params.Uri = uri
	
	err = m.RPCQuery("/json_rpc", "POST", "parse_uri", params, &result)
	if err != nil { return }
	
	if result.Uri.Address == "" {
		err = errors.New("The uri does not contain an address.")
		return
	}
	if result.Uri.Amount == 0 {
		err = errors.New("The uri does not contain an amount.")
		return
	}
	
	uriS = result.Uri
	return
}

func (m *Monero) PayURI(uri string) (err error) {
	var uriS XMRURI
	var dest []XMRDestination = make([]XMRDestination, 1)
	
	uriS, err = m.ParseURI(uri)
	if err != nil { return }
	
	dest[0].Address = uriS.Address
	dest[0].Amount  = uriS.Amount
	
	return m.Transfer(EmptyXMRPaymentID, dest)
}
