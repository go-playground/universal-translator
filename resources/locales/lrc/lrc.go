package lrc

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lrc struct {
	locale string
}

// New returns a new instance of translator for the 'lrc' locale
func New() locales.Translator {
	return &lrc{
		locale: "lrc",
	}
}

// Locale returns the current translators string locale
func (l *lrc) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *lrc) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
