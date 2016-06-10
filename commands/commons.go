package commands

import (
	"fmt"
	"strings"

	"github.com/caarlos0/coinbase/api"
)

func BalancePrintf(native bool, acc api.Account) {
	MoneyPrintf(native, acc.Balance, acc.NativeBalance)
}

func MoneyPrintf(native bool, money, nativeMoney api.Money) {
	var amount string
	var currency string
	if native {
		amount = nativeMoney.Amount
		currency = nativeMoney.Currency
	} else {
		amount = money.Amount
		currency = money.Currency
	}
	if !strings.HasPrefix(amount, "-") {
		amount = "+" + amount
	}
	fmt.Printf("%s %s", amount, currency)
}
