package bo_IN

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "bo_IN",
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
