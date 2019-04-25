package Analysis

/*Tokens - таблица токенов для распарса*/
var Tokens map[string]string

/*Result - таблица соответсвий блоков и токенов*/
var Result map[string]string

/*ColorScheme - схема цветов для подсветки*/
var ColorScheme map[string]string

/*InitColors - инициализация цветов*/
func InitColors() map[string]string {
	ColorScheme := map[string]string{
		"base03":  "#002b36",
		"base02":  "#073642",
		"base01":  "#586e75",
		"base00":  "#657b83",
		"base0":   "#839496",
		"base1":   "#93a1a1",
		"base2":   "#eee8d5",
		"base3":   "#fdf6e3",
		"yellow":  "#b58900",
		"orange":  "#cb4b16",
		"red":     "#dc322f",
		"magenta": "#d33682",
		"violet":  "#6c71c4",
		"blue":    "#268bd2",
		"cyan":    "#2aa198",
		"green":   "#859900",
	}
	return ColorScheme
}

/*InitTokens - инициализация токенов*/
func InitTokens() {
	Tokens["if"] = "if"
	Tokens["then"] = "then"
	Tokens["else"] = "else"
	Tokens["id"] = "(id [a-zA-Z_][a-zA-Z0-9]*)"
	Tokens["num"] = "(num \\d*)"
	Tokens["leftArrow"] = "\\d{1, 10}[ ]{0, }<[ ]{0, }\\d{1, 10}"
	Tokens["leftArrowEquals"] = "\\d{1, 10}[ ]{0, }<=[ ]{0, }\\d{1, 10}"
	Tokens["equals"] = "\\d{1, 10}[ ]{0, }=[ ]{0, }\\d{1, 10}"
	Tokens["notequals"] = "\\d{1, 10}[ ]{0, }<>[ ]{0, }\\d{1, 10}"

	Result["if"] = `<span style={{color: "` + ColorScheme["blue"] + `"}}>if</span>`
	Result["then"] = `<span style={{color: "` + ColorScheme["blue"] + `"}}>then</span>`
	Result["else"] = `<span style={{color: "` + ColorScheme["blue"] + `"}}>else</span>`
	Result["id"] = `<span style={{color: "` + ColorScheme["red"] + `"}}>id</span>`
	Result["num"] = `<span style={{color: "` + ColorScheme["red"] + `"}}>num</span>`
	Result["leftArrow"] = `<span style={{color: "` + ColorScheme["base01"] + `"}}>&lt;</span>`
	Result["leftArrowEquals"] = `<span style={{color: "` + ColorScheme["base01"] + `"}}>&lt;=</span>`
	Result["equals"] = `<span style={{color: "` + ColorScheme["base01"] + `"}}>=</span>`
	Result["notequals"] = `<span style={{color: "` + ColorScheme["base01"] + `"}}>&lt;&gt;</span>`

}

/*GetToken - получение токена по ключу*/
func GetToken(key string) (value string) {
	return Tokens[key]
}

/*GetResult - получение раскраски по ключу*/
func GetResult(key string) (value string) {
	return Result[key]
}
