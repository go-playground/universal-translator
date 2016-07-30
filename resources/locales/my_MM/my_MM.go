package my_MM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type my_MM struct {
	locale string
}

// New returns a new instance of translator for the 'my_MM' locale
func New() locales.Translator {
	return &my_MM{
		locale: "my_MM",
	}
}

// Locale returns the current translators string locale
func (l *my_MM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *my_MM) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
