package commands

import (
	"fmt"
	"strings"

	"github.com/caarlos0/coinbase/api"
	"github.com/urfave/cli"
)

func printfBalance(native bool, acc api.Account) {
	printfMoney(native, acc.Balance, acc.NativeBalance)
}

func printfMoney(native bool, money, nativeMoney api.Money) {
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

var nativeFlag = cli.BoolFlag{
	Name:  "native, n",
	Usage: "Use native currency",
}
