package en_150

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "en_150",
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
