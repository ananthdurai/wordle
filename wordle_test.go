package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchWithLetterPosition(t *testing.T) {
	wordList := []string{"DAINE", "AIERY", "BLING", "CHIMP", "DRUID"}
	expected := []string{"DAINE", "AIERY", "CHIMP", "DRUID"}
	actual := MatchLetterWithPosition(wordList, "L", 1)
	assert.Equal(t, expected, actual, "The two words should be the same.")
}

func TestMatchAnyLetter(t *testing.T) {
	wordList := []string{"DAINE", "AIERY", "BLING", "CHIMP", "DRUID"}
	expected := []string{"DAINE", "BLING", "CHIMP"}
	actual := MatchLetterWithAny(wordList, "R")
	assert.Equal(t, expected, actual, "The two words should be the same.")
}
