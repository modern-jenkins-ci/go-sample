package main

import (
	"testing"
)

var DockerHubUser string
var DockerHubPassword string
var AwsAccess string
var AwsSecret string

func TestPass(t *testing.T) {
	if !PassTest() {
		t.Error("expected test to pass.")
	}
}

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func assertNotEqual(t *testing.T, a interface{}, b interface{}) {
	if a == b {
		t.Fatalf("%s == %s", a, b)
	}
}
