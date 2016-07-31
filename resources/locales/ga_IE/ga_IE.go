package ga_IE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ga_IE struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ga_IE' locale
func New() locales.Translator {
	return &ga_IE{
		locale:  "ga_IE",
		plurals: []locales.PluralRule{2, 3, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *ga_IE) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ga_IE'
func (t *ga_IE) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ga_IE'
func (t *ga_IE) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	} else if n == 2 {
		return locales.PluralRuleTwo, nil
	} else if n >= 3 && n <= 6 {
		return locales.PluralRuleFew, nil
	} else if n >= 7 && n <= 10 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
