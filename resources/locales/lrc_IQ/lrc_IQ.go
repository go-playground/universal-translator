package lrc_IQ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lrc_IQ struct {
	locale string
}

// New returns a new instance of translator for the 'lrc_IQ' locale
func New() locales.Translator {
	return &lrc_IQ{
		locale: "lrc_IQ",
	}
}

// Locale returns the current translators string locale
func (l *lrc_IQ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *lrc_IQ) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
