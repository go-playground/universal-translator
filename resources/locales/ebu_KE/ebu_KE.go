package ebu_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ebu_KE struct {
	locale string
}

// New returns a new instance of translator for the 'ebu_KE' locale
func New() locales.Translator {
	return &ebu_KE{
		locale: "ebu_KE",
	}
}

// Locale returns the current translators string locale
func (l *ebu_KE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ebu_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
