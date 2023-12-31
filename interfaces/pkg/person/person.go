package person

type Person interface {
	GetAge() uint
}

func FindMaxAge(persones ...Person) uint {
	if len(persones) == 0 {
		return 0
	}

	max := persones[0].GetAge()
	for i := 0; i < len(persones); i++ {
		if persones[i].GetAge() > max {
			max = persones[i].GetAge()
		}
	}

	return max
}
