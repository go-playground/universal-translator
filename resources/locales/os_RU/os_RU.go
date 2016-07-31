package os_RU

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type os_RU struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'os_RU' locale
func New() locales.Translator {
	return &os_RU{
		locale:  "os_RU",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *os_RU) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'os_RU'
func (t *os_RU) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'os_RU'
func (t *os_RU) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
