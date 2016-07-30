package yo_BJ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type yo_BJ struct {
	locale string
}

// New returns a new instance of translator for the 'yo_BJ' locale
func New() locales.Translator {
	return &yo_BJ{
		locale: "yo_BJ",
	}
}

// Locale returns the current translators string locale
func (l *yo_BJ) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *yo_BJ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
