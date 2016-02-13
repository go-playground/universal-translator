package en_NG

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "en_NG",
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
