package en_VU

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "en_VU",
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
