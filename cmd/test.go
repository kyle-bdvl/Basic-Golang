package main

import (
	"fmt"
)

func Exercise1() {
	score := 80

	fmt.Println(score)
	fmt.Println(&score)
}

func Exercise2() {
	score := 1000
	p := &score
	fmt.Println("Address : ", p)
	fmt.Println("*p : ", *p)

}

func Exercise3() {
	score := 80
	p := &score
	*p = 95
	fmt.Println(*p)
}

func addBonus(score *int) {
	*score += 10
}
func Exercise4() {
	score := 80
	addBonus(&score)
	fmt.Println(score)
}

func swap(x, y *float64) {
	*y, *x = *x, *y
}
func Exercise5() {
	x, y := 4.0, 8.0
	swap(&x, &y)
	fmt.Println(x)
}

type Student struct {
	Name  string
	Score int
}

func increaseScore(s *Student) {
	s.Score += 20
}

func Exercise6() {
	s := Student{
		Name:  "Alice",
		Score: 80,
	}

	increaseScore(&s)
	fmt.Printf("Scores : %v\n", s.Score)
}

func LevelUp(s *Student) {
	s.Name = "Pro " + s.Name
	s.Score += 50
}

func (s *Student) LevelUp() {
	s.Name = "Legendary " + s.Name
	s.Score += 100
}

func Exercise7() {
	r := Student{
		Name:  "Bob",
		Score: 70,
	}
	fmt.Printf("%p\n", &r)
	// LevelUp(&r)
	r.LevelUp()
	LevelUp(&r)
	fmt.Printf("Name : %v\n", r.Name)

	fmt.Printf("Scores : %v\n", r.Score)

}

// func main() {
// 	Exercise1()
// 	fmt.Println("========================================")
// 	Exercise2()
// 	fmt.Println("========================================")
// 	Exercise3()
// 	fmt.Println("========================================")
// 	Exercise4()
// 	fmt.Println("========================================")
// 	Exercise5()
// 	fmt.Println("========================================")
// 	Exercise6()
// 	fmt.Println("========================================")
// 	Exercise7()
// }
