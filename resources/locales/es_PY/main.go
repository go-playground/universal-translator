package es_PY

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "es_PY",
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
