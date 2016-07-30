package dav

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dav struct {
	locale string
}

// New returns a new instance of translator for the 'dav' locale
func New() locales.Translator {
	return &dav{
		locale: "dav",
	}
}

// Locale returns the current translators string locale
func (l *dav) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *dav) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
