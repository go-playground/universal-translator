package fr_CA

import "github.com/go-playground/universal-translator"

var (
	symbols = ut.Symbols{Decimal: "٫", Group: "٬", Negative: "\u200f−", Percent: "٪", PerMille: "؉"}
	formats = ut.NumberFormats{Decimal: "0 milliard", Currency: "0\u00a0mn\u00a0¤", CurrencyAccounting: "0\u00a0mn\u00a0¤", Percent: "#,##0\u00a0%"}
)
