package ebu

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ebu struct {
	locale string
}

// New returns a new instance of translator for the 'ebu' locale
func New() locales.Translator {
	return &ebu{
		locale: "ebu",
	}
}

// Locale returns the current translators string locale
func (l *ebu) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ebu) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
