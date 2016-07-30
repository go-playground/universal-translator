package twq_NE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type twq_NE struct {
	locale string
}

// New returns a new instance of translator for the 'twq_NE' locale
func New() locales.Translator {
	return &twq_NE{
		locale: "twq_NE",
	}
}

// Locale returns the current translators string locale
func (l *twq_NE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *twq_NE) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
