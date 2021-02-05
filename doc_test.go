package dontpanic

import (
	"testing"
)

func TestRender(t *testing.T) {
	if err = Render("README.md", "index.html"); err != nil {
		t.Errorf("Expected nil; Got: %s", err.Error())
	}
}
