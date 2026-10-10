package main

import "testing"

func TestFeatureSCRUM3(t *testing.T) {
	if FeatureSCRUM3() != "SCRUM-3" {
		t.Fatal("unexpected")
	}
}
