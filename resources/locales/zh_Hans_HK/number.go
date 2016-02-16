package zh_Hans_HK

import "github.com/go-playground/universal-translator"

var (
	symbols = ut.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"}
	formats = ut.NumberFormats{Decimal: "0万亿", Currency: "¤#,##0.00", CurrencyAccounting: "¤#,##0.00;(¤#,##0.00)", Percent: "#,##0%"}
)
