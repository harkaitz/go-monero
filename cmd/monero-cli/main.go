package main

import (
	"os"
	"fmt"
	"github.com/pborman/getopt/v2"
	"github.com/harkaitz/go-monero"
)

const help string =
`Usage: monero-cli [-p PORT] ...

A simple monero RPC client.

  -B                  : View balance.
  -R AMOUNT [-d DESC] : Print payment URI and payment ID, height to charge.
  -P URI [AMOUNT]     : Pay payment URI.
  -C HEIGHT PAYID...  : Check payments.
  -L aiopfd           : List transfers (all,in,out,pending,failed,[d]pool)
  -A                  : Get height.

Copyright (c) 2023 Harkaitz Agirre, harkaitz.aguirre@gmail.com`

func main() {
	
	var err         error
	var m           monero.Monero
	var balance     monero.Balance
	var amount      monero.XMRAtom
	var height      monero.XMRHeight
	var ids       []monero.XMRPaymentID
	var paymentID   monero.XMRPaymentID
	var pays      []monero.XMRPayment
	var pay         monero.XMRPayment
	var transfers []monero.XMRTransfer
	var transfer    monero.XMRTransfer
	var id          string
	var i           int
	var uri         string
	var args      []string
	
	defer func() {
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err.Error())
			os.Exit(1)
		}
	}()
	
	pFlag := getopt.Int('p', 0)
	dFlag := getopt.String('d', "")
	BFlag := getopt.Bool('B')
	RFlag := getopt.String('R', "")
	PFlag := getopt.String('P', "")
	CFlag := getopt.String('C', "")
	LFlag := getopt.String('L', "")
	AFlag := getopt.Bool('A')
	
	getopt.SetUsage(func() { fmt.Println(help) })
	getopt.Parse()
	
	m = monero.CreateMonero(*pFlag)
	
	switch {
	case *BFlag:
		balance, err = m.GetBalance()
		if err != nil { return }
		
		fmt.Print(balance)
	case *RFlag != "":
		amount, err = monero.StrXMRAtom(*RFlag)
		if err != nil { return }
		
		uri, paymentID, height, err = m.MakeURI(amount, *dFlag)
		if err != nil { return }
		
		fmt.Printf("%v\n%v %v\n", uri, height, paymentID)
	case *PFlag != "":
		args = getopt.Args()
		if len(args) == 1 {
			amount, err = monero.StrXMRAtom(args[0])
			if err != nil { return }
			err = m.PayURIAmount(*PFlag, amount)
			if err != nil { return }
		} else {
			err = m.PayURI(*PFlag)
			if err != nil { return }
		}
	case *CFlag != "":
		
		height, err = monero.StrXMRHeight(*CFlag)
		if err != nil { return }
		
		ids = make([]monero.XMRPaymentID, len(getopt.Args()))
		for i, id = range getopt.Args() {
			ids[i], err = monero.StrXMRPaymentID(id)
			if err != nil { return }
		}
		
		pays, err = m.GetBulkPayments(height, ids...)
		if err != nil { return }
		
		fmt.Println(monero.TitleXMRPayment)
		for _, pay = range pays {
			fmt.Println(pay)
		}
	case *LFlag != "":
		if *LFlag == "a" {
			*LFlag = "iopfd"
		}
		transfers, err = m.GetTransfers(*LFlag)
		if err != nil {
			return
		}
		fmt.Println(monero.TitleXMRTransfer)
		for _, transfer = range transfers {
			fmt.Println(transfer)
		}
	case *AFlag:
		height, err = m.GetHeight()
		if err != nil {
			return
		}
		fmt.Printf("%v\n", height)
	default:
		getopt.Usage()
	}
	
	return
}
