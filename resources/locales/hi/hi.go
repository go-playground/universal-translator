package hi

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type hi struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'hi' locale
func New() locales.Translator {
	return &hi{
		locale:  "hi",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *hi) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'hi'
func (t *hi) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'hi'
func (t *hi) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (i == 0) || (n == 1) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
