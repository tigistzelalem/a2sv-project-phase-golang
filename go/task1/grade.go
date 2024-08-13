package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// calculate average grade
func calculateAverage(grades []float64) float64 {
	var total float64
	for _, grade := range grades {
		total += grade
	}

	return total / float64(len(grades))
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("Enter your number of subjects: ")
	numSubjectsStr, _ := reader.ReadString('\n')
	numSubjectsStr = strings.TrimSpace(numSubjectsStr)
	numSubjects, err := strconv.Atoi(numSubjectsStr)
	if err != nil || numSubjects <= 0 {
		log.Fatal("The number of subjects must be a valid integer greater than zero")
	}

	subjects := make([]string, numSubjects)
	grades := make([]float64, numSubjects)

	for i := 0; i < numSubjects; i++ {
		fmt.Printf("Enter the name of the subject %d: ", i+1)
		subjects[i], _ = reader.ReadString('\n')
		subjects[i] = strings.TrimSpace(subjects[i])

		fmt.Printf("Enter the grade for %s: ", subjects[i])
		gradeStr, _ := reader.ReadString('\n')
		gradeStr = strings.TrimSpace(gradeStr)
		grade, err := strconv.ParseFloat(gradeStr, 64)
		if err != nil || grade < 0 || grade > 100 {
			log.Fatalf("Grade for %s must be a valid number between 0 and 100.", subjects[i])
		}
		grades[i] = grade
	}

	averageGrade := calculateAverage(grades)
	fmt.Printf("\nStudent Name: %s\n", name)
	fmt.Println("Grades:")
	for i := 0; i < numSubjects; i++ {
		fmt.Printf("%s: %.2f\n", subjects[i], grades[i])
	}
	fmt.Printf("Average Grade: %.2f\n", averageGrade)
}
