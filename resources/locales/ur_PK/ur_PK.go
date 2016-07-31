package ur_PK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ur_PK struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ur_PK' locale
func New() locales.Translator {
	return &ur_PK{
		locale:  "ur_PK",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ur_PK) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ur_PK'
func (t *ur_PK) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ur_PK'
func (t *ur_PK) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
