package guz

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type guz struct {
	locale string
}

// New returns a new instance of translator for the 'guz' locale
func New() locales.Translator {
	return &guz{
		locale: "guz",
	}
}

// Locale returns the current translators string locale
func (l *guz) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *guz) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
