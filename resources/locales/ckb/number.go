package ckb

import "github.com/go-playground/universal-translator"

var (
	symbols = ut.Symbols{Decimal: "", Group: "", Negative: "\u200e-", Percent: "٪", PerMille: ""}
	formats = ut.NumberFormats{Decimal: "", Currency: "", CurrencyAccounting: "", Percent: ""}
)
