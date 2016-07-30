package ff_MR

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ff_MR struct {
	locale string
}

// New returns a new instance of translator for the 'ff_MR' locale
func New() locales.Translator {
	return &ff_MR{
		locale: "ff_MR",
	}
}

// Locale returns the current translators string locale
func (l *ff_MR) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ff_MR) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
