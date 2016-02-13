package pa_Arab

import "github.com/go-playground/universal-translator"

var (
	symbols = ut.Symbols{Decimal: "٫", Group: ",", Negative: "\u200e-", Percent: "٪", PerMille: "؉"}
	formats = ut.NumberFormats{Decimal: "#,##,##0.###", Currency: "¤\u00a00K", Percent: "#,##,##0%"}
)
