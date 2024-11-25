package main

import "fmt"

// Есть неотсортированные билеты с существующим маршрутом из точки А в B.
// требуется отсортировать по порядку следования за O(N)

type ticket struct {
	from string
	to   string
}

func (t ticket) String() string {
	return fmt.Sprintf("%s->%s", t.from, t.to)

}

func main() {
	tickets := []ticket{
		{"SJC", "KKA"},
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
		{"KKA", "KUU"},
		{"LHR", "SFO"},
	}

	printTickets(tickets)

	err := sortTickets(&tickets)
	if err != nil {
		fmt.Printf("Sorting error %V", err)
	}

	printTickets(tickets)

}

func sortTickets(tickets *[]ticket) error {
	// ищем первый билет.
	fromMap := make(map[string]int)
	toMap := make(map[string]struct{})

	for i, t := range *tickets {
		fromMap[t.from] = i
		toMap[t.to] = struct{}{}
	}

	var firstIndex = -1
	for key, idx := range fromMap {
		_, exist := toMap[key]
		if !exist {
			firstIndex = idx
			break
		}
	}
	if firstIndex == -1 {
		return fmt.Errorf("first ticket not found")
	}

	// ставим этот билет на первое место
	(*tickets)[0], (*tickets)[firstIndex] = (*tickets)[firstIndex], (*tickets)[0]
	fromMap[(*tickets)[0].from] = 0
	fromMap[(*tickets)[firstIndex].from] = firstIndex
	nextFrom := (*tickets)[0].to

	// перебираем билеты со второго и ставим на свое место
	for i := 1; i < len(*tickets); i++ {
		nextIndex, ok := fromMap[nextFrom]
		if !ok {
			break
		}

		(*tickets)[nextIndex], (*tickets)[i] = (*tickets)[i], (*tickets)[nextIndex]
		fromMap[(*tickets)[i].from] = i
		fromMap[(*tickets)[nextIndex].from] = nextIndex
		nextFrom = (*tickets)[i].to
	}

	return nil
}

func printTickets(tickets []ticket) {
	for i, t := range tickets {
		fmt.Printf("%d:%s ", i, t)
	}
	fmt.Println("")
}
