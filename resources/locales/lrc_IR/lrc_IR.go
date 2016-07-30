package lrc_IR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lrc_IR struct {
	locale string
}

// New returns a new instance of translator for the 'lrc_IR' locale
func New() locales.Translator {
	return &lrc_IR{
		locale: "lrc_IR",
	}
}

// Locale returns the current translators string locale
func (l *lrc_IR) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *lrc_IR) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
