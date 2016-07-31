package smn_FI

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type smn_FI struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'smn_FI' locale
func New() locales.Translator {
	return &smn_FI{
		locale:  "smn_FI",
		plurals: []locales.PluralRule{2, 3, 6},
	}
}

// Locale returns the current translators string locale
func (t *smn_FI) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'smn_FI'
func (t *smn_FI) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'smn_FI'
func (t *smn_FI) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
