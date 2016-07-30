package kam_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kam_KE struct {
	locale string
}

// New returns a new instance of translator for the 'kam_KE' locale
func New() locales.Translator {
	return &kam_KE{
		locale: "kam_KE",
	}
}

// Locale returns the current translators string locale
func (l *kam_KE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *kam_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
