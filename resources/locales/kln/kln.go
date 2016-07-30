package kln

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kln struct {
	locale string
}

// New returns a new instance of translator for the 'kln' locale
func New() locales.Translator {
	return &kln{
		locale: "kln",
	}
}

// Locale returns the current translators string locale
func (l *kln) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *kln) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
