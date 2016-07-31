package zh_Hant_HK

import "github.com/go-playground/universal-translator"

var (
	symbols = ut.Symbols{Decimal: ".", Group: ",", Negative: "-", Percent: "%", PerMille: "‰"}
	formats = ut.NumberFormats{Decimal: "0K", Currency: "¤0K", CurrencyAccounting: "¤0K", Percent: "#,##0%"}
)
