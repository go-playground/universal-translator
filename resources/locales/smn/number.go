package smn

import "github.com/go-playground/universal-translator"

func newSymbols() ut.Symbols {
	return ut.Symbols{}
}

func newFormats() ut.NumberFormats {
	return ut.NumberFormats{Decimal: "0 tuhháát", Currency: "¤\u00a00K", Percent: ""}
}
