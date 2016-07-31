package chr_US

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type chr_US struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'chr_US' locale
func New() locales.Translator {
	return &chr_US{
		locale:  "chr_US",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *chr_US) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'chr_US'
func (t *chr_US) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'chr_US'
func (t *chr_US) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
