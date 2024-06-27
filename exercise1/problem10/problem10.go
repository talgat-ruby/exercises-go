package problem10

func factory() (map[string]int, func (name string) func(increment int)) {
	brands := make(map[string]int, 0)

	makeBrand := func (name string) func(increment int)  {
		brands[name] = 0


		increment := func (increment int)  {
			brands[name] = brands[name] + increment
		}

		return increment
	}

	return brands, makeBrand
}
