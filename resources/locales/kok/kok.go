package kok

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kok struct {
	locale string
}

// New returns a new instance of translator for the 'kok' locale
func New() locales.Translator {
	return &kok{
		locale: "kok",
	}
}

// Locale returns the current translators string locale
func (l *kok) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *kok) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
