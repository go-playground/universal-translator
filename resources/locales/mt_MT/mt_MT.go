package mt_MT

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type mt_MT struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'mt_MT' locale
func New() locales.Translator {
	return &mt_MT{
		locale:  "mt_MT",
		plurals: []locales.PluralRule{2, 4, 5, 6},
	}
}

// Locale returns the current translators string locale
func (t *mt_MT) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'mt_MT'
func (t *mt_MT) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'mt_MT'
func (t *mt_MT) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	} else if (n == 0) || (n%100 >= 2 && n%100 <= 10) {
		return locales.PluralRuleFew, nil
	} else if n%100 >= 11 && n%100 <= 19 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
