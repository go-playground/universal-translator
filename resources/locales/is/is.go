package is

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type is struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'is' locale
func New() locales.Translator {
	return &is{
		locale:  "is",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *is) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'is'
func (t *is) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'is'
func (t *is) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	t, err := locales.T(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (t == 0 && i%10 == 1 && i%100 != 11) || (t != 0) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
