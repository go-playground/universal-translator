package fil_PH

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type fil_PH struct {
	locale string
}

// New returns a new instance of translator for the 'fil_PH' locale
func New() locales.Translator {
	return &fil_PH{
		locale: "fil_PH",
	}
}

// Locale returns the current translators string locale
func (l *fil_PH) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *fil_PH) CardinalPluralRule(num string) (locales.PluralRule, error) {

	f, err := locales.F(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if (v == 0 && (i == 1 || i == 2 || i == 3)) || (v == 0 && (i%10 != 4 && i%10 != 6 && i%10 != 9)) || (v != 0 && (f%10 != 4 && f%10 != 6 && f%10 != 9)) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
