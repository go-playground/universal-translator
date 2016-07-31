package yo_BJ

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "yo_BJ",
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
