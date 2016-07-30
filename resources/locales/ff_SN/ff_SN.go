package ff_SN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ff_SN struct {
	locale string
}

// New returns a new instance of translator for the 'ff_SN' locale
func New() locales.Translator {
	return &ff_SN{
		locale: "ff_SN",
	}
}

// Locale returns the current translators string locale
func (l *ff_SN) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ff_SN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
