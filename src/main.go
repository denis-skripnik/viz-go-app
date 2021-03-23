package main

import (
    "fmt"
    "strconv"
    "os"
    "encoding/json"
    "log"
    "viz-app/account"
)

func vizSender(p int, acc account.User) {
if p == 1 {
    account.Award(acc)
    } else if p == 2 {
account.Transfer(acc)
    }

fmt.Println("Для закрытия приложения нажмите Enter")
var conferm string
fmt.Scanln(&conferm)
}

func accountMenu(acc account.User) {
fmt.Println("Выбран аккаунт", acc.Login)
fmt.Println(`Выберите, что хотите сделать
1. Наградить кого-то
2. Перевести VIZ
0. Закрыть программу`)
var p int
fmt.Scanln(&p)
if p > 0 {
    vizSender(p, acc)
}
}

func selectAccount(u []account.User) {
	var choice int
fmt.Scanln(&choice)
if choice == 0 {
	os.Exit(2)
} else if choice > 0 && len(u) >= 1 {
accountMenu(u[choice-1])
}
}

func main() {
    var users []account.User

	file, err := os.Open("config.json")
    if err != nil {
        // здесь перехватывается ошибка
        return
    }
    defer file.Close()
	
    stat, err := file.Stat()
    if err != nil {
        return
    }

	    // чтение файла
		bs := make([]byte, stat.Size())
		_, err = file.Read(bs)
		if err != nil {
			return
		}
	
    err2 := json.Unmarshal(bs, &users)

    if err2 != nil {

        log.Fatal(err2)
    }
fmt.Println("Выберите аккаунт, нажав на соответствующую клавишу и Enter. Выход из приложения - 0")
    for i := range users {
		var number string = strconv.Itoa(i + 1)
		fmt.Println(number + ".", users[i].Login)
    }
selectAccount(users)
}