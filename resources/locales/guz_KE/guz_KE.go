package guz_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type guz_KE struct {
	locale string
}

// New returns a new instance of translator for the 'guz_KE' locale
func New() locales.Translator {
	return &guz_KE{
		locale: "guz_KE",
	}
}

// Locale returns the current translators string locale
func (l *guz_KE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *guz_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
