package ur_PK

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ur_PK struct {
	locale string
}

// New returns a new instance of translator for the 'ur_PK' locale
func New() locales.Translator {
	return &ur_PK{
		locale: "ur_PK",
	}
}

// Locale returns the current translators string locale
func (l *ur_PK) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ur_PK) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	v := locales.V(num)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
