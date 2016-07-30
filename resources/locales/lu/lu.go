package lu

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lu struct {
	locale string
}

// New returns a new instance of translator for the 'lu' locale
func New() locales.Translator {
	return &lu{
		locale: "lu",
	}
}

// Locale returns the current translators string locale
func (l *lu) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *lu) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
