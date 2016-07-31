package da

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type da struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'da' locale
func New() locales.Translator {
	return &da{
		locale:  "da",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *da) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'da'
func (t *da) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'da'
func (t *da) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	t, err := locales.T(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (n == 1) || (t != 0 && (i == 0 || i == 1)) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
