package sah

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sah struct {
	locale string
}

// New returns a new instance of translator for the 'sah' locale
func New() locales.Translator {
	return &sah{
		locale: "sah",
	}
}

// Locale returns the current translators string locale
func (l *sah) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *sah) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
