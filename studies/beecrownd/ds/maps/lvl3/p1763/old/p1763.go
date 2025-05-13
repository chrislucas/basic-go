// https://judge.beecrowd.com/en/problems/view/1763
// cases
// https://www.udebug.com/URI/1763
// DONE

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func Scanf(f string, a ...interface{}) (n int, e error) { return fmt.Fscanf(reader, f, a...) }

var ChristmasGreetings = map[string]string{
	"brasil":         "Feliz Natal!",
	"alemanha":       "Frohliche Weihnachten!",
	"austria":        "Frohe Weihnacht!",
	"coreia":         "Chuk Sung Tan!",
	"espanha":        "Feliz Navidad!",
	"grecia":         "Kala Christougena!",
	"estados-unidos": "Merry Christmas!",
	"inglaterra":     "Merry Christmas!",
	"australia":      "Merry Christmas!",
	"portugal":       "Feliz Natal!",
	"suecia":         "God Jul!",
	"turquia":        "Mutlu Noeller",
	"argentina":      "Feliz Navidad!",
	"chile":          "Feliz Navidad!",
	"mexico":         "Feliz Navidad!",
	"antardida":      "Merry Christmas!",
	"canada":         "Merry Christmas!",
	"irlanda":        "Nollaig Shona Dhuit!",
	"belgica":        "Zalig Kerstfeest!",
	"italia":         "Buon Natale!",
	"libia":          "Buon Natale!",
	"siria":          "Milad Mubarak!",
	"marrocos":       "Milad Mubarak!",
	"japao":          "Merii Kurisumasu!",
}

func RunUntilEndOFIle[T any](inv func(string) T) {
	for {
		var text string
		if _, err := Scanf("%s\n", &text); err != nil && err == io.EOF /*|| text == "" */ {
			break
		}
		text = strings.TrimSpace(text)
		inv(text)
	}
}

func ReaderValue[T any](tranform func(string) T) (T, error) {
	text, error := reader.ReadString('\n')
	if error != nil && error == io.EOF {
		var zero T
		return zero, error
	}
	text = strings.TrimSpace(text)
	return tranform(text), nil
}

func RunUntilCondition[T any](inv func(string) T) {
	for {
		text, error := ReaderValue(func(text string) string {
			return text
		})
		if error == io.EOF /*|| text == "" */ {
			break
		}
		//text = strings.TrimSpace(text)
		inv(text)
	}
}

func Solver() {
	RunUntilEndOFIle(func(text string) any {
		if greeting, found := ChristmasGreetings[text]; found {
			fmt.Println(greeting)
		} else {
			fmt.Println("--- NOT FOUND ---")
		}
		return nil
	})
}

func ClearString(text string) string {
	return strings.TrimSpace(text)
}

func SolverII() {
	answer := make([]string, 1)
	RunUntilCondition(func(text string) any {
		if greeting, found := ChristmasGreetings[text]; found {
			/*
				if len(answer) > 0 {
					answer = append(answer, fmt.Sprintf("\n%s", greeting))
				} else {
					answer = append(answer, greeting)
				}
			*/
			answer = append(answer, fmt.Sprintf("%s\n", greeting))
		} else {
			/*
				if len(answer) > 0 {
					answer = append(answer, fmt.Sprintf("\n%s", "--- NOT FOUND ---"))
				} else {
					answer = append(answer, "--- NOT FOUND ---\n")
				}
			*/
			answer = append(answer, "--- NOT FOUND ---\n")
		}
		return nil
	})

	fmt.Println(strings.Join(answer, ""))
}

func SolverIII() {
	RunUntilCondition(func(text string) any {
		if greeting, found := ChristmasGreetings[text]; found {
			fmt.Println(greeting)
		} else {
			fmt.Println("--- NOT FOUND ---")
		}
		return nil
	})
}

func RunUntilEOF(inv func(string)) {
	for {
		text, err := reader.ReadString('\n')
		if err != nil /*&& err == io.EOF*/ {
			break
		}
		inv(strings.TrimSpace(text))
	}
}

func SolverIV() {
	RunUntilEOF(func(text string) {
		text = strings.ToLower(text)
		if greeting, found := ChristmasGreetings[text]; found {
			fmt.Println(greeting)
		} else {
			fmt.Println("--- NOT FOUND ---")
		}
	})
}

func main() {
	SolverIV()
}

/*

fmt Print prints % end golang

fmt.Print:
It prints the given arguments to the console without adding any
 spaces between them and without appending a newline character at the end.
 Subsequent output will appear on the same line.

 fmt.Println:
It prints the given arguments to the console,
adding spaces between them if there are multiple arguments,
and always appends a newline character at the end, moving the cursor to the next line.


If you observe a % symbol at the end of your output when using fmt.Printf,
it's likely because you haven't added a newline character (\n) at the end of your format string,
and the shell prompt is appearing immediately after the output.
*/
