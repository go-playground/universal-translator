package tzm_MA

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type tzm_MA struct {
	locale string
}

// New returns a new instance of translator for the 'tzm_MA' locale
func New() locales.Translator {
	return &tzm_MA{
		locale: "tzm_MA",
	}
}

// Locale returns the current translators string locale
func (l *tzm_MA) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *tzm_MA) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if (n >= 0 && n <= 1) || (n >= 11 && n <= 99) {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
