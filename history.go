package monero

import (
	"fmt"
)

type XMRTransfer struct {
	Address        XMRAddress     `json:"address"`
	Amount         XMRAtom        `json:"amount"`
	Fee            XMRAtom        `json:"fee"`
	Locked         bool           `json:"locked"`
	PaymentID      XMRPaymentID   `json:"payment_id"`
	TxID           string         `json:"txid"`
	Type           string         `json:"type"`
	UnlockTime     float64        `json:"unlock_time"`
	Destinations []XMRDestination `json:"destinations"`
}

func (m *Monero) GetTransfers(sels string) (tl []XMRTransfer, err error) {
	type params struct {
		In      bool `json:"in"`
		Out     bool `json:"out"`
		Pending bool `json:"pending"`
		Failed  bool `json:"failed"`
		Pool    bool `json:"pool"`
	}
	type result struct {
		In      []XMRTransfer `json:"in"`
		Out     []XMRTransfer `json:"out"`
		Pending []XMRTransfer `json:"pending"`
		Failed  []XMRTransfer `json:"failed"`
		Pool    []XMRTransfer `json:"pool"`
	}
	
	var req params
	var res result
	
	for _, sel := range sels {
		switch sel {
		case 'i': req.In      = true;
		case 'o': req.Out     = true;
		case 'p': req.Pending = true;
		case 'f': req.Failed  = true;
		case 'd': req.Pool    = true;
		}
	}
	
	err = m.RPCQuery("/json_rpc", "POST", "get_transfers", req, &res)
	
	if err != nil { return }
	
	tl = []XMRTransfer{}
	if res.In      != nil { tl = append(tl, res.In...)      }
	if res.Out     != nil { tl = append(tl, res.Out...)     }
	if res.Pending != nil { tl = append(tl, res.Pending...) }
	if res.Failed  != nil { tl = append(tl, res.Failed...)  }
	if res.Pool    != nil { tl = append(tl, res.Pool...)    }
	
	return
}

var TitleXMRTransfer string = fmt.Sprintf(
	"%-8v %-8v %-12v %v",
	"Type",
	"Locked",
	"Amount",
	"PaymentID",
)

func (t XMRTransfer) String() string {
	return fmt.Sprintf(
		"%-8v %-8v %-12v %v",
		t.Type,
		t.Locked,
		t.TotalAmount().ToXMR(),
		t.PaymentID,
	)
}

func (t *XMRTransfer) TotalAmount() (r XMRAtom) {
	if len(t.Destinations) > 0 {
		for _, d := range t.Destinations {
			r += d.Amount
		}
	} else {
		r = t.Amount
	}
	return
}
