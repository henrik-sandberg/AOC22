package main

import (
	"fmt"
	"strings"
)

type game_option struct {
	name  string
	beats string
	loses string
	value int
}

func Day02(input []string) {

	rock := game_option{name: "rock", beats: "scissor", loses: "paper", value: 1}
	paper := game_option{name: "paper", beats: "rock", loses: "scissor", value: 2}
	scissor := game_option{name: "scissor", beats: "paper", loses: "rock", value: 3}

	game_options := make(map[string]game_option)
	game_options[rock.name] = rock
	game_options[paper.name] = paper
	game_options[scissor.name] = scissor

	fmt.Println("Part 1: ", day02_part1(input, game_options))
	fmt.Println("Part 2: ", day02_part2(input, game_options))
}

func day02_part1(input []string, game_options map[string]game_option) (score int) {
	mappings := make(map[string]string)
	mappings["A"] = "rock"
	mappings["B"] = "paper"
	mappings["C"] = "scissor"

	mappings["X"] = "rock"
	mappings["Y"] = "paper"
	mappings["Z"] = "scissor"

	for _, round := range input {
		s := strings.Split(round, " ")
		score += calculate_score(game_options[mappings[s[1]]], game_options[mappings[s[0]]])
	}
	return
}

func day02_part2(input []string, game_options map[string]game_option) (score int) {
	mappings := make(map[string]string)
	mappings["A"] = "rock"
	mappings["B"] = "paper"
	mappings["C"] = "scissor"

	for _, round := range input {
		s := strings.Split(round, " ")
		they := game_options[mappings[s[0]]]
		var us game_option
		if s[1] == "X" {
			us = game_options[they.beats]
		} else if s[1] == "Y" {
			us = game_options[they.name]
		} else {
			us = game_options[they.loses]
		}
		score += calculate_score(us, they)
	}
	return
}

func calculate_score(us game_option, they game_option) int {
	// Score: 0 if lost, 3 if draw, 6 if win
	// plus value
	if us.name == they.name {
		return 3 + us.value
	} else if they.name == us.beats {
		return 6 + us.value
	} else {
		return us.value
	}
}
