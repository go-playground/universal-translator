package vi

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type vi struct {
	locale string
}

// New returns a new instance of translator for the 'vi' locale
func New() locales.Translator {
	return &vi{
		locale: "vi",
	}
}

// Locale returns the current translators string locale
func (l *vi) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *vi) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
