package kln_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kln_KE struct {
	locale string
}

// New returns a new instance of translator for the 'kln_KE' locale
func New() locales.Translator {
	return &kln_KE{
		locale: "kln_KE",
	}
}

// Locale returns the current translators string locale
func (l *kln_KE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *kln_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
