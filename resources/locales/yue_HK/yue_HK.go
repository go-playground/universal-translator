package yue_HK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type yue_HK struct {
	locale string
}

// New returns a new instance of translator for the 'yue_HK' locale
func New() locales.Translator {
	return &yue_HK{
		locale: "yue_HK",
	}
}

// Locale returns the current translators string locale
func (l *yue_HK) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *yue_HK) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
