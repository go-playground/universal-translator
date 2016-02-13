package zh_Hant

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "zh_Hant",
	Number: ut.Number{
		Symbols:    symbols,
		Formats:    formats,
		Currencies: currencies,
	},
	Calendar:   calendar,
	PluralRule: pluralRule,
}

func init() {
	ut.RegisterLocale(locale)
}
