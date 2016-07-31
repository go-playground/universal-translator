package he

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type he struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'he' locale
func New() locales.Translator {
	return &he{
		locale:  "he",
		plurals: []locales.PluralRule{2, 3, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *he) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'he'
func (t *he) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'he'
func (t *he) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	} else if i == 2 && v == 0 {
		return locales.PluralRuleTwo, nil
	} else if v == 0 && n < 0 && n > 10 && n%10 == 0 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
