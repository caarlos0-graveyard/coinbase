package commands

import (
	"fmt"
	"strings"

	"github.com/caarlos0/coinbase/api"
)

// BalancePrintf prints balance to stdout
func BalancePrintf(native bool, acc api.Account) {
	MoneyPrintf(native, acc.Balance, acc.NativeBalance)
}

// MoneyPrintf prints money to stdout
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
	fmt.Printf("%s %s\t", amount, currency)
}
