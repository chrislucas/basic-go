// https://medium.com/openskill/golang-demistifying-interface-value-and-pointer-receiver-365de07c2293

package main

import "fmt"

/*
	Uma inrerface tem 2 campos
		- Um pointeiro pra uma informacao de um tipo especifico.
		podemos pensar nisso como um type
		- Data Pointer: Se um dado armazenado é um ponteiro,
		ele é armazenado diretamente, se for um valor, entao
		um ponteiro para o valor é armazenado
*/

type error interface {
	Error() string
}

func New(test string) error {
	return &ErrorMessage{text: test}
}

type ErrorMessage struct {
	text string
}

func (e *ErrorMessage) Error() string {
	return e.text
}

type HttpError struct {
	statusCode int
	message    string
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("HTTP Error: %d - %s", e.statusCode, e.message)
}

func NewHttpError(statusCode int, message string) *HttpError {
	return &HttpError{statusCode: statusCode, message: message}
}

func main() {
	var customError = New("test")

	fmt.Println(customError)
	fmt.Println(customError.Error())
}
