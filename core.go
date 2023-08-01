package monero

import (
	"github.com/harkaitz/go-u27"
	"strconv"
)

type Monero struct {
	u.RPC
	RecipientName string
}

func CreateMonero(port int) (m Monero) {
	if port == 0 { port = 18081 }
	m.URL = "http://127.0.0.1:" + strconv.Itoa(port)
	return
}
