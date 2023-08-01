# GO-MONERO

A monero RPC client.

## TODO

- Add digest authentication.

## Go struct Monero

    package monero // import "."
    
    type Monero struct {
        u.RPC
        RecipientName string
    }
    
    func CreateMonero(port int) (m Monero)
    func (m *Monero) GetBalance() (b Balance, err error)
    func (m *Monero) GetBulkPayments(height XMRHeight, ids ...XMRPaymentID) (tl []XMRPayment, err error)
    func (m *Monero) GetHeight() (h XMRHeight, err error)
    func (m *Monero) GetTransfers(sels string) (tl []XMRTransfer, err error)
    func (m *Monero) MakeIntegratedAddress() (a XMRAddress, i XMRPaymentID, h XMRHeight, err error)
    func (m *Monero) MakeURI(amount XMRAtom, desc string) (uri string, id XMRPaymentID, h XMRHeight, err error)
    func (m *Monero) ParseURI(uri string) (uriS XMRURI, err error)
    func (m *Monero) PayURI(uri string) (err error)
    func (m *Monero) SearchPayment(l []XMRPayment, id XMRPaymentID) (*XMRPayment, bool)
    func (m *Monero) Transfer(id XMRPaymentID, d []XMRDestination) (err error)
    func (m *Monero) UpdateBlockAge(tl []XMRPayment) (err error)

