package bas_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bas_CM struct {
	locale string
}

// New returns a new instance of translator for the 'bas_CM' locale
func New() locales.Translator {
	return &bas_CM{
		locale: "bas_CM",
	}
}

// Locale returns the current translators string locale
func (l *bas_CM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *bas_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
