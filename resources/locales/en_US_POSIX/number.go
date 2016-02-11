package en_US_POSIX

import "github.com/go-playground/universal-translator"

func newSymbols() ut.Symbols {
	return ut.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: "0/00"}
}

func newFormats() ut.NumberFormats {
	return ut.NumberFormats{Decimal: "#0.######", Currency: "¤\u00a0#0.00", Percent: "#0%"}
}
