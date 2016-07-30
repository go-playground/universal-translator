package dje_NE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dje_NE struct {
	locale string
}

// New returns a new instance of translator for the 'dje_NE' locale
func New() locales.Translator {
	return &dje_NE{
		locale: "dje_NE",
	}
}

// Locale returns the current translators string locale
func (l *dje_NE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *dje_NE) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
