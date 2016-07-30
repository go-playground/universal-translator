package agq_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type agq_CM struct {
	locale string
}

// New returns a new instance of translator for the 'agq_CM' locale
func New() locales.Translator {
	return &agq_CM{
		locale: "agq_CM",
	}
}

// Locale returns the current translators string locale
func (l *agq_CM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *agq_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
