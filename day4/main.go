package main

import (
	"fmt"
)

type num struct {
	n        int
	selected bool
}

type card struct {
	numbers [][]*num
	won     bool
}

func (c *card) selectnum(choice int) bool {
	for _, row := range c.numbers {
		for _, num := range row {
			if num.n == choice {
				num.selected = true
				return true
			}
		}
	}
	return false
}

func (c *card) check() bool {
	for _, row := range c.numbers {
		rowwin := true
		for _, num := range row {
			if !num.selected {
				rowwin = false
			}
		}
		if rowwin {
			c.won = true
			return c.won
		}
	}

	for i := 0; i < len(c.numbers[0]); i++ {
		colwin := true
		for j := 0; j < len(c.numbers); j++ {
			if !c.numbers[j][i].selected {
				colwin = false
			}
		}
		if colwin {
			c.won = true
			return c.won
		}
	}
	return false
}

func (c *card) score(win int) (score int) {
	for _, row := range c.numbers {
		for _, num := range row {
			if !num.selected {
				score += num.n
			}
		}
	}
	return score * win
}

type bingo struct {
	choices []int
	cards   []*card
}

func (b bingo) play() (*card, int) {
	for _, choice := range b.choices {
		for _, card := range b.cards {
			card.selectnum(choice)
			if card.check() {
				return card, choice
			}
		}
	}
	return b.cards[0], -1
}

func checkForLastOpenCard(cards []*card) (bool, *card) {
	var loosers []*card
	for _, card := range cards {
		if !card.won {
			loosers = append(loosers, card)
		}
	}
	if len(loosers) == 1 {
		return true, loosers[0]
	}
	return false, nil
}

func (b bingo) playToDeath() (*card, int) {
	for _, choice := range b.choices {
		for _, card := range b.cards {
			last, c := checkForLastOpenCard(b.cards)
			card.selectnum(choice)
			card.check()
			if last && c.won {
				return c, choice
			}
		}
	}
	return b.cards[0], -1
}

func task1(b *bingo) (result int) {
	winningcard, winningnum := b.play()
	if winningnum != -1 {
		result = winningcard.score(winningnum)
	}
	return result
}

func task2(b *bingo) (result int) {
	winningcard, winningnum := b.playToDeath()
	if winningnum != -1 {
		result = winningcard.score(winningnum)
	}
	return result
}

func main() {
	input := "input.txt"

	b := readdata(input)
	fmt.Println("Task 1 - # Score of winner =  ", task1(b))

	b = readdata(input)
	fmt.Println("Task 2 - # Score of looser =  ", task2(b))
}
