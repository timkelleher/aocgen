package aoc

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

func Args() (int, int) {
	year, _, _ := time.Now().Date()
	day := 0

	args := os.Args[1:]
	if len(args) > 0 {
		year, _ = strconv.Atoi(args[0])
	}
	if len(args) > 1 {
		day, _ = strconv.Atoi(args[1])
	}

	return year, day
}

func Year() int {
	year, _, _ := time.Now().Date()
	return year
}

func Input(year, day int) []string {
	fileName := fmt.Sprintf("./pkg/year%d/inputs/%d.txt", year, day)
	return readFile(fileName)
}

func TestInput(year, day int) []string {
	fileName := fmt.Sprintf("../../pkg/year%d/inputs/%d.txt", year, day)
	return readFile(fileName)
}

func readFile(fileName string) []string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		logrus.Fatal(err)
	}

	lines := strings.Split(string(b), "\n")
	return lines
}
