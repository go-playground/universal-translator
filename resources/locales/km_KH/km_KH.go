package km_KH

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type km_KH struct {
	locale string
}

// New returns a new instance of translator for the 'km_KH' locale
func New() locales.Translator {
	return &km_KH{
		locale: "km_KH",
	}
}

// Locale returns the current translators string locale
func (l *km_KH) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *km_KH) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
