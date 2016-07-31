package sl

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type sl struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'sl' locale
func New() locales.Translator {
	return &sl{
		locale:  "sl",
		plurals: []locales.PluralRule{2, 3, 4, 6},
	}
}

// Locale returns the current translators string locale
func (t *sl) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'sl'
func (t *sl) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'sl'
func (t *sl) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if v == 0 && i%100 == 1 {
		return locales.PluralRuleOne, nil
	} else if v == 0 && i%100 == 2 {
		return locales.PluralRuleTwo, nil
	} else if (v == 0 && i%100 >= 3 && i%100 <= 4) || (v != 0) {
		return locales.PluralRuleFew, nil
	}

	return locales.PluralRuleOther, nil
}
