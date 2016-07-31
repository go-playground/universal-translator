package fr_CH

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "fr_CH",
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
