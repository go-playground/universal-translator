package gl_ES

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type gl_ES struct {
	locale  string
	plurals []locales.PluralRule
}

// New returns a new instance of translator for the 'gl_ES' locale
func New() locales.Translator {
	return &gl_ES{
		locale:  "gl_ES",
		plurals: []locales.PluralRule{2, 6},
	}
}

// Locale returns the current translators string locale
func (t *gl_ES) Locale() string {
	return t.locale
}

// Plurals returns the list of plurals associated with 'gl_ES'
func (t *gl_ES) Plurals() []locales.PluralRule {
	return t.plurals
}

// CardinalPluralRule returns the PluralRule given 'num' for 'gl_ES'
func (t *gl_ES) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
