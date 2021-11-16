package main

func main() {
	groupCity := map[int][]string{
		10:   []string{"aaa", "bbb", "ccc"}, // города с населением 10-99 тыс. человек
		100:  []string{"ddd", "eee"}, // города с населением 100-999 тыс. человек
		1000: []string{"fff", "ggg"}, // города с населением 1000 тыс. человек и более
	}
	
	cityPopulation := map[string]int {
		"aaa": 10, "bbb": 10, "ddd": 100, "eee": 100, "fff": 1000 }
	
	outer:
	for name, _ := range cityPopulation {
		for _, name2 := range groupCity[100] {
			if name == name2 {
				continue outer
			}
		}
		delete(cityPopulation, name)
	}
}
