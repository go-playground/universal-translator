package ut

import (
	"testing"

	"github.com/go-playground/locales/en"
)

func TestFindTranslatorTrimSpace(t *testing.T) {
	enLoc := en.New()
	tr := New(enLoc, enLoc)
	got, found := tr.FindTranslator("  en  ")
	if !found || got == nil {
		t.Fatalf("found=%v got=%v", found, got)
	}
	got, found = tr.GetTranslator(" en ")
	if !found || got == nil {
		t.Fatalf("get found=%v", found)
	}
}
