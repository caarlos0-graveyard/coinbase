package commands

import (
	"fmt"

	"github.com/caarlos0/coinbase/api"
)

func BalancePrintf(native bool, acc api.Account) {
	var amount string
	var currency string
	if native {
		amount = acc.NativeBalance.Amount
		currency = acc.NativeBalance.Currency
	} else {
		amount = acc.Balance.Amount
		currency = acc.Balance.Currency
	}
	fmt.Printf("%s %s", amount, currency)
}
