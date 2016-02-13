package es_419

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "es_419",
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
