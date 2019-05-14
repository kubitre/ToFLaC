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
	ERROR        = iota + -2 // iota == 0 => -2
	NONTYPE      = iota + -1 // iota == 1 => 0
	INT          = iota + 0  // iota == 2 => 2
	FLOAT        = iota + 1  // iota == 3 => 4
	IDENTIFIER   = iota + 2  // iota == 4 => 6
	ENDSTATEMENT = iota + 3  // iota == 5 => 8
	POINTER      = iota + 4  // iota == 6 => 10
	SPACE        = iota + 5  // iota == 7 => 12
	NEWLINE      = iota + 6  // iota == 8 => 14
	COMMA        = iota + 7  // iota == 9 => 16
)
