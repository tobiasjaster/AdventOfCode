// Code generated by aocgen; DO NOT EDIT.
package year2022

import (
	"fmt"
	"math"
	"crypto/md5"
	"encoding/json"
)

type Day19 struct{}

type Material int 

const (
	Ore 	 Material = 0
	Clay 	 Material = 1
	Obsidian Material = 2
	Geode 	 Material = 3
)

type RobotFactory struct{
	Robots map[Material]map[Material]int
}

type Game19 struct {
	Robots map[Material]int
	Material map[Material]int
	Minutes int
	RoundRobots map[int]map[Material]int
	RoundMaterial map[int]map[Material]int
}

func (g *Game19)produce() {
	material := g.Material
	material[Ore] += g.Robots[Ore]
	material[Clay] += g.Robots[Clay]
	material[Obsidian] += g.Robots[Obsidian]
	material[Geode] += g.Robots[Geode]
	g.Material = material
}

func (g *Game19)maxGeode(maxMinutes int, factory RobotFactory) int {
	geode := g.Material[Geode]
	geode += (maxMinutes-g.Minutes)*g.Robots[Geode]
	for i:=1;i<(maxMinutes-g.Minutes);i +=2 {
		geode += i
	}
	return geode
}

func (g *Game19)getId() [16]byte {
    jsonBytes, _ := json.Marshal(g)
    return md5.Sum(jsonBytes)
}

func NewGame19() *Game19 {
	g := new(Game19)
	g.Robots = map[Material]int{Ore:1,Clay:0,Obsidian:0,Geode:0}
	g.Material = map[Material]int{Ore:0,Clay:0,Obsidian:0,Geode:0}
	roundRobots := make(map[int]map[Material]int)
	roundMaterial := make(map[int]map[Material]int)
	g.Minutes = 0
	g.RoundRobots = roundRobots
	g.RoundMaterial = roundMaterial
	return g
}

func CopyGame19(game Game19) *Game19 {
	g := new(Game19)
	g.Robots = make(map[Material]int)
	for k,v := range game.Robots{
		g.Robots[k] = v
	}
	g.Material = make(map[Material]int)
	for k,v := range game.Material{
		g.Material[k] = v
	}
	roundRobots := make(map[int]map[Material]int)
	for round, robots := range(game.RoundRobots) {
		roundRobots[round] = make(map[Material]int)
		for material, count := range robots{
			roundRobots[round][material] = count
		}
	}
	roundMaterial := make(map[int]map[Material]int)
	for round, robots := range(game.RoundMaterial) {
		roundMaterial[round] = make(map[Material]int)
		for material, count := range robots{
			roundMaterial[round][material] = count
		}
	}
	g.RoundRobots = roundRobots
	g.RoundMaterial = roundMaterial
	g.Minutes = int(game.Minutes)
	return g
}

func parseInput19(lines []string) map[int]RobotFactory {
	robotFactories := make(map[int]RobotFactory)

	for _, line := range lines {
		if line == "" { continue }
		factoryNumber := 0
		var oreRobotMaterialOre int
		var clayRobotMaterialOre int
		var obsidianRobotMaterialOre int
		var obsidianRobotMaterialClay int
		var geodeRobotMaterialOre int
		var geodeRobotMaterialObsidian int
		fmt.Sscanf(
			line,
			"Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&factoryNumber,
			&oreRobotMaterialOre,
			&clayRobotMaterialOre,
			&obsidianRobotMaterialOre,
			&obsidianRobotMaterialClay,
			&geodeRobotMaterialOre,
			&geodeRobotMaterialObsidian,
		)		
		robots := map[Material]map[Material]int{Ore:{Ore:oreRobotMaterialOre}, Clay:{Ore:clayRobotMaterialOre}, Obsidian:{Ore:obsidianRobotMaterialOre, Clay:obsidianRobotMaterialClay}, Geode:{Ore:geodeRobotMaterialOre, Obsidian:geodeRobotMaterialObsidian}}
		robotFactory := RobotFactory{Robots: robots}
		robotFactories[factoryNumber] = robotFactory
	}
	return robotFactories
}

func mostGeode(game Game19, maxMinutes int, factory RobotFactory) int {
	geode := 0
	queue := []Game19{game}
	maxOre := int(math.Max(math.Max(float64(factory.Robots[Ore][Ore]), float64(factory.Robots[Clay][Ore])), math.Max(float64(factory.Robots[Obsidian][Ore]), float64(factory.Robots[Geode][Ore]))))
	visited := make(map[[16]byte]int)
	for ok := true; ok; ok = len(queue)>0 {
		actGame := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if actGame.Material[Geode] > geode {
			fmt.Println(actGame)
			geode = actGame.Material[Geode]
		}
		if actGame.Minutes>=maxMinutes { 
			continue
		}
		if actGame.maxGeode(maxMinutes, factory) <= geode {
			continue
		}
		actGame.Minutes++
		for _, robot := range []Material{Ore, Clay, Obsidian, Geode} {
			nextGame := *CopyGame19(actGame)
			buildable := true
			for _, material := range []Material{Ore, Clay, Obsidian, Geode} {
				if nextGame.Material[material] < factory.Robots[robot][material] {
					buildable = false
					break
				}
			}
			if !buildable { continue }
			switch robot {
			case Ore:
				if nextGame.Robots[Ore] > maxOre {
					continue
				}
			case Clay:
				if nextGame.Robots[Clay] >= factory.Robots[Obsidian][Clay] {
					continue
				}
			case Obsidian:
				if nextGame.Robots[Obsidian] >= factory.Robots[Geode][Obsidian] {
					continue
				}
			}
			materialMap := nextGame.Material
			for _, material := range []Material{Ore, Clay, Obsidian, Geode} {
				materialMap[material] -= factory.Robots[robot][material]
			}
			nextGame.Material = materialMap
			nextGame.produce()
			nextGame.Robots[robot]++
			nextGame.RoundMaterial[nextGame.Minutes] = nextGame.Material
			nextGame.RoundRobots[nextGame.Minutes] = nextGame.Robots
			_, err := visited[nextGame.getId()]
			if !err {
				visited[nextGame.getId()] = nextGame.Material[Geode]
				queue = append(queue, nextGame)
			}
		}
		nextGame := *CopyGame19(actGame)
		nextGame.produce()
		nextGame.RoundMaterial[actGame.Minutes] = nextGame.Material
		nextGame.RoundRobots[actGame.Minutes] = nextGame.Robots
		_, err := visited[nextGame.getId()]
		if !err {
			visited[nextGame.getId()] = nextGame.Material[Geode]
			queue = append(queue, nextGame)
		}
	}
	fmt.Println("Visited: ", len(visited))
	return geode
}


func (p Day19) PartA(lines []string) any {
	robotFactories := parseInput19(lines)
	mostGeodeMap := make(map[int]int)
	geode := 0
	for nr, factory := range robotFactories {
		game := *NewGame19()
		mostGeodeMap[nr] = mostGeode(game, 24, factory)
	}
	fmt.Println(mostGeodeMap)
	for _,v := range mostGeodeMap {
		if v>geode {
			geode = v
		}
	}
	return geode
}

func (p Day19) PartB(lines []string) any {
	return "implement_me"
}
