package ig_NG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ig_NG struct {
	locale string
}

// New returns a new instance of translator for the 'ig_NG' locale
func New() locales.Translator {
	return &ig_NG{
		locale: "ig_NG",
	}
}

// Locale returns the current translators string locale
func (l *ig_NG) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ig_NG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
