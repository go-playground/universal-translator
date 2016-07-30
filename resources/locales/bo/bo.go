package bo

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type bo struct {
	locale string
}

// New returns a new instance of translator for the 'bo' locale
func New() locales.Translator {
	return &bo{
		locale: "bo",
	}
}

// Locale returns the current translators string locale
func (l *bo) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *bo) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
