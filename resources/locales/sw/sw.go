package sw

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sw struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sw' locale
func New() locales.Translator {
	return &sw{
		locale:  "sw",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *sw) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sw'
func (t *sw) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sw'
func (t *sw) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
