package ta_SG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ta_SG struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ta_SG' locale
func New() locales.Translator {
	return &ta_SG{
		locale:  "ta_SG",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ta_SG) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ta_SG'
func (t *ta_SG) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ta_SG'
func (t *ta_SG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
