package main

import (
	"os"
	"testing"
)

func TestGetEnvSet(t *testing.T) {
	os.Setenv("MYVAR", "MYVALUE")
	if val := GetEnv("MYVAR", "FALLBACK"); val != "MYVALUE" {
		panic("Wrong value")
	}
}

func TestGetEnvNotSet(t *testing.T) {
	os.Clearenv()
	if val := GetEnv("MYVAR", "FALLBACK"); val != "FALLBACK" {
		panic("Wrong value")
	}
}
