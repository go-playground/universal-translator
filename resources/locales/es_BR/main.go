package es_BR

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "es_BR",
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
