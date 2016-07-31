package pa_Arab

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "pa_Arab",
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
