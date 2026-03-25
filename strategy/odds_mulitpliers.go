package strategy

func GetStdOddsMultipliers() map[int]int {
	return map[int]int{
		2:  1,
		3:  2,
		4:  3,
		5:  4,
		6:  5,
		8:  5,
		9:  4,
		10: 3,
		11: 2,
		12: 1,
	}
}

func Get100xMultipliers() map[int]int {
	return map[int]int{
		2:  100,
		3:  100,
		4:  100,
		5:  100,
		6:  100,
		8:  100,
		9:  100,
		10: 100,
		11: 100,
		12: 100,
	}
}

func Get2xMultipliers() map[int]int {
	return map[int]int{
		2:  2,
		3:  2,
		4:  2,
		5:  2,
		6:  2,
		8:  2,
		9:  2,
		10: 2,
		11: 2,
		12: 2,
	}
}
