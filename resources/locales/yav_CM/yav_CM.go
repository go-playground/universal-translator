package yav_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type yav_CM struct {
	locale string
}

// New returns a new instance of translator for the 'yav_CM' locale
func New() locales.Translator {
	return &yav_CM{
		locale: "yav_CM",
	}
}

// Locale returns the current translators string locale
func (l *yav_CM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *yav_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
