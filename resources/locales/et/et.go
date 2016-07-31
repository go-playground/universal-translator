package et

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type et struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'et' locale
func New() locales.Translator {
	return &et{
		locale:  "et",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *et) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'et'
func (t *et) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'et'
func (t *et) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
