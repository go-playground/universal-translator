package af_NA

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "af_NA",
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
