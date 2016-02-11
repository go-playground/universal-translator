package es_DO

import "github.com/go-playground/universal-translator"

func newSymbols() ut.Symbols {
	return ut.Symbols{}
}

func newFormats() ut.NumberFormats {
	return ut.NumberFormats{Decimal: "", Currency: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"}
}
