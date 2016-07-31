package as

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type as struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'as' locale
func New() locales.Translator {
	return &as{
		locale:  "as",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *as) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'as'
func (t *as) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'as'
func (t *as) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
