package ff_CM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ff_CM struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ff_CM' locale
func New() locales.Translator {
	return &ff_CM{
		locale:  "ff_CM",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ff_CM) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ff_CM'
func (t *ff_CM) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ff_CM'
func (t *ff_CM) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
