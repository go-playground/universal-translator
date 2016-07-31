package ar_SD

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ar_SD struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ar_SD' locale
func New() locales.Translator {
	return &ar_SD{
		locale:  "ar_SD",
		plurals: []locales.PluralRule{1, 2, 3, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *ar_SD) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ar_SD'
func (t *ar_SD) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ar_SD'
func (t *ar_SD) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
	} else if n%100 >= 3 && n%100 <= 10 {
		return locales.PluralRuleFew, nil
	} else if n%100 >= 11 && n%100 <= 99 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
