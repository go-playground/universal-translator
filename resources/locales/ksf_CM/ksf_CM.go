package ksf_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ksf_CM struct {
	locale string
}

// New returns a new instance of translator for the 'ksf_CM' locale
func New() locales.Translator {
	return &ksf_CM{
		locale: "ksf_CM",
	}
}

// Locale returns the current translators string locale
func (l *ksf_CM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ksf_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
