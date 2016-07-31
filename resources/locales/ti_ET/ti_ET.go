package ti_ET

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ti_ET struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ti_ET' locale
func New() locales.Translator {
	return &ti_ET{
		locale:  "ti_ET",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ti_ET) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ti_ET'
func (t *ti_ET) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ti_ET'
func (t *ti_ET) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
