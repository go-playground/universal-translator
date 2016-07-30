package yav

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type yav struct {
	locale string
}

// New returns a new instance of translator for the 'yav' locale
func New() locales.Translator {
	return &yav{
		locale: "yav",
	}
}

// Locale returns the current translators string locale
func (l *yav) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *yav) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
