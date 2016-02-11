package ses

import "github.com/go-playground/universal-translator"

func newSymbols() ut.Symbols {
	return ut.Symbols{Decimal: "", Group: "\u00a0", Negative: "", Percent: "", PerMille: ""}
}

func newFormats() ut.NumberFormats {
	return ut.NumberFormats{Decimal: "", Currency: "#,##0.00¤", Percent: ""}
}
