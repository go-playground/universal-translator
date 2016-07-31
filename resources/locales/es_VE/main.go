package es_VE

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "es_VE",
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
