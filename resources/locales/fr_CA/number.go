package fr_CA

import "github.com/go-playground/universal-translator"

func newSymbols() ut.Symbols {
	return ut.Symbols{}
}

func newFormats() ut.NumberFormats {
	return ut.NumberFormats{Decimal: "0 milliard", Currency: "0\u00a0mn\u00a0¤", Percent: ""}
}
