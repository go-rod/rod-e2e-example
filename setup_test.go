// This is the setup file for this test suite.

package main

import (
	"testing"

	"github.com/go-rod/rod"
	"github.com/ysmood/got"
)

// entry point for tests
func Test(t *testing.T) {
	browser := rod.New().MustConnect()

	got.Each(t, func(t *testing.T) T {
		t.Parallel() // run each test concurrently

		return T{got.New(t), browser}
	})
}

// test context
type T struct {
	got.G

	browser *rod.Browser
}

// a helper function to create an incognito page
func (t T) page(url string) *rod.Page {
	page := t.browser.MustIncognito().MustPage(url)
	t.Cleanup(page.MustClose)
	return page
}
