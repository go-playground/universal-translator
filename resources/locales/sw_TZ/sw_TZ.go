package sw_TZ

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sw_TZ struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sw_TZ' locale
func New() locales.Translator {
	return &sw_TZ{
		locale:  "sw_TZ",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sw_TZ) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sw_TZ'
func (t *sw_TZ) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sw_TZ'
func (t *sw_TZ) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
