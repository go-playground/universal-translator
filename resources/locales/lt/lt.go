package lt

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type lt struct {
	locale string
}

// New returns a new instance of translator for the 'lt' locale
func New() locales.Translator {
	return &lt{
		locale: "lt",
	}
}

// Locale returns the current translators string locale
func (l *lt) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *lt) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	f, err := locales.F(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n%10 == 1 && n%100 < 11 && n%100 > 19 {
		return locales.PluralRuleOne, nil
	} else if n%10 >= 2 && n%10 <= 9 && n%100 < 11 && n%100 > 19 {
		return locales.PluralRuleFew, nil
	} else if f != 0 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
