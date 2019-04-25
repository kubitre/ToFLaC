package regularAutomat

import (
	"fmt"
)

type RegularExp struct {
	automat *AutomatState
}

/*Reader - последовательный считыватель*/
func (regexp *RegularExp) Reader() {
	// reader := bufio.NewReader(os.Stdin)
	for {
		// input, _ := reader.ReadString('\n')
		// input = string([]byte(input))
		input := "123.23.21.41"
		regexp.parser(input)
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
