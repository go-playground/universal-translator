package af_NA

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type af_NA struct {
	locale string
}

// New returns a new instance of translator for the 'af_NA' locale
func New() locales.Translator {
	return &af_NA{
		locale: "af_NA",
	}
}

// Locale returns the current translators string locale
func (l *af_NA) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *af_NA) CardinalPluralRule(num string) (locales.PluralRule, error) {

	n, err := locales.N(num)
	if err != nil {
		return locales.PluralRuleUnknown, &locales.ErrBadNumberValue{NumberValue: num, InnerError: err}
	}

	if n == 1 {
		return locales.PluralRuleOne, nil
	}

	return locales.PluralRuleOther, nil
}
