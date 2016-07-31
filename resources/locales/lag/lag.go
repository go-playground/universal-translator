package lag

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lag struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'lag' locale
func New() locales.Translator {
	return &lag{
		locale:  "lag",
		plurals: []locales.PluralRule{1, 2, 6},
	}
}

// Locale returns the current translators string locale
func (t *lag) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'lag'
func (t *lag) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'lag'
func (t *lag) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 0 {
		return locales.PluralRuleZero, nil
	} else if (i == 0 || i == 1) && n != 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
