package ha_GH

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ha_GH struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ha_GH' locale
func New() locales.Translator {
	return &ha_GH{
		locale:  "ha_GH",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ha_GH) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ha_GH'
func (t *ha_GH) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ha_GH'
func (t *ha_GH) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
