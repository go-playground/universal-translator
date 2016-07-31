package naq_NA

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type naq_NA struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'naq_NA' locale
func New() locales.Translator {
	return &naq_NA{
		locale:  "naq_NA",
		plurals: []locales.PluralRule{2, 3, 6},
	}
}

// Locale returns the current translators string locale
func (t *naq_NA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'naq_NA'
func (t *naq_NA) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'naq_NA'
func (t *naq_NA) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	} else if n == 2 {
		return locales.PluralRuleTwo, nil
	}

	return locales.PluralRuleOther, nil
}
