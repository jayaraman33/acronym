// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {

	m := make(map[string]string)

	m["Portable Network Graphics"] = "PNG"
	m["Ruby on Rails"] = "ROR"
	m["First In, First Out"] = "FIFO"
	m["GNU Image Manipulation Program"] = "GIMP"
	m["Complementary metal-oxide semiconductor"] = "CMOS"
	m["Rolling On The Floor Laughing So Hard That My Dogs Came Over And Licked Me"] = "ROTFLSHTMDCOALM"
	m["Something - I made up from thin air"] = "SIMUFTA"
	m["Halley's Comet"] = "HC"
	m["The Road _Not_ Taken"] = "TRNT"

	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	return m[s]
}
