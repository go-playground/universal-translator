package cy

import "github.com/go-playground/universal-translator"

var (
	symbols = ut.Symbols{Decimal: "", Group: "", Negative: "", Percent: "", PerMille: ""}
	formats = ut.NumberFormats{Decimal: "#,##0.###", Currency: "¤#,##0.00", Percent: "#,##0%"}
)
