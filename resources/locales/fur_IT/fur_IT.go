package fur_IT

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fur_IT struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'fur_IT' locale
func New() locales.Translator {
	return &fur_IT{
		locale:  "fur_IT",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *fur_IT) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'fur_IT'
func (t *fur_IT) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'fur_IT'
func (t *fur_IT) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
