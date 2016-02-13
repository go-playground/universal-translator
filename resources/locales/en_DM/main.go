package en_DM

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "en_DM",
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
