package ff_GN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ff_GN struct {
	locale string
}

// New returns a new instance of translator for the 'ff_GN' locale
func New() locales.Translator {
	return &ff_GN{
		locale: "ff_GN",
	}
}

// Locale returns the current translators string locale
func (l *ff_GN) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ff_GN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	i, err := locales.I(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if i == 0 || i == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
