package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

/*func (s Student) getAge() int {
	return s.age
}*/

/*func (s *Student) setAge(age int) {
	s.age = age
}*/

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, val := range s.grades {
		sum += val
	}
	return float32(sum) / float32(len(s.grades))
}

func main() {
	s1 := Student{"Tim", []int{80, 85, 90, 95}, 18}
	s2 := Student{"Yash", []int{80, 80, 90, 91, 99}, 21}
	//s1.getAge()
	average := s1.getAverageGrade()
	average2 := s2.getAverageGrade()
	fmt.Println(average, "\n", average2)
}
