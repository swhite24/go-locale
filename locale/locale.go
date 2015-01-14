/**
 * locale.go
 *
 * (C) Steven White 2015
 */

// Package locale provides logic for parsing and interacting with the
// Accept-Language HTTP header.
package locale

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type (
	// Locale represents an individual locale
	Locale struct {
		Language string
		Country  string
		Score    float64
	}

	// Locales represents a collection of *Locales, providing utility
	// methods for interacting with them
	Locales []*Locale
)

// NewLocale delivers a reference to a Locale for an individual
// locale definition in an Accept-Language header.
//
// Example definitions:
//
//   en-US
//   en_GB
//   en;q=0.8
func NewLocale(l string) (*Locale, error) {
	// Break l into language / country and q
	parts := strings.Split(l, ";")

	// Grab language and country, with country being optional
	lc := lcRegex.FindAllStringSubmatch(parts[0], -1)
	if len(lc) == 0 {
		return nil, errors.New("No language or country provided")
	}

	// Initialize locale with language / country
	locale := &Locale{Language: lc[0][0]}
	locale.Language = lc[0][0]
	if len(lc) > 1 {
		locale.Country = lc[1][0]
	}

	// Determine if score is specified
	if len(parts) > 1 {
		score := qRegex.FindAllStringSubmatch(parts[1], -1)
		if len(score) > 0 {
			locale.Score, _ = strconv.ParseFloat(score[0][0], 64)
		}
	}

	// Default score
	if locale.Score == 0 {
		locale.Score = 1.0
	}

	return locale, nil
}

// Read receives a complete Accept-Language header and delivers a reference
// to a Locales, containing references to a *Locale for each definition in
// the Accept-Language header.
func Read(header string) Locales {
	// Initialize Locales
	locales := Locales{}

	// Split out individual definitions
	pieces := strings.Split(header, ",")

	// Add Locale for each definition
	for _, l := range pieces {
		if locale, err := NewLocale(l); err == nil {
			locales = append(locales, locale)
		}
	}

	return locales
}

// Best delivers the locale with the highest quality score
func (ls Locales) Best() *Locale {
	best := (*Locale)(nil)
	score := 0.0
	for _, l := range ls {
		if best == nil || l.Score > score {
			best = l
			score = l.Score
		}
	}
	return best
}

func (l *Locale) String() string {
	s := l.Language
	if l.Country != "" {
		s = s + "_" + l.Country
	}

	return s
}

// --------------------------------------------------------
// Locale utilities
// --------------------------------------------------------

var (
	lcRegex = regexp.MustCompile("[a-zA-Z]+")
	qRegex  = regexp.MustCompile("[0-9]+.?[0-9]*")
)
