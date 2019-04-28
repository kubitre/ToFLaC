package token

/*Type 			- тип токена
ERROR 			- ошибочное распознание
NONTYPE 		- не токен
INT 			- тип целочисленный
FLOAT 			- тип с плавающей точкой
IDENTIFIER 		- идентификатор
ENDSTATEMENT 	- окончания объявления переменной
POINTER 		- указатель
SPACE 			- пробел
NEWLINE 		- новая строка
COMMA 			- запятая для перечисления идентификаторов
*/
const (
	ERROR        = iota + -2
	NONTYPE      = iota + -1
	INT          = iota + 0
	FLOAT        = iota + 1
	IDENTIFIER   = iota + 2
	ENDSTATEMENT = iota + 3
	POINTER      = iota + 4
	SPACE        = iota + 5
	NEWLINE      = iota + 6
	COMMA        = iota + 7
)
