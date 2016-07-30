package ksf

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ksf struct {
	locale string
}

// New returns a new instance of translator for the 'ksf' locale
func New() locales.Translator {
	return &ksf{
		locale: "ksf",
	}
}

// Locale returns the current translators string locale
func (l *ksf) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ksf) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
