package utils

type Random struct {
	seed uint
}

func (r *Random) Seed(seed uint) {
	seed = seed % 2147483647
	if seed <= 0 {
		seed += 2147483646
	}

	r.seed = seed
}

func (r *Random) Int() int {
	r.seed = r.seed * 16807 % 2147483647

	return int(r.seed)

}

func (r *Random) Intn(n int) int {
	return r.Int() % n
}

func NewRandom(seed uint) *Random {
	r := &Random{}
	r.Seed(seed)

	return r
}
