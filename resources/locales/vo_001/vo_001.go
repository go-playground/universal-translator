package vo_001

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type vo_001 struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'vo_001' locale
func New() locales.Translator {
	return &vo_001{
		locale:  "vo_001",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *vo_001) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'vo_001'
func (t *vo_001) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'vo_001'
func (t *vo_001) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
