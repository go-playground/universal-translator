package fa

import "github.com/go-playground/universal-translator"

var (
	symbols = ut.Symbols{Decimal: "٫", Group: "٬", Negative: "", Percent: "٪", PerMille: "؉"}
	formats = ut.NumberFormats{Decimal: "#,##0.###", Currency: "\u200e¤#,##0.00", Percent: "#,##0%"}
)
