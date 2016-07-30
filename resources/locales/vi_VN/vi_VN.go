package vi_VN

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type vi_VN struct {
	locale string
}

// New returns a new instance of translator for the 'vi_VN' locale
func New() locales.Translator {
	return &vi_VN{
		locale: "vi_VN",
	}
}

// Locale returns the current translators string locale
func (l *vi_VN) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *vi_VN) CardinalPluralRule(num string) (locales.PluralRule, error) {

	return locales.PluralRuleOther, nil
}
