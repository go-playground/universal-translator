package ca_ES_VALENCIA

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "ca_ES_VALENCIA",
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
