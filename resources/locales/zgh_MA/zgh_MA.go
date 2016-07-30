package zgh_MA

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zgh_MA struct {
	locale string
}

// New returns a new instance of translator for the 'zgh_MA' locale
func New() locales.Translator {
	return &zgh_MA{
		locale: "zgh_MA",
	}
}

// Locale returns the current translators string locale
func (l *zgh_MA) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *zgh_MA) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
