package dz

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dz struct {
	locale string
}

// New returns a new instance of translator for the 'dz' locale
func New() locales.Translator {
	return &dz{
		locale: "dz",
	}
}

// Locale returns the current translators string locale
func (l *dz) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *dz) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
