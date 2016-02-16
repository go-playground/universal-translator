package ta

import "github.com/go-playground/universal-translator"

var (
	symbols = ut.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"}
	formats = ut.NumberFormats{Decimal: "#,##,##0.###", Currency: "¤\u00a0#,##,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##,##0%"}
)
