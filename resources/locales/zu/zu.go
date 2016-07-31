package zu

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type zu struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'zu' locale
func New() locales.Translator {
	return &zu{
		locale:  "zu",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *zu) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'zu'
func (t *zu) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'zu'
func (t *zu) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
