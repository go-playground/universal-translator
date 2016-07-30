package kde

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kde struct {
	locale string
}

// New returns a new instance of translator for the 'kde' locale
func New() locales.Translator {
	return &kde{
		locale: "kde",
	}
}

// Locale returns the current translators string locale
func (l *kde) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *kde) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
