package he_IL

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type he_IL struct {
	locale string
}

// New returns a new instance of translator for the 'he_IL' locale
func New() locales.Translator {
	return &he_IL{
		locale: "he_IL",
	}
}

// Locale returns the current translators string locale
func (l *he_IL) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *he_IL) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	} else if i == 2 && v == 0 {
		return locales.PluralRuleTwo, nil
	} else if v == 0 && n < 0 && n > 10 && n%10 == 0 {
		return locales.PluralRuleMany, nil
	}

	return locales.PluralRuleOther, nil
}
