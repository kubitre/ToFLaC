package Middleware

import (
	"time"

	"github.com/fatih/color"
)

/*Logger - слой логирования для запросов*/
type Logger struct {
	APIRoute      string        // маршрут запроса
	RequestMethod string        // метод запроса
	ContentType   string        // тип контента
	RequestBody   string        // содержимое запроса
	Time          time.Duration // время ответа на запрос
	RequestStatus bool          //произошла ли ошибка во время зарпоса
	OutputMethod  *color.Color
}

/*ConfigTrace - конфигуратор лога*/
func (log *Logger) ConfigTrace(method, uri, contenttype string, time time.Duration) *Logger {
	log = &Logger{
		RequestMethod: method,
		ContentType:   contenttype,
		Time:          time,
		APIRoute:      uri,
	}
	return log
}

/*PrintTraceRoute - вывод данных по запросу*/
func (log *Logger) PrintTraceRoute() {
	if log.RequestStatus {
		log.OutputMethod = color.New(color.FgRed).Add(color.Underline)
	} else {
		log.OutputMethod = color.New(color.FgGreen)
	}
	log.printMethodType()
	log.printAPIRoute()
	log.printContentType()
	log.printRequestBody()
	log.printElapsedTime()
}

func (log *Logger) printMethodType() {
	log.OutputMethod.Print("\n[" + log.RequestMethod + "]\t")
	return
}

func (log *Logger) printAPIRoute() {
	log.OutputMethod.Print(log.APIRoute + "\t")
	return
}

func (log *Logger) printContentType() {
	log.OutputMethod.Print("Content-Type: " + log.ContentType + "\t")
	return
}

func (log *Logger) printRequestBody() {
	log.OutputMethod.Print("Body: " + log.RequestBody + "\t")
	return
}

func (log *Logger) printElapsedTime() {
	log.OutputMethod.Print("Elapsed Time: " + string(log.Time.Nanoseconds()))
	return
}
