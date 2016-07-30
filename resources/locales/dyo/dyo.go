package dyo

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dyo struct {
	locale string
}

// New returns a new instance of translator for the 'dyo' locale
func New() locales.Translator {
	return &dyo{
		locale: "dyo",
	}
}

// Locale returns the current translators string locale
func (l *dyo) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *dyo) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
