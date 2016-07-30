package ms_BN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ms_BN struct {
	locale string
}

// New returns a new instance of translator for the 'ms_BN' locale
func New() locales.Translator {
	return &ms_BN{
		locale: "ms_BN",
	}
}

// Locale returns the current translators string locale
func (l *ms_BN) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ms_BN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
