package et_EE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type et_EE struct {
	locale string
}

// New returns a new instance of translator for the 'et_EE' locale
func New() locales.Translator {
	return &et_EE{
		locale: "et_EE",
	}
}

// Locale returns the current translators string locale
func (l *et_EE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *et_EE) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
