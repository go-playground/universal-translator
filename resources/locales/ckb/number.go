package ckb

import "github.com/go-playground/universal-translator"

func newSymbols() ut.Symbols {
	return ut.Symbols{Decimal: "", Group: "", Negative: "\u200e-", Percent: "٪", PerMille: ""}
}

func newFormats() ut.NumberFormats {
	return ut.NumberFormats{}
}
