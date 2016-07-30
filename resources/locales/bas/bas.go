package bas

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bas struct {
	locale string
}

// New returns a new instance of translator for the 'bas' locale
func New() locales.Translator {
	return &bas{
		locale: "bas",
	}
}

// Locale returns the current translators string locale
func (l *bas) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *bas) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
