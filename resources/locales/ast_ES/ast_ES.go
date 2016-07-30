package ast_ES

import (
	"github.com/go-playground/universal-translator/resources/locales"
)

type ast_ES struct {
	locale string
}

// New returns a new instance of translator for the 'ast_ES' locale
func New() locales.Translator {
	return &ast_ES{
		locale: "ast_ES",
	}
}

// Locale returns the current translators string locale
func (l *ast_ES) Locale() string {
	return l.locale
}

// CardinalPluralRule returns the PluralRule given 'num'
func (l *ast_ES) CardinalPluralRule(num string) (locales.PluralRule, error) {

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
