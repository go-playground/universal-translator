package vai_Vaii

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type vai_Vaii struct {
	locale string
}

// New returns a new instance of translator for the 'vai_Vaii' locale
func New() locales.Translator {
	return &vai_Vaii{
		locale: "vai_Vaii",
	}
}

// Locale returns the current translators string locale
func (l *vai_Vaii) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *vai_Vaii) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
