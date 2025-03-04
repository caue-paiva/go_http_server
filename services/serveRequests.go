package services

import (
	"fmt"
	"os"
)

// retorna um array de strings para cada linha do texto da request e o index da quebra de linha que separa os headers do body
func GetRequestContents(filePath string) ([]string, int) {
	bytes, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println(err.Error())
	}

	lines := make([]string, 0, 20)
	startByte := 0
	lineCounter := 0
	var lineBreak int = -1 //guarda o index do linebrak

	for i, by := range bytes {
		if by == byte('\n') {
			if (i - startByte) <= 1 { //achamos o linebreak
				lineBreak = lineCounter
			} else {
				lines = append(lines, string(bytes[startByte:i]))
			}

			lineCounter++
			startByte = i + 1
		}
	}

	if startByte < len(bytes) {
		lines = append(lines, string(bytes[startByte:])) //append da linha final
	}
	return lines, lineBreak
}

// transforma o slice de bytes da request em um slice de strings, cada entrada representando uma linha
func DeserializeRequest(requestBytes []byte) ([]string, int) {
	lines := make([]string, 0, 20)
	startByte := 0
	lineCounter := 0
	var lineBreak int = -1 //guarda o index do linebrak

	for i := 0; i < len(requestBytes)-1; i++ {

		if requestBytes[i] == byte('\r') && requestBytes[i+1] == byte('\n') { //final da linha
			line := string(requestBytes[startByte:i])
			lines = append(lines, line)
			if len(line) == 0 && lineBreak == -1 {
				lineBreak = lineCounter
			}

			lineCounter++
			startByte = i + 2 //avança byte do começo da prox linha, pulando o \n também
			i++               //pula o \n
		}

	}

	if startByte < len(requestBytes) {
		lines = append(lines, string(requestBytes[startByte:])) //append da linha final
	}
	return lines, lineBreak
}
