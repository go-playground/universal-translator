package mer_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mer_KE struct {
	locale string
}

// New returns a new instance of translator for the 'mer_KE' locale
func New() locales.Translator {
	return &mer_KE{
		locale: "mer_KE",
	}
}

// Locale returns the current translators string locale
func (l *mer_KE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *mer_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
