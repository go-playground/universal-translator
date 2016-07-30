package cy

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type cy struct {
	locale string
}

// New returns a new instance of translator for the 'cy' locale
func New() locales.Translator {
	return &cy{
		locale: "cy",
	}
}

// Locale returns the current translators string locale
func (l *cy) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *cy) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 0 {
		return locales.PluralRuleZero, nil
	} else if n == 1 {
		return locales.PluralRuleOne, nil
	} else if n == 2 {
		return locales.PluralRuleTwo, nil
	} else if n == 3 {
		return locales.PluralRuleFew, nil
	} else if n == 6 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
