// http://judge.beecrowd.com/en/problems/view/1430
// https://judge.beecrowd.com/en/problems/view/1430
// DONE
package main

import (
	"bufio"
	"fmt"

	//"io"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func Scanf(f string, a ...interface{}) (n int, e error) { return fmt.Fscanf(reader, f, a...) }

/*
func RunUntilEndOFIle[T any](inv func(string) T) {
	for {
		var text string
		if _, err := Scanf("%s\n", &text); err != nil && err == io.EOF {
			break
		}

		inv(text)
	}
}
*/

func filter(arr []string, condition func(string) bool) []string {
	result := []string{}
	for _, str := range arr {
		if condition(str) {
			result = append(result, str)
		}
	}
	return result
}

func main() {

	/*
		/HH/QQQQ/XXXTXTEQH/W/HW/
		/W/W/SQHES/
		/WE/TEX/THES/
		*
	*/

	noteDurations := map[string]float64{
		"W": 1.0,
		"H": 0.5,
		"Q": 0.25,
		"E": 0.125,
		"S": 0.0625,
		"T": 0.03125,
		"X": 0.015625,
	}

	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if text == "*" {
			break
		}
		parts := strings.Split(text, "/")
		parts = filter(parts, func(part string) bool {
			return part != ""
		})

		var answer int
		for _, part := range parts {
			note := strings.TrimSpace(part)
			duration := 0.0
			for _, char := range note {
				duration += noteDurations[string(char)]
			}
			if duration == 1.0 {
				answer++
			}
		}
		fmt.Println(answer)
		/*
			durationCount := make(map[float64]int)
			for _, part := range parts {
				note := strings.TrimSpace(part)
				duration := 0.0
				for _, char := range note {
					duration += noteDurations[string(char)]
				}
				durationCount[duration]++
			}
			var answer int
			for k, v := range durationCount {
				if k == 1.0 {
					answer = v
					break
				}
			}
			fmt.Println(answer)
		*/
	}
}

/*
	func BuilderString[T any](strs []T, sep string) string {
		var sb strings.Builder
		for _, value := range strs {
			sb.WriteString(fmt.Sprintf("%v%s", value, sep))
		}
		return sb.String()
	}

	func Stringify[T any](value T) string {
		return fmt.Sprintf("%v", value)
	}


	func ReadLine() string {
		line, _ := reader.ReadString('\n')
		return strings.TrimSpace(line)
	}

	func ReadInt() int {
		n, _ := strconv.Atoi(ReadLine())
		return n
	}
*/
