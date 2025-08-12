package main

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

func die(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
}

func main() {
	o, err := parseOpts(os.Args[1:])
	die(err)

	r, err := git.PlainOpen("")
	die(err)

	head, err := r.Head()
	die(err)
	var base *plumbing.Reference
	if o.base == "" {
		base = head
	} else {
		branch, err := r.Branch(o.base)
		die(err)
		base, err = r.Reference(plumbing.ReferenceName(branch.Name), true)
		die(err)
	}

	baseComit, err := r.CommitObject(base.Hash())
	die(err)

	refs, _ := r.References()
	var report []item
	refs.ForEach(func(ref *plumbing.Reference) error {
		if (ref.Name().IsBranch() && !o.heads) || (ref.Name().IsRemote() && !o.remotes) || (!ref.Name().IsBranch() && !ref.Name().IsRemote()) {
			return nil
		}
		if ref.Hash().IsZero() {
			return nil
		}

		commit, err := r.CommitObject(ref.Hash())
		die(err)
		ancestors, _ := commit.MergeBase(baseComit)
		ahead, _ := walk(commit, ancestors)
		behind, _ := walk(baseComit, ancestors)
		var name string
		if o.heads && ref.Name().IsRemote() {
			name = fmt.Sprintf("remotes/%s", ref.Name().Short())
		} else {
			name = ref.Name().Short()
		}
		report = append(report, item{
			name:  name,
			h:     ref.Hash(),
			ahead: ahead, behind: behind,
			current: *ref == *head,
		})
		return nil
	})

	if len(report) == 0 {
		return
	}
	var nameW int
	for _, i := range report {
		if len(i.name) > nameW {
			nameW = len(i.name)
		}
	}
	for _, i := range report {
		if i.current {
			fmt.Print("* ")
		} else {
			fmt.Print("  ")
		}
		fmt.Printf("%-*s  %s", nameW, i.name, i.h.String()[:8])
		fmt.Printf("  %3d", i.behind)
		if i.behind > 100 {
			fmt.Print("━")
		} else if i.behind > 10 {
			fmt.Print("╺")
		} else {
			fmt.Print(" ")
		}
		if i.behind > 0 && i.ahead > 0 {
			fmt.Print("┿")
		} else if i.behind > 0 {
			fmt.Print("┥")
		} else if i.ahead > 0 {
			fmt.Print("┝")
		} else {
			fmt.Print("│")
		}
		if i.ahead > 100 {
			fmt.Print("━")
		} else if i.ahead > 10 {
			fmt.Print("╸")
		} else {
			fmt.Print(" ")
		}
		fmt.Printf("%-3d", i.ahead)
		fmt.Println()
	}
}

type item struct {
	name          string
	h             plumbing.Hash
	ahead, behind uint
	current       bool
}

func walk(leaf *object.Commit, roots []*object.Commit) (uint, error) {
	var count uint
	todo := []*object.Commit{leaf}
	cache := make(map[plumbing.Hash]bool, len(roots))
	for _, c := range roots {
		cache[c.Hash] = true
	}
	for len(todo) > 0 {
		c := todo[0]
		todo = todo[1:]
		if cache[c.Hash] {
			continue
		}
		count += 1
		c.Parents().ForEach(func(p *object.Commit) error {
			todo = append(todo, p)
			return nil
		})
		cache[c.Hash] = true
	}
	return count, nil
}

type walker struct {
	seen map[plumbing.Hash]bool
}
