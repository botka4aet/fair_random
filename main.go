package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Количество граней кубика
var dice = 10

// Количество бросков, которые должны полностью закрыть грани кубика, перед промежуточным отчетом
var step = 100

// Количество серий бросков перед выводом результата
var steplim = 100

func init() {
}

func main() {
	//Отклонение от нормы в процентах.
	//0 - полное совпадение, 100 - одна из граней выпала в два раза больше
	var percentage float32
	//Отклонение от нормы в бросках
	var delta int

	//Стандартный рандом
	normalrandom := make(map[int]int)
	for i := 1; i <= steplim; i++ {
		for j := 0; j < step*dice; j++ {
			normalrandom[rand.Intn(dice)]++
		}
		for _, v := range normalrandom {
			tdelta := AbsInt(v - step*i)
			if tdelta > delta {
				delta = tdelta
			}
		}
	}
	percentage = float32(delta*100) / float32(step*steplim)
	fmt.Printf("[Обыч]Max отклонение - %.2f%%\n", percentage)
	//Стандартный рандом конец

	//Рандом со срезом
	//Сколько в срезе должно быть элементов перед попыткой его дозаполнить
	sllimit := 1000

	slicerandom := make(map[int]int)
	slcounter := []int{}
	delta = 0
	for i := 1; i <= steplim; i++ {
		for j := 0; j < step*dice; j++ {
			for len(slcounter) < sllimit {
				for n := 0; n < dice; n++ {
					slcounter = append(slcounter, n)
				}
			}
			slposition := rand.Intn(len(slcounter))
			slicerandom[slcounter[slposition]]++
			slcounter = remove(slcounter, slposition)
		}
		for _, v := range slicerandom {
			tdelta := AbsInt(v - step*i)
			if tdelta > delta {
				delta = tdelta
			}
		}
	}
	percentage = float32(delta*100) / float32(step*steplim)
	fmt.Printf("[Срез]Max отклонение - %.2f%%\n", percentage)
	//Рандом со срезом конец
}

func AbsInt(delta int) int {
	if delta < 0 {
		delta = -delta
	}
	return delta
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

var Rng *rand.Rand

func init() {
	Rng = rand.New(rand.NewSource(time.Now().UnixNano()))

}
