package bankomat

/*UserAccount - аккаунт в банкомате*/
type UserAccount struct {
	ID      int
	Balance int
}

//GetCash - снять наличные в размере amount
func (ac *UserAccount) GetCash(amount int) {
	ac.Balance -= amount
}

//GetCurrentBalance - получение текущего баланса юзвера
func (ac *UserAccount) GetCurrentBalance() int {
	return ac.Balance
}

/*UniverseBank - хранилище денег аккаунтов*/
type UniverseBank struct {
	Accounts []UserAccount
}

/*GetCashClient - удалить у клиента id amount денег*/
func (ub *UniverseBank) GetCashClient(id, amount int) {
	var elementU UserAccount
	for index, element := range ub.Accounts {
		if element.ID == id {
			elementU = element
			elementU.GetCash(amount)
			ub.Accounts[index] = elementU
		}
	}
}
