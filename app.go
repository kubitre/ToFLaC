package main

import (
	"bufio"
	"os"
	"strings"
	"time"

	"tflac_labs/bankomat"
	"tflac_labs/bankomat/states"
	"tflac_labs/regularAutomat"

	tm "github.com/buger/goterm"
)

// func init() {
// 	server := Routes.EntrypointRoute{}
// 	server.CreateAllRoutes()
// }

func testStatePattern() {
	// automat := states.NewAutomat()
}

/*ConfigureBankomat - конфигурация автомата*/
func ConfigureBankomat() *states.AutomatState {
	return CreateNewBankomat(
		&bankomat.UniverseBank{
			Accounts: []bankomat.UserAccount{
				bankomat.UserAccount{
					ID:      1,
					Balance: 999999,
				},
			},
		},
		999,
	)
}

/*CreateNewBankomat - создание нового банкомата*/
func CreateNewBankomat(
	usersWithBalance *bankomat.UniverseBank,
	amountPapersAvailable int,
) *states.AutomatState {
	automat := &states.AutomatState{
		Statess: &states.AllStates{
			ZeroInstance: &states.ZeroState{},
			Z1Instance:   &states.Z1State{},
			Z2Instance:   &states.Z2State{},
			Z21Instance:  &states.Z21State{},
			Z22Instance:  &states.Z22State{},
			Z3Instance:   &states.Z3State{},
			Z31Instance:  &states.Z31State{},
			Z32Instance:  &states.Z32State{},
		},
		Bank:         usersWithBalance,
		State:        &states.ZeroState{},
		AmountPapers: amountPapersAvailable,
	}

	return automat
}

func BankomatEntryPoint() {
	tm.Clear() // Clear current screen
	tm.MoveCursor(1, 1)
	tm.Println("Server was started!")
	automat := ConfigureBankomat()

	reader := bufio.NewReader(os.Stdin)
	// tm.MoveCursor(1, 1)
	tm.Println("Bankomat")
	tm.Println("---------------------")

	tm.Background("string", tm.RED)

	for {
		tm.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			tm.Println("hello, Yourself")
		}
		tm.Flush() // Call it every time at the end of rendering

		time.Sleep(time.Second)

		// elementsDebug := []string{"1", "2", "5000", "3"}
		// for _, element := range elementsDebug {
		// 	automat.NewEnter(element)
		// }

		automat.NewEnter(text)
		tm.Clear()
		tm.MoveCursor(1, 1)
		tm.Println(automat.GetLog())

	}
}

func RegularExpEntryPoint() {
	regexp := &regularAutomat.RegularExp{}
	regexp = regexp.New()

	regexp.Reader()
}

func main() {
	RegularExpEntryPoint()
	// BankomatEntryPoint()
}
