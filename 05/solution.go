package day_05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type ruleset map[int]struct {
	before []int
}

func (r ruleset) isBefore(p, other int) bool {
	if p == other {
		fmt.Println("pages equal")
		return false
	}

	before, ok := r[p]
	if !ok {
		fmt.Println("no rules for page", p)
		return false
	}

	for _, b := range before.before {
		if b == other {
			return true
		}
	}

	return false // or unsure
}

type update []int

func (u update) validate(r ruleset) bool {
	for i := 0; i < len(u)-1; i++ {
		for j := i + 1; j < len(u); j++ {
			if !r.isBefore(u[i], u[j]) {
				return false
			}
		}
	}

	return true
}

func (u update) getMiddlePage() int {
	if len(u)%2 != 1 {
		fmt.Println("invalid update length")
		return -1
	}

	return u[len(u)/2]
}

func (u update) sorted(r ruleset) (new update, didchange bool) {
	if u.validate(r) {
		return u, false
	}

	sorted := u

	slices.SortFunc(sorted, func(i, j int) int {
		if i == j {
			fmt.Println("pages equal")
			return -1
		}

		if r.isBefore(i, j) {
			return -1
		}

		return 1
	})

	return sorted, true
}

func parseRuleset(input []string) (ruleset, error) {
	r := make(ruleset)

	for _, line := range input {
		pages := strings.Split(line, "|")
		if len(pages) != 2 {
			return nil, fmt.Errorf("invalid rule: %s", line)
		}

		before, err := strconv.Atoi(pages[0])
		if err != nil {
			return nil, err
		}

		after, err := strconv.Atoi(pages[1])
		if err != nil {
			return nil, err
		}

		if _, ok := r[before]; !ok {
			r[before] = struct {
				before []int
			}{
				before: []int{after},
			}
		} else {
			r[before] = struct {
				before []int
			}{
				before: append(r[before].before, after),
			}
		}
	}

	return r, nil
}

func parseUpdates(input []string) ([]update, error) {
	var updates []update
	for _, line := range input {
		pages := strings.Split(line, ",")

		var pagesInt []int
		for _, page := range pages {
			p, err := strconv.Atoi(page)
			if err != nil {
				return nil, err
			}

			pagesInt = append(pagesInt, p)
		}

		updates = append(updates, update(pagesInt))
	}

	return updates, nil
}

func parseInput(input string) (ruleset, []update, error) {
	parts := strings.Split(input, "\n\n")
	rulesetLines := strings.Split(parts[0], "\n")
	rules, err := parseRuleset(rulesetLines)
	if err != nil {
		return nil, nil, err
	}

	updatesLines := strings.Split(parts[1], "\n")
	updates, err := parseUpdates(updatesLines)
	if err != nil {
		return nil, nil, err
	}

	return rules, updates, nil
}

func PartOne(input string) (int, error) {
	rules, updates, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	var acc int

	for _, u := range updates {
		if u.validate(rules) {
			acc += u.getMiddlePage()
		}
	}

	return acc, nil
}

func PartTwo(input string) (int, error) {
	rules, updates, err := parseInput(input)
	if err != nil {
		return 0, err
	}

	var acc int

	for _, u := range updates {
		sorted, didChange := u.sorted(rules)
		if didChange {
			acc += sorted.getMiddlePage()
		}
	}

	return acc, nil
}
