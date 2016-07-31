package kab

import "github.com/go-playground/universal-translator"

var (
	symbols = ut.Symbols{Decimal: ",", Group: "\u00a0", Negative: "", Percent: "", PerMille: ""}
	formats = ut.NumberFormats{Decimal: "", Currency: "#,##0.00¤", CurrencyAccounting: "#,##0.00¤", Percent: ""}
)
