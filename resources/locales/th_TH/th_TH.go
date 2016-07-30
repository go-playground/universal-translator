package th_TH

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type th_TH struct {
	locale string
}

// New returns a new instance of translator for the 'th_TH' locale
func New() locales.Translator {
	return &th_TH{
		locale: "th_TH",
	}
}

// Locale returns the current translators string locale
func (l *th_TH) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *th_TH) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
