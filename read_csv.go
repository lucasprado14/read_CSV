//ler csv, coluna "document", cada item é um cpf
//fazer uma chamada para http, para um desses document, para uma URL X a ser definida
//acontecer em paralelo, 100 go routines
//Ao termino do programa, printar o numero dos cpf que foram consultados com sucesso na URL

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

// CSV : csv
type CSV struct { //estrutura que receberá os dados do CSV
	CPF string
}

func checkErr(err error) { //checa erros
	if err != nil {
		log.Panic("ERROR: " + err.Error())
	}
}

func main() {

	// com arquivo
	csvFile, err := os.Open("test.csv") //abre arquivo
	checkErr(err)

	reader := csv.NewReader(bufio.NewReader(csvFile)) //lê arquivo
	reader.Comma = ';'                                //define delimitador

	var person []CSV

	for {
		line, err := reader.Read() //para cada linha
		if err == io.EOF {
			break
		} else if err != nil {
			checkErr(err)
		}
		person = append(person, CSV{ //adiciona uma pessoa
			CPF: line[0],
		})
	}

	fmt.Println(person[4].CPF) //exibe dados do csv

}
