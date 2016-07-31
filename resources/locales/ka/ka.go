package ka

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ka struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ka' locale
func New() locales.Translator {
	return &ka{
		locale:  "ka",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ka) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ka'
func (t *ka) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ka'
func (t *ka) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
