package wae

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type wae struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'wae' locale
func New() locales.Translator {
	return &wae{
		locale:  "wae",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *wae) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'wae'
func (t *wae) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'wae'
func (t *wae) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
