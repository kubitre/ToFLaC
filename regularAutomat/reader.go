package regularAutomat

import (
	"fmt"
	// "bufio"
	// "os"
)

type RegularExp struct {
	automat *AutomatState
}

/*Reader - последовательный считыватель*/
func (regexp *RegularExp) Reader() {
	// reader := bufio.NewReader(os.Stdin)
	// for {
	// 	fmt.Println("For exit write the @q")
	// input, _ := reader.ReadString('\n')
	// input = string([]byte(input))
	// fmt.Println("reading: ", input)
	// if input == "@q"{
	// 	break
	// }
	input := "192.12.3.2 test new ip 93.43.34.43"
	regexp.parser(input)
	// }

	fmt.Println("Finding ips: ")
	for _, value := range regexp.automat.AllMatches {
		fmt.Println(value)
	}
}

func (reg *RegularExp) parser(str string) {
	for _, value := range str {
		fmt.Println("char: " + string(value))
		reg.automat.NewSymbs(string(value))
	}
	reg.automat.NewSymbs(" ")
	fmt.Println(reg.automat.GetLog())
	reg.automat.ClearLogs()
}

/*New - создание нового автомата для регулярного парсинга*/
func (regexp *RegularExp) New() *RegularExp {
	reg := &RegularExp{
		automat: CreateAutomat(),
	}

	return reg
}
