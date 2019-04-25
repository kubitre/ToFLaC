package Analysis

import "regexp"

/*Analysator - структура результат поиска выражения в строке*/
type Analysator struct {
	Token   string
	Atribut string
}

/*GetAnalys - получение анализа*/
func (anal *Analysator) GetAnalys(value string) map[string]string {
	resultregexpIF := regexp.MustCompile(GetToken("if"))
	return map[string]string{
		"result": resultregexpIF.ReplaceAllString(value, GetResult("if")),
	}
}

/*ConfigureRegexp - конфигурация паттернов и раскрасок*/
func (anal *Analysator) ConfigureRegexp() {

}
