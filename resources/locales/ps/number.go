package ps

import "github.com/go-playground/universal-translator"

var (
	symbols = ut.Symbols{Decimal: "٫", Group: "٬", Negative: "", Percent: "٪", PerMille: ""}
	formats = ut.NumberFormats{Decimal: "", Currency: "#,##0.00\u00a0¤", CurrencyAccounting: "#,##0.00\u00a0¤", Percent: ""}
)
