package ki_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ki_KE struct {
	locale string
}

// New returns a new instance of translator for the 'ki_KE' locale
func New() locales.Translator {
	return &ki_KE{
		locale: "ki_KE",
	}
}

// Locale returns the current translators string locale
func (l *ki_KE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ki_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
