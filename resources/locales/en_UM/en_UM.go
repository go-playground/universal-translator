package en_UM

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type en_UM struct {
	locale string
}

// New returns a new instance of translator for the 'en_UM' locale
func New() locales.Translator {
	return &en_UM{
		locale: "en_UM",
	}
}

// Locale returns the current translators string locale
func (l *en_UM) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *en_UM) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
