package nmg_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nmg_CM struct {
	locale string
}

// New returns a new instance of translator for the 'nmg_CM' locale
func New() locales.Translator {
	return &nmg_CM{
		locale: "nmg_CM",
	}
}

// Locale returns the current translators string locale
func (l *nmg_CM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *nmg_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
