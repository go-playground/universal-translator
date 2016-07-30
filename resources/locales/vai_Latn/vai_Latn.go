package vai_Latn

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type vai_Latn struct {
	locale string
}

// New returns a new instance of translator for the 'vai_Latn' locale
func New() locales.Translator {
	return &vai_Latn{
		locale: "vai_Latn",
	}
}

// Locale returns the current translators string locale
func (l *vai_Latn) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *vai_Latn) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
