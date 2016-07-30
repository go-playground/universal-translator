package ses

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ses struct {
	locale string
}

// New returns a new instance of translator for the 'ses' locale
func New() locales.Translator {
	return &ses{
		locale: "ses",
	}
}

// Locale returns the current translators string locale
func (l *ses) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ses) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
