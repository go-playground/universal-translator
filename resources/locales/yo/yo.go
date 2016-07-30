package yo

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type yo struct {
	locale string
}

// New returns a new instance of translator for the 'yo' locale
func New() locales.Translator {
	return &yo{
		locale: "yo",
	}
}

// Locale returns the current translators string locale
func (l *yo) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *yo) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
