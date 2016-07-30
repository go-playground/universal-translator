package nmg

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type nmg struct {
	locale string
}

// New returns a new instance of translator for the 'nmg' locale
func New() locales.Translator {
	return &nmg{
		locale: "nmg",
	}
}

// Locale returns the current translators string locale
func (l *nmg) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *nmg) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
