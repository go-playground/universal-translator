package yo_NG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type yo_NG struct {
	locale string
}

// New returns a new instance of translator for the 'yo_NG' locale
func New() locales.Translator {
	return &yo_NG{
		locale: "yo_NG",
	}
}

// Locale returns the current translators string locale
func (l *yo_NG) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *yo_NG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
