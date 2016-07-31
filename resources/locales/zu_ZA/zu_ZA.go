package zu_ZA

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zu_ZA struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zu_ZA' locale
func New() locales.Translator {
	return &zu_ZA{
		locale:  "zu_ZA",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *zu_ZA) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zu_ZA'
func (t *zu_ZA) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'zu_ZA'
func (t *zu_ZA) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
