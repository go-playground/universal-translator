package om_ET

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type om_ET struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'om_ET' locale
func New() locales.Translator {
	return &om_ET{
		locale:  "om_ET",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *om_ET) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'om_ET'
func (t *om_ET) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'om_ET'
func (t *om_ET) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
