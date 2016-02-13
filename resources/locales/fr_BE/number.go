package fr_BE

import "github.com/go-playground/universal-translator"

var (
	symbols = ut.Symbols{Decimal: "٫", Group: ".", Negative: "\u200f−", Percent: "٪", PerMille: "؉"}
	formats = ut.NumberFormats{Decimal: "#,##0.###", Currency: "#,##0.00\u00a0¤", Percent: "#,##0\u00a0%"}
)
