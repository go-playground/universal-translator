package om_KE

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type om_KE struct {
	locale string
}

// New returns a new instance of translator for the 'om_KE' locale
func New() locales.Translator {
	return &om_KE{
		locale: "om_KE",
	}
}

// Locale returns the current translators string locale
func (l *om_KE) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *om_KE) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
