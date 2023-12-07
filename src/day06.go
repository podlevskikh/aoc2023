package src

func (s *Service) CalculateDay06FirstPart() int {
	races := [][]int{{35, 213}, {69, 1168}, {68, 1086}, {87, 1248}}

	res := 1
	for _, race := range races {
		time := race[0]
		distance := race[1]

		winNum := 0
		for i := 1; i < time; i++ {
			if (time-i)*i > distance {
				winNum++
			}
		}

		res *= winNum
	}

	return res
}

func (s *Service) CalculateDay06SecondPart() int {
	races := [][]int{{35696887, 213116810861248}}

	res := 1
	for _, race := range races {
		time := race[0]
		distance := race[1]

		winNum := 0
		for i := 1; i < time; i++ {
			if (time-i)*i > distance {
				winNum++
			}
		}

		res *= winNum
	}

	return res
}
