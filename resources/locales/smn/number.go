package smn

import "github.com/go-playground/universal-translator"

var (
	symbols = ut.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""}
	formats = ut.NumberFormats{Decimal: "0 tuhháát", Currency: "¤\u00a00K", CurrencyAccounting: "¤\u00a00K", Percent: ""}
)
