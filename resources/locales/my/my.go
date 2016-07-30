package my

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type my struct {
	locale string
}

// New returns a new instance of translator for the 'my' locale
func New() locales.Translator {
	return &my{
		locale: "my",
	}
}

// Locale returns the current translators string locale
func (l *my) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *my) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
