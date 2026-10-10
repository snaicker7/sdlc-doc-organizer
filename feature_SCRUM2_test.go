package main

import "testing"

func TestFeatureSCRUM2(t *testing.T) {
	if FeatureSCRUM2() != "SCRUM-2" {
		t.Fatal("unexpected")
	}
}
