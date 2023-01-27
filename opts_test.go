package main

import "testing"

func Test_parseOpts(t *testing.T) {
	for _, tst := range []struct {
		name     string
		args     []string
		expected *opts
		error    string
	}{
		{name: "defaults", args: []string{}, expected: &opts{true, false, ""}},
		{name: "--remotes", args: []string{"--remotes"}, expected: &opts{false, true, ""}},
		{name: "-r", args: []string{"-r"}, expected: &opts{false, true, ""}},
		{name: "--all", args: []string{"--all"}, expected: &opts{true, true, ""}},
		{name: "-a", args: []string{"-a"}, expected: &opts{true, true, ""}},
		{name: "--base foo", args: []string{"--base", "foo"}, expected: &opts{true, false, "foo"}},
		{name: "--base", args: []string{"--base"}, error: "option `base` requires a value"},
	} {
		t.Run(tst.name, func(t *testing.T) {
			actual, err := parseOpts(tst.args)
			if tst.expected != nil && *tst.expected != *actual {
				t.Fatalf("expected %v, got %v", tst.expected, actual)
			}
			if tst.error != "" && err.Error() != tst.error {
				t.Fatalf("expected error \"%v\", got \"%v\"", tst.error, err.Error())
			}
		})
	}
}
