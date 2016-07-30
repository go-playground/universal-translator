package dav_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type dav_KE struct {
	locale string
}

// New returns a new instance of translator for the 'dav_KE' locale
func New() locales.Translator {
	return &dav_KE{
		locale: "dav_KE",
	}
}

// Locale returns the current translators string locale
func (l *dav_KE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *dav_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
