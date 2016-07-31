package es_PH

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "es_PH",
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
