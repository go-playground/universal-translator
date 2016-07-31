package kl_GL

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type kl_GL struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'kl_GL' locale
func New() locales.Translator {
	return &kl_GL{
		locale:  "kl_GL",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *kl_GL) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'kl_GL'
func (t *kl_GL) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'kl_GL'
func (t *kl_GL) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
