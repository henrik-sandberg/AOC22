package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

func Day19(input []string) {
	blueprints := make([]blueprint, len(input))
	for i, line := range input {
		bp, err := parseBlueprint(line)
		if err != nil {
			log.Fatal("Could not parse blueprint", err)
		}
		blueprints[i] = bp
	}
	fmt.Println("Part 1: ", day19_part1(blueprints))
	fmt.Println("Part 2: ", day19_part2(blueprints))
}

func day19_part1(blueprints []blueprint) int {
	result := 0
	for _, bp := range blueprints {
		result += bp.id * bp.calculateHighestGeodeLevel(24)
	}
	return result
}

func day19_part2(blueprints []blueprint) int {
	result := 1
	for _, bp := range blueprints[:3] {
		result *= bp.calculateHighestGeodeLevel(32)
	}
	return result
}

type blueprint struct {
	id                                             int
	oreRobot, clayRobot, obsidianRobot, geodeRobot robotCost
}

type robotCost struct {
	ore, clay, obsidian, geode int
}

func (bp blueprint) calculateHighestGeodeLevel(cycles int) (result int) {
	maxOreCost := max(bp.oreRobot.ore, bp.clayRobot.ore, bp.obsidianRobot.ore, bp.geodeRobot.ore)
	maxClayCost := bp.obsidianRobot.clay
	maxObsidianCost := bp.geodeRobot.obsidian
	queue := []robotState{{oreRobots: 1}}
	for len(queue) > 0 {
		state := queue[0]
		queue = queue[1:]
		possibleFutureStates := []robotState{}
		if state.oreRobots < maxOreCost {
			requiredOre := max(0, bp.oreRobot.ore-state.ore)
			time := (requiredOre+state.oreRobots-1)/state.oreRobots + 1
			if state.cycle+time < cycles {
				cp := state.progress(time)
				cp.ore -= bp.oreRobot.ore
				cp.oreRobots++
				possibleFutureStates = append(possibleFutureStates, cp)
			}
		}
		if state.clayRobots < maxClayCost {
			requiredOre := max(0, bp.clayRobot.ore-state.ore)
			time := (requiredOre+state.oreRobots-1)/state.oreRobots + 1
			if state.cycle+time < cycles {
				cp := state.progress(time)
				cp.ore -= bp.clayRobot.ore
				cp.clayRobots++
				possibleFutureStates = append(possibleFutureStates, cp)
			}
		}
		if state.clayRobots > 0 && state.obsidianRobots < maxObsidianCost {
			requiredOre := max(0, bp.obsidianRobot.ore-state.ore)
			requiredClay := max(0, bp.obsidianRobot.clay-state.clay)
			time := max(
				(requiredOre+state.oreRobots-1)/state.oreRobots,
				(requiredClay+state.clayRobots-1)/state.clayRobots,
			) + 1
			if state.cycle+time < cycles {
				cp := state.progress(time)
				cp.ore -= bp.obsidianRobot.ore
				cp.clay -= bp.obsidianRobot.clay
				cp.obsidianRobots++
				possibleFutureStates = append(possibleFutureStates, cp)
			}
		}
		if state.obsidianRobots > 0 {
			requiredOre := max(0, bp.geodeRobot.ore-state.ore)
			requiredObsidian := max(0, bp.geodeRobot.obsidian-state.obsidian)
			time := max(
				(requiredOre+state.oreRobots-1)/state.oreRobots,
				(requiredObsidian+state.obsidianRobots-1)/state.obsidianRobots,
			) + 1
			if state.cycle+time < cycles {
				cp := state.progress(time)
				cp.ore -= bp.geodeRobot.ore
				cp.obsidian -= bp.geodeRobot.obsidian
				cp.geodeRobots++
				possibleFutureStates = append(possibleFutureStates, cp)
			}
		}
		if len(possibleFutureStates) > 0 {
			queue = append(queue, possibleFutureStates...)
		} else {
			result = max(result, state.geode+(cycles-state.cycle)*state.geodeRobots)
		}
	}
	return
}

func parseBlueprint(line string) (blueprint, error) {
	numberRegex := regexp.MustCompile("\\d+")
	match := numberRegex.FindAllString(line, -1)
	if len(match) != 7 {
		return blueprint{}, errors.New("Unexpected format")
	}
	bp := blueprint{}
	bp.id, _ = strconv.Atoi(match[0])

	oreForOre, _ := strconv.Atoi(match[1])
	bp.oreRobot = robotCost{ore: oreForOre}

	oreForClay, _ := strconv.Atoi(match[2])
	bp.clayRobot = robotCost{ore: oreForClay}

	oreForObsidian, _ := strconv.Atoi(match[3])
	clayForObsidian, _ := strconv.Atoi(match[4])
	bp.obsidianRobot = robotCost{ore: oreForObsidian, clay: clayForObsidian}

	oreForGeode, _ := strconv.Atoi(match[5])
	obsidianForGeode, _ := strconv.Atoi(match[6])
	bp.geodeRobot = robotCost{ore: oreForGeode, obsidian: obsidianForGeode}

	return bp, nil
}

type robotState struct {
	cycle                                              int
	ore, clay, obsidian, geode                         int
	oreRobots, clayRobots, obsidianRobots, geodeRobots int
}

func (state robotState) progress(n int) robotState {
	state.cycle += n
	state.ore += state.oreRobots * n
	state.clay += state.clayRobots * n
	state.obsidian += state.obsidianRobots * n
	state.geode += state.geodeRobots * n
	return state
}
