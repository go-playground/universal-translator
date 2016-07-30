package gsw_FR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type gsw_FR struct {
	locale string
}

// New returns a new instance of translator for the 'gsw_FR' locale
func New() locales.Translator {
	return &gsw_FR{
		locale: "gsw_FR",
	}
}

// Locale returns the current translators string locale
func (l *gsw_FR) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *gsw_FR) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
