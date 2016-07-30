package vai

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type vai struct {
	locale string
}

// New returns a new instance of translator for the 'vai' locale
func New() locales.Translator {
	return &vai{
		locale: "vai",
	}
}

// Locale returns the current translators string locale
func (l *vai) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *vai) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
