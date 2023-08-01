package monero

import (
	"fmt"
)

type XMRPayment struct {
	PaymentID   XMRPaymentID `json:"payment_id"`
	Amount      XMRAtom      `json:"amount"`
	BlockHeight XMRHeight    `json:"block_height"`
	BlockAge    XMRHeight    `json:"-"`
}

func (m *Monero) MakeIntegratedAddress() (a XMRAddress, i XMRPaymentID, h XMRHeight, err error) {
	
	var result struct {
		IntegratedAddress XMRAddress   `json:"integrated_address"`
		PaymentID         XMRPaymentID `json:"payment_id"`
		Height            XMRHeight    `json:"height"`
	}
	
	err = m.RPCQuery("/json_rpc", "POST", "get_height", nil, &result)
	if err != nil { return }
	
	err = m.RPCQuery("/json_rpc", "POST", "make_integrated_address", nil, &result)
	if err != nil { return }
	
	a = result.IntegratedAddress
	i = result.PaymentID
	h = result.Height
	
	return
}

func (m *Monero) GetBulkPayments(height XMRHeight, ids ...XMRPaymentID) (tl []XMRPayment, err error) {
	var params struct {
		PaymentsIDs    []XMRPaymentID `json:"payment_ids"`
		MinBlockHeight   string       `json:"min_block_height"`
	}
	var result struct {
		Payments []XMRPayment `json:"payments"`
	}
	
	params.PaymentsIDs = ids
	params.MinBlockHeight = (height-10).String()
	
	err = m.RPCQuery("/json_rpc", "POST", "get_bulk_payments", params, &result)
	if err != nil { return }
	
	tl = result.Payments
	err = m.UpdateBlockAge(tl)
	if err != nil { return }
	
	return
}

func (m *Monero) GetHeight() (h XMRHeight, err error) {
	var res struct {
		Height XMRHeight `json:"height"`
	}
	err = m.RPCQuery("/json_rpc", "POST", "get_height", nil, &res)
	if err != nil {
		return
	}
	h = res.Height
	return
}

func (m *Monero) UpdateBlockAge(tl []XMRPayment) (err error) {
	var h XMRHeight
	var i int
	h, err = m.GetHeight()
	if err != nil { return }
	for i, _ = range tl {
		tl[i].BlockAge = h - tl[i].BlockHeight
	}
	return
}

func (m *Monero) SearchPayment(l []XMRPayment, id XMRPaymentID) (*XMRPayment, bool) {
	for i, _ := range l {
		if l[i].PaymentID == id {
			return &l[i], true
		}
	}
	return nil, false
}

func (p XMRPayment) String() string {
	return fmt.Sprintf(
		"%-20v %-10v %-12v %-12v",
		p.PaymentID,
		p.Amount.ToXMR(),
		p.BlockHeight,
		p.BlockAge,
	)
}

var TitleXMRPayment string = fmt.Sprintf(
	"%-20v %-10v %-12v %-12v",
	"PaymentID",
	"Amount",
	"BlockHeight",
	"BlockAge",
)

