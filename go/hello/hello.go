package main

type change struct {
	Banknote int
	Amount   int
}

var banknote = []int{1000, 500, 100, 50, 20, 10, 5, 2, 1}

func sell(price, pay int) []change {
	result := []change{}
	c := pay - price
	for _, note := range banknote {
		if c >= note {
			noteAmount := c / note
			c = c - (noteAmount * note)
			result = append(result, change{Banknote: note, Amount: noteAmount})
		}
	}
	return result
}
