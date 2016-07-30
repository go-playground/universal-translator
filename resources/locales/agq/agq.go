package agq

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type agq struct {
	locale string
}

// New returns a new instance of translator for the 'agq' locale
func New() locales.Translator {
	return &agq{
		locale: "agq",
	}
}

// Locale returns the current translators string locale
func (l *agq) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *agq) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
