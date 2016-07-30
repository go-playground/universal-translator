package nus_SS

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nus_SS struct {
	locale string
}

// New returns a new instance of translator for the 'nus_SS' locale
func New() locales.Translator {
	return &nus_SS{
		locale: "nus_SS",
	}
}

// Locale returns the current translators string locale
func (l *nus_SS) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *nus_SS) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
