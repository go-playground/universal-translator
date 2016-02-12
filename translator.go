package ut

import (
	"fmt"
	"log"
)

type translation struct {
	text string
}

// map[key]map[plural type othe, many, few, single]*translation
type translations map[PluralRule]map[string]*translation
type groups map[string][]*translation

// Translator holds the locale translation instance
type Translator struct {
	Locale       *Locale
	ruler        PluralRuler
	translations translations
	groups       groups
}

func newTranslator(locale string) (*Translator, error) {
	return nil, nil
	// loc, err := GetLocale(locale)
	// if err != nil {
	// 	return nil, err
	// }

	// return &Translator{
	// 	Locale:       loc,
	// 	ruler:        pluralRules[loc.PluralRule],
	// 	translations: make(translations),
	// 	groups:       make(groups),
	// }, nil
}

// Add registers a new translation to the Translator using the
// key for a unique translation and group allows for retrieving
// translations for dev purposes later i.e. load all translations
// for the "homepage" group
// plural is optional, can be blank, but allows adding a plural translation
// at the same time as a singular
func (t *Translator) Add(rule PluralRule, group string, key string, text string) {

	trans := &translation{
		text: text,
	}

	if _, ok := t.translations[rule][key]; ok {
		panic(fmt.Sprintf("Translation with key '%s' already exists", key))
	}

	t.translations[rule][key] = trans

	if _, ok := t.groups[group]; !ok {
		t.groups[group] = make([]*translation, 0)
	}

	t.groups[group] = append(t.groups[group], trans)
}

func (t *Translator) pluralRule(count NumberValue) (rule PluralRule) {
	return t.ruler.FindRule(count)
}

// T translates the text associated with the given key with the
// arguments passed in
func (t *Translator) T(key string, a ...interface{}) string {
	return t.P(key, 0, a...)
}

// P translates the plural text associated with the given key with the
// arguments passed in
func (t *Translator) P(key string, count interface{}, a ...interface{}) string {

	rule := t.ruler.FindRule(count)

	trans, ok := t.translations[rule][key]
	if !ok {
		s := "***** WARNING:***** Translation Key " + key + " Not Found"
		log.Println(s)
		return s
	}

	return fmt.Sprintf(trans.text, a...)
}

// // TSafe translates the text associated with the given key with the
// // arguments passed in just like T() but doesn't panic, but instead
// // returns an error
// func (t *Translator) TSafe(key string, a ...interface{}) (string, error) {

// 	trans, ok := t.translations[key]
// 	if !ok {
// 		return "", errors.New("*** Translation Key " + key + " Not Found")
// 	}

// 	return fmt.Sprintf(trans.singular, a...), nil
// }
