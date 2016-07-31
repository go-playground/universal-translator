package ff

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ff struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'ff' locale
func New() locales.Translator {
	return &ff{
		locale:  "ff",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *ff) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'ff'
func (t *ff) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'ff'
func (t *ff) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
