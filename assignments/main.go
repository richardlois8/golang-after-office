package main

import (
	"fmt"
	"sort"
	"bytes"
	"math/rand"
)

func main() {
	// fmt.Println("hello world")

	fmt.Println("Array Sign")
	// arraySign([]int{2, 1})                    // 1
	// arraySign([]int{-2, 1})                   // -1
	// arraySign([]int{-1, -2, -3, -4, 3, 2, 1}) // 1
	fmt.Println(arraySign([]int{2, 1}))                    // 1
	fmt.Println(arraySign([]int{-2, 1}))                   // -1
	fmt.Println(arraySign([]int{-1, -2, -3, -4, 3, 2, 1})) // 1

	fmt.Println("Is Anagram")
	// isAnagram("anak", "kana") // true
	// isAnagram("anak", "mana") // false
	// isAnagram("anagram", "managra") // true
	fmt.Println(isAnagram("anak", "kana")) // true
	fmt.Println(isAnagram("anak", "mana")) // false
	fmt.Println(isAnagram("anagram", "managra")) // true

	fmt.Println("Find The Difference")
	// findTheDifference("abcd", "abcde") // 'e'
	// findTheDifference("abcd", "abced") // 'e'
	// findTheDifference("", "y")         // 'y'
	fmt.Println(string(findTheDifference("abcd", "abcde"))) // 'e'
	fmt.Println(string(findTheDifference("abcd", "abced"))) // 'e'
	fmt.Println(string(findTheDifference("", "y")))         // 'y'

	fmt.Println("Can Make Arithmetic Progression From Sequence")
	// canMakeArithmeticProgression([]int{1, 5, 3})    // true; 1, 3, 5 adalah baris aritmatik +2
	// canMakeArithmeticProgression([]int{5, 1, 9})    // true; 9, 5, 1 adalah baris aritmatik -4
	// canMakeArithmeticProgression([]int{1, 2, 4, 8}) // false; 1, 2, 4, 8 bukan baris aritmatik, melainkan geometrik x2
	fmt.Println(canMakeArithmeticProgression([]int{1, 5, 3}))    // true; 1, 3, 5 adalah baris aritmatik +2
	fmt.Println(canMakeArithmeticProgression([]int{5, 1, 9}))    // true; 9, 5, 1 adalah baris aritmatik -4
	fmt.Println(canMakeArithmeticProgression([]int{1, 2, 4, 8}))
	
	tesDeck()
}

// https://leetcode.com/problems/sign-of-the-product-of-an-array
func arraySign(nums []int) int {
	// write code here
	product := 1
	for _, value := range nums {
		product *= value
	}

	if product > 0 {
		return 1
	} else {
		return -1
	}
}

// https://leetcode.com/problems/valid-anagram
func isAnagram(s string, t string) bool {
	// write code here
	bytesS := []byte(s)
	bytesT := []byte(t)
	if len(bytesS) != len(bytesT) {
		return false
	}
	if compareSortedBytes(bytesS, bytesT) {
		return true
	}
	return false	
}

func sortBytes(b []byte) []byte {
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	return b
}

func compareSortedBytes(bytesS, bytesT []byte) bool {
	return bytes.Equal(sortBytes(bytesS), sortBytes(bytesT))
}

// https://leetcode.com/problems/find-the-difference
func findTheDifference(s string, t string) byte {
	// write code here
	var byteS, byteT byte
	for i := 0; i < len(s); i++ {
		byteS += s[i]
	}
	for i := 0; i < len(t); i++ {
		byteT += t[i]
	}
	b := byteT - byteS
	return b
}

// https://leetcode.com/problems/can-make-arithmetic-progression-from-sequence
func canMakeArithmeticProgression(arr []int) bool {
	// write code here
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 0; i < len(arr) - 1; i++ {
		if arr[i+1] - arr[i] != diff {
			return false
		}
	}
	return true
}

// Deck represent "standard" deck consist of 52 cards
type Deck struct {
	cards []Card
}

// Card represent a card in "standard" deck
type Card struct {
	symbol int // 0: spade, 1: heart, 2: club, 3: diamond
	number int // Ace: 1, Jack: 11, Queen: 12, King: 13
}

// New insert 52 cards into deck d, sorted by symbol & then number.
// [A Spade, 2 Spade,  ..., A Heart, 2 Heart, ..., J Diamond, Q Diamond, K Diamond ]
// assume Ace-Spade on top of deck.
func (d *Deck) New() {
	// write code here
	d.cards = make([]Card, 52)
	index := 0
	for symbol := 0; symbol < 4; symbol++ {
		for number := 1; number <= 13; number++ {
			d.cards[index] = Card{symbol: symbol, number: number}
			index++
		}
	}
}

// PeekTop return n cards from the top
func (d Deck) PeekTop(n int) []Card {
	// write code here
	return d.cards[:n]
}

// PeekTop return n cards from the bottom
func (d Deck) PeekBottom(n int) []Card {
	// write code here
	return d.cards[len(d.cards)-n:]
}

// PeekCardAtIndex return a card at specified index
func (d Deck) PeekCardAtIndex(idx int) Card {
	return d.cards[idx]
}

// Shuffle randomly shuffle the deck
func (d *Deck) Shuffle() {
	// write code here
	for i := range d.cards {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// Cut perform single "Cut" technique. Move n top cards to bottom
// e.g. Deck: [1, 2, 3, 4, 5]. Cut(3) resulting Deck: [4, 5, 1, 2, 3]
func (d *Deck) Cut(n int) {
	// write code here
	d.cards = append(d.cards[n:], d.cards[:n]...)
}

func (c Card) ToString() string {
	textNum := ""
	switch c.number {
	case 1:
		textNum = "Ace"
	case 11:
		textNum = "Jack"
	case 12:
		textNum = "Queen"
	case 13:
		textNum = "King"
	default:
		textNum = fmt.Sprintf("%d", c.number)
	}
	texts := []string{"Spade", "Heart", "Club", "Diamond"}
	return fmt.Sprintf("%s %s", textNum, texts[c.symbol])
}

func tesDeck() {
	deck := Deck{}
	deck.New()
	top5Cards := deck.PeekTop(3)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}
	fmt.Println("---\n")

	fmt.Println(deck.PeekCardAtIndex(12).ToString()) // Queen Spade
	fmt.Println(deck.PeekCardAtIndex(13).ToString()) // King Spade
	fmt.Println(deck.PeekCardAtIndex(14).ToString()) // Ace Heart
	fmt.Println(deck.PeekCardAtIndex(15).ToString()) // 2 Heart
	fmt.Println("---\n")

	deck.Shuffle()
	top5Cards = deck.PeekTop(10)
	for _, c := range top5Cards {
		fmt.Println(c.ToString())
	}

	fmt.Println("---\n")
	deck.New()
	deck.Cut(5)
	bottomCards := deck.PeekBottom(10)
	for _, c := range bottomCards {
		fmt.Println(c.ToString())
	}
}
