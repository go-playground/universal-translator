package es_DO

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "es_DO",
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
