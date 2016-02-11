package kok

import "github.com/go-playground/universal-translator"

func newSymbols() ut.Symbols {
	return ut.Symbols{}
}

func newFormats() ut.NumberFormats {
	return ut.NumberFormats{Decimal: "#,##,##0.###", Currency: "¤\u00a0#,##,##0.00", Percent: "#,##,##0%"}
}
