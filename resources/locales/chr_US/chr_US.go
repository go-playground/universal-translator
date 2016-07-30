package chr_US

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type chr_US struct {
	locale string
}

// New returns a new instance of translator for the 'chr_US' locale
func New() locales.Translator {
	return &chr_US{
		locale: "chr_US",
	}
}

// Locale returns the current translators string locale
func (l *chr_US) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *chr_US) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
