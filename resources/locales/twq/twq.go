package twq

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type twq struct {
	locale string
}

// New returns a new instance of translator for the 'twq' locale
func New() locales.Translator {
	return &twq{
		locale: "twq",
	}
}

// Locale returns the current translators string locale
func (l *twq) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *twq) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
