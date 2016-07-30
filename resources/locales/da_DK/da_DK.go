package da_DK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type da_DK struct {
	locale string
}

// New returns a new instance of translator for the 'da_DK' locale
func New() locales.Translator {
	return &da_DK{
		locale: "da_DK",
	}
}

// Locale returns the current translators string locale
func (l *da_DK) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *da_DK) CardinalPluralRule(num string) (locales.PluralRule, error) {

	t, err := locales.T(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (n == 1) || (t != 0 && (i == 0 || i == 1)) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
