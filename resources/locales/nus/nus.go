package nus

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nus struct {
	locale string
}

// New returns a new instance of translator for the 'nus' locale
func New() locales.Translator {
	return &nus{
		locale: "nus",
	}
}

// Locale returns the current translators string locale
func (l *nus) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *nus) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
