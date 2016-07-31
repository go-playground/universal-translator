package fr_KM

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "fr_KM",
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
