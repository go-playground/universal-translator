package sah_RU

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sah_RU struct {
	locale string
}

// New returns a new instance of translator for the 'sah_RU' locale
func New() locales.Translator {
	return &sah_RU{
		locale: "sah_RU",
	}
}

// Locale returns the current translators string locale
func (l *sah_RU) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *sah_RU) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
