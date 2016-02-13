package fa_AF

import "github.com/go-playground/universal-translator"

var locale = &ut.Locale{
	Locale: "fa_AF",
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
