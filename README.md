# GO-MONERO

A monero RPC client library and program.

## TODO

- Add digest authentication.

## Go programs

    Usage: monero-cli [-p PORT] ...
    
    A simple monero RPC client.
    
      -B                  : View balance.
      -R AMOUNT [-d DESC] : Print payment URI and payment ID, height to charge.
      -P URI              : Pay payment URI.
      -C HEIGHT PAYID...  : Check payments.
      -L aiopfd           : List transfers (all,in,out,pending,failed,[d]pool)
      -A                  : Get height.
    
    Copyright (c) 2023 Harkaitz Agirre, harkaitz.aguirre@gmail.com

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

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
