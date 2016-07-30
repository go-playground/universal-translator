package ewo

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ewo struct {
	locale string
}

// New returns a new instance of translator for the 'ewo' locale
func New() locales.Translator {
	return &ewo{
		locale: "ewo",
	}
}

// Locale returns the current translators string locale
func (l *ewo) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ewo) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
