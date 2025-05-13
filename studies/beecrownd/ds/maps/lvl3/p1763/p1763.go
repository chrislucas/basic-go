// https://judge.beecrowd.com/en/problems/view/1763
// cases
// https://www.udebug.com/URI/1763
// DONE

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

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

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

func RunUntilEOF(inv func(string)) {
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			break
		}
		inv(strings.TrimSpace(text))
	}
}

func AnotherRunUntilEOF(inv func(string)) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		inv(strings.TrimSpace(text))
	}
}

func Solver() {
	AnotherRunUntilEOF(func(text string) {
		text = strings.ToLower(text)
		if greeting, found := ChristmasGreetings[text]; found {
			fmt.Println(greeting)
		} else {
			fmt.Println("--- NOT FOUND ---")
		}
	})
}

func main() {
	Solver()
}

/*
	A, B
	n = Inteiro
	O(n), O(n log n)

	array 10_000_000
	Busca Normal = O(n)
	for	i ate n
		if array[i] == X
			return i
		end
	end

	A => O (n) = 10_000_000
	B => O (log n) = 7
	n = 10_000_000
	A > B

	Busca Binaria = O(log n)
	[Menu(id=1), Menu(id=2), Menu(id=3), Menu(id=4), Menu(id=5), Menu(id=6), Menu(id=7)]

	busca_binaria(id int)
	id = 2

	esquerda = 0
	direita = len(array) - 1
	meio = (esquerda + direita) / 2

	2 > 4
	[Menu(id=1), Menu(id=2), Menu(id=3)]
	direita = meio - 1
	meio = (esquerda + direita) / 2

	1000
	500
	250
	125

	List.find(compara Menu == Menu(1)) => O(n)
	List.binary_search(compara Menu == Menu(1)) => O(log n)

	F -> O(n log n) + O(log n)
	n = 10000
	132000

	G -> O(n) = 10000

	F > G

*/
