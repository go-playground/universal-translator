package qu

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type qu struct {
	locale string
}

// New returns a new instance of translator for the 'qu' locale
func New() locales.Translator {
	return &qu{
		locale: "qu",
	}
}

// Locale returns the current translators string locale
func (l *qu) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *qu) CardinalPluralRule(num string) (locales.PluralRule, error) {

}
