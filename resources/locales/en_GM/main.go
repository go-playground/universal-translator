package en_GM

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "en_GM",
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
