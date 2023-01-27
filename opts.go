package main

import "errors"

type opts struct {
	heads, remotes bool
	base           string
}

func parseOpts(args []string) (*opts, error) {
	res := &opts{heads: true}
	for len(args) > 0 {
		arg := args[0]
		args = args[1:]
		switch {
		case arg == "-r" || arg == "--remotes":
			res.heads, res.remotes = false, true
		case arg == "-a" || arg == "--all":
			res.heads, res.remotes = true, true
		case arg == "--base":
			if len(args) == 0 {
				return nil, errors.New("option `base` requires a value")
			}
			res.base = args[0]
			args = args[1:]
		}
	}
	return res, nil
}
