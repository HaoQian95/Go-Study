package test

import "testing"

func TestPalindrome(t *testing.T) {
	if !IsPalindrome("été") {
		t.Error(`IsPalindrome("été") = false!`)
	}
	if !IsPalindrome("kayak")  {
		t.Error(`IsPalindrome("kayak") = false!`)
	}
}

func TestNonPalindrome(t *testing.T) {
	input := "A man, a plan, a canal: Panama"
    if IsPalindrome(input) {
        t.Errorf(`IsPalindrome(%q) = true!`, input)
    }
}
