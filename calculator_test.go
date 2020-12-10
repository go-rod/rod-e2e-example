package main

// test case: 1 + 2 = 3
func (t T) Add() {
	p := t.page("https://ahfarmer.github.io/calculator")

	p.MustElementR("button", "1").MustClick()
	p.MustElementR("button", `^\+$`).MustClick()
	p.MustElementR("button", "2").MustClick()
	p.MustElementR("button", "=").MustClick()

	// assert the result with t.Eq
	t.Eq(p.MustElement(".component-display").MustText(), "3")
}

// test case: 2 * 3 = 6
func (t T) Multiple() {
	p := t.page("https://ahfarmer.github.io/calculator")

	// use for-loop to click each button
	for _, regex := range []string{"2", "x", "3", "="} {
		p.MustElementR("button", regex).MustClick()
	}

	t.Eq(p.MustElement(".component-display").MustText(), "6")
}
