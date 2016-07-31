package mn_MN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mn_MN struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mn_MN' locale
func New() locales.Translator {
	return &mn_MN{
		locale:  "mn_MN",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *mn_MN) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mn_MN'
func (t *mn_MN) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'mn_MN'
func (t *mn_MN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
