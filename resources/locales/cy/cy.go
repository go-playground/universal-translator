package cy

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type cy struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'cy' locale
func New() locales.Translator {
	return &cy{
		locale:  "cy",
		plurals: []locales.PluralRule{1, 2, 3, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *cy) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'cy'
func (t *cy) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'cy'
func (t *cy) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 0 {
		return locales.PluralRuleZero, nil
	} else if n == 1 {
		return locales.PluralRuleOne, nil
	} else if n == 2 {
		return locales.PluralRuleTwo, nil
	} else if n == 3 {
		return locales.PluralRuleFew, nil
	} else if n == 6 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
