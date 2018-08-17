package simulation

import (
	"fmt"
)

// Actor represents a single actor in the simulation.
type Actor struct {
	name  string
	attrs map[string]int
	team  int // The team number that the Actor will be on
	x     int // The actor's current x coordinate
	y     int // The actor's current y coordinate
}

// GlobalState represents the global simulation state.
type GlobalState struct {
	actors []Actor
	x      int // The max x coordinate
	y      int // The max y coordinate
}

// NewGlobalState creates a fully built Simulation object that can be used to run the simulation against.
func NewGlobalState() *GlobalState {
	s := new(GlobalState)
	s.actors = []Actor{
		{
			name: "fred",
			attrs: map[string]int{
				"hp":  5,
				"ac":  13,
				"atk": 2,
				"dmg": 4,
			},
			team: 1,
			x:    0,
			y:    1,
		},
		{
			name: "bob",
			attrs: map[string]int{
				"hp":  10,
				"ac":  10,
				"atk": 2,
				"dmg": 4,
			},
			team: 2,
			x:    3,
			y:    2,
		},
	}
	s.x = 5
	s.y = 5
	return s
}

// TeamHealthPools gets the list of team health pools
func TeamHealthPools(s *GlobalState) map[int]int {
	teamHealth := make(map[int]int)
	for _, actor := range s.actors {
		teamHealth[actor.team] += actor.attrs["hp"]
	}
	return teamHealth
}

const maxRounds = 10

// RunSimulation main entry point.
func RunSimulation() {
	// runs until end condition is reached
	fmt.Println("simulation start")
	s := NewGlobalState()
	curTurn := 0
	round := 0
	teamsHealth := TeamHealthPools(s)

	// Loop the simulation until only one team
	for len(teamsHealth) != 1 && round < maxRounds {
		if curTurn == 0 {
			round++
		}
		fmt.Printf("round %d, current actor: %s\n", round, s.actors[curTurn].name)

		teamsHealth = TeamHealthPools(s)
		fmt.Printf("team healths: %v\n", teamsHealth)

		// Set next player's turn.
		curTurn = (curTurn + 1) % len(s.actors)
	}
	if len(teamsHealth) > 1 {
		fmt.Printf("tied: %v\n", teamsHealth)
		return
	}
	winner := teamsHealth
	fmt.Printf("winning team: %v\n", winner)
}
