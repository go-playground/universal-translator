package kea_CV

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kea_CV struct {
	locale string
}

// New returns a new instance of translator for the 'kea_CV' locale
func New() locales.Translator {
	return &kea_CV{
		locale: "kea_CV",
	}
}

// Locale returns the current translators string locale
func (l *kea_CV) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *kea_CV) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
