package dje

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dje struct {
	locale string
}

// New returns a new instance of translator for the 'dje' locale
func New() locales.Translator {
	return &dje{
		locale: "dje",
	}
}

// Locale returns the current translators string locale
func (l *dje) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *dje) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
