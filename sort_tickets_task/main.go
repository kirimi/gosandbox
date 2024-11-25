package main

import "fmt"

type ticket struct {
	from string
	to   string
}

func (t ticket) String() string {
	return fmt.Sprintf("%s->%s", t.from, t.to)

}

func main() {
	tickets := []ticket{
		{"MUC", "LHR"},
		{"JFK", "MUC"},
		{"SFO", "SJC"},
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
	toMap := make(map[string]int)

	for i, t := range *tickets {
		fromMap[t.from] = i
		toMap[t.to] = i
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
	nextFrom := (*tickets)[0].to

	// перебираем билеты со второго и ставим на свое место
	for i := 1; i < len(*tickets); i++ {
		rightTickets := (*tickets)[i:]
		nextIndex := getTicketIndexWithFrom(&rightTickets, nextFrom)
		if nextIndex == -1 {
			return nil
		}

		(*tickets)[nextIndex+i], (*tickets)[i] = (*tickets)[i], (*tickets)[nextIndex+i]
		nextFrom = (*tickets)[i].to
	}

	return nil
}

func getTicketIndexWithFrom(tickets *[]ticket, from string) int {
	for i, t := range *tickets {
		if t.from == from {
			return i
		}
	}
	return -1
}

func printTickets(tickets []ticket) {
	for i, t := range tickets {
		fmt.Printf("%d:%s ", i, t)
	}
	fmt.Println("")
}
