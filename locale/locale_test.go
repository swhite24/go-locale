package locale

import "testing"

func TestNewLocale(t *testing.T) {
	l, e := NewLocale("en-US")

	if e != nil {
		t.Error("Failed to create new locale with proper format")
	}

	if l.Language != "en" {
		t.Errorf("Failed to populate locale with correct language.  Want %s, got %s.", "en", l.Language)
	}
	if l.Country != "US" {
		t.Errorf("Failed to populate locale with correct language.  Want %s, got %s.", "US", l.Country)
	}
}

func TestNewLocaleOnlyLanguage(t *testing.T) {
	l, e := NewLocale("en")

	if e != nil {
		t.Error("Failed to create new locale with proper format")
	}

	if l.Language != "en" {
		t.Errorf("Failed to populate locale with correct language.  Want %s, got %s.", "en", l.Language)
	}
	if l.Country != "" {
		t.Errorf("Failed to populate locale with empty language.  Got %s.", l.Country)
	}
}

func TestNewLocaleWithQuality(t *testing.T) {
	l, e := NewLocale("en-US;q=0.9")
	if e != nil {
		t.Error("Failed to create new locale with proper format")
	}

	if l.Language != "en" {
		t.Errorf("Failed to populate locale with correct language.  Want %s, got %s.", "en", l.Language)
	}
	if l.Country != "US" {
		t.Errorf("Failed to populate locale with correct language.  Want %s, got %s.", "US", l.Country)
	}
	if l.Score != 0.9 {
		t.Errorf("Failed to populate score with correct value. Want %f, got %f.", 0.9, l.Score)
	}
}

func TestNewLocaleEmpty(t *testing.T) {
	l, e := NewLocale("")

	if l != nil || e == nil {
		t.Error("Incorrectly created new locale with empty string")
	}
}

func TestRead(t *testing.T) {
	ls := Read("en_US,en;q=0.8")

	if len(ls) != 2 {
		t.Errorf("Failed to generate locale for each definition")
	}
}

func TestBest(t *testing.T) {
	ls := Read("en_US,en;q=0.8")

	l := ls.Best()

	if l.Country != "US" || l.Language != "en" || l.Score != 1 {
		t.Errorf("Failed to select best language")
	}
}
