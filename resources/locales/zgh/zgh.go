package zgh

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zgh struct {
	locale string
}

// New returns a new instance of translator for the 'zgh' locale
func New() locales.Translator {
	return &zgh{
		locale: "zgh",
	}
}

// Locale returns the current translators string locale
func (l *zgh) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *zgh) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
