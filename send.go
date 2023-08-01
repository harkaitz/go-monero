package monero

func (m *Monero) Transfer(id XMRPaymentID, d []XMRDestination) (err error) {
	var params struct {
		PaymentID    XMRPaymentID     `json:"payment_id,omitempty"`
		Destinations []XMRDestination `json:"destinations"`
	}
	params.Destinations = d
	params.PaymentID = id
	return m.RPCQuery("/json_rpc", "POST", "transfer_split", params, nil)
}
