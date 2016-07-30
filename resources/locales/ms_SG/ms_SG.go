package ms_SG

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ms_SG struct {
	locale string
}

// New returns a new instance of translator for the 'ms_SG' locale
func New() locales.Translator {
	return &ms_SG{
		locale: "ms_SG",
	}
}

// Locale returns the current translators string locale
func (l *ms_SG) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ms_SG) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
