package yue

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type yue struct {
	locale string
}

// New returns a new instance of translator for the 'yue' locale
func New() locales.Translator {
	return &yue{
		locale: "yue",
	}
}

// Locale returns the current translators string locale
func (l *yue) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *yue) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
