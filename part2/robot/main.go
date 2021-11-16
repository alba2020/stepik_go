package main

type Robot struct {
	On bool
	Ammo int
	Power int
}

func (r *Robot) Shoot() bool {
	if !r.On || r.Ammo < 1 {
		return false
	} else {
		r.Ammo--
		return true
	}
}

func (r *Robot) RideBike() bool {
	if ! r.On || r.Power < 1 {
		return false
	} else {
		r.Power--
		return true
	}
}

func main() {
	testStruct := new(Robot)
	_ = testStruct
}