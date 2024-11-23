package data

import (
	"fe3h-randomizer/common"
	"fmt"
	"math/rand"
)

type weight struct {
	LevelOneOdds   int
	LevelTwoOdds   int
	LevelThreeOdds int
	LevelFourOdds  int
	LevelFiveOdds  int
	LevelSixOdds   int
}

func getWeights(difficultyLevel uint) weight {
	var weighted weight = weight{}
	if difficultyLevel == 1 {
		weighted.LevelOneOdds = 80
		weighted.LevelTwoOdds = 15
		weighted.LevelThreeOdds = 4
		weighted.LevelFourOdds = 1
	} else if difficultyLevel == 2 {
		weighted.LevelOneOdds = 20
		weighted.LevelTwoOdds = 50
		weighted.LevelThreeOdds = 15
		weighted.LevelFourOdds = 5
	} else if difficultyLevel == 3 {
		weighted.LevelOneOdds = 10
		weighted.LevelTwoOdds = 20
		weighted.LevelThreeOdds = 55
		weighted.LevelFourOdds = 15
	} else if difficultyLevel == 4 {
		weighted.LevelOneOdds = 1
		weighted.LevelTwoOdds = 10
		weighted.LevelThreeOdds = 20
		weighted.LevelFourOdds = 50
		weighted.LevelFiveOdds = 19
	} else if difficultyLevel == 5 {
		weighted.LevelThreeOdds = 15
		weighted.LevelFourOdds = 25
		weighted.LevelFiveOdds = 40
		weighted.LevelSixOdds = 20
	} else {
		weighted.LevelFourOdds = 5
		weighted.LevelFiveOdds = 45
		weighted.LevelSixOdds = 55
	}

	return weighted
}

func GenerateWeightedCharacter(difficulty uint) error {
	var weights = getWeights(difficulty)
	var classRoll = rand.Intn(101)
	var classRank = getRankFromWeight(weights, classRoll)
	var class = common.PickNItems(SkillLevels[classRank].Classes, 1)[0]
	fmt.Println("Class:", class)
	var weaponsRoll = make([]string, difficulty)
	for i := range weaponsRoll {
		var slotRoll = getRankFromWeight(weights, rand.Intn(101))
		var weapon = common.PickNItems(SkillLevels[slotRoll].Weapons, 1)
		weaponsRoll[i] = weapon[0]
	}

	fmt.Println(weaponsRoll)

	return nil
}

func getRankFromWeight(weighted weight, roll int) uint {
	var i int = 0

	i += weighted.LevelOneOdds
	if i >= roll {
		return 1
	}

	i += weighted.LevelTwoOdds
	if i >= roll {
		return 2
	}

	i += weighted.LevelThreeOdds
	if i >= roll {
		return 3
	}

	i += weighted.LevelFourOdds
	if i >= roll {
		return 4
	}

	i += weighted.LevelFiveOdds
	if i >= roll {
		return 5
	}

	return 6
}
