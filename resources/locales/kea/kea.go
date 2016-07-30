package kea

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kea struct {
	locale string
}

// New returns a new instance of translator for the 'kea' locale
func New() locales.Translator {
	return &kea{
		locale: "kea",
	}
}

// Locale returns the current translators string locale
func (l *kea) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *kea) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
