package ne_NP

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ne_NP struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ne_NP' locale
func New() locales.Translator {
	return &ne_NP{
		locale:  "ne_NP",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ne_NP) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ne_NP'
func (t *ne_NP) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ne_NP'
func (t *ne_NP) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
