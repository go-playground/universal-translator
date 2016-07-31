package hr_BA

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "hr_BA",
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
