package main

import (
    "fmt"
    "strconv"
    "os"
    "math"
    "encoding/json"
	"log"
    "github.com/VIZ-Blockchain/viz-go-lib"
	"github.com/VIZ-Blockchain/viz-go-lib/operations"
    "github.com/VIZ-Blockchain/viz-go-lib/types"
)

type User struct {
    Login         string
    RegularKey       string
    ActiveKey string
}

func vizSender(p int, acc User) {
	client, _ := viz.NewClient("https://viz-node.dpos.space")
	defer client.Close()

	var trx []operations.Operation

if p == 1 {
        ans, err := client.API.GetAccounts(acc.Login)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
        config_mass, err := client.API.GetConfig()
        if (err != nil) {
    return
}

props, err := client.API.GetDynamicGlobalProperties();
if (err != nil) {
    return
}

last_vote_time := ans[0].LastVoteTime
current_time := props.Time.Unix() * 1000
last_vote_seconds := last_vote_time.Unix() * 1000
fastpower := float64(10000 / float64(config_mass.EnergyRegenerationSeconds))
    after_vote_time := float64(current_time -last_vote_seconds)
volume_not := (float64(ans[0].Energy) + (after_vote_time /1000) * fastpower)/100; //расчет текущей Voting Power
     volume := math.Round(volume_not*100)/100
var charge float64 = 0
    if (volume >=100) {
    charge = 100
    } else {
        charge= volume
    }

        fmt.Println("Энергия на данный момент:", charge)
    }
        
    client.SetKeys(&viz.Keys{PKey: []string{acc.RegularKey}})
    fmt.Println("Укажите получателя награды")
    var to string
    fmt.Scanln(&to)
    fmt.Println("Процент энергии")
    var energy float64
fmt.Scanln(&energy)
fmt.Println("Укажите memo (заметку к награде)")
var memo string
fmt.Scanln(&memo)

beneficiaries := []types.Beneficiary{
		{Account: "denis-skripnik", Weight: 500},  // 5%
	}
	tx := &operations.AwardOperation{
		Initiator:      acc.Login,
		Receiver:       to,
		Energy:         uint16(energy * 100),
		CustomSequence: 0,
		Memo:           memo,
		Beneficiaries:  beneficiaries,
	}
	trx = append(trx, tx)

	resp, err := client.SendTrx(acc.Login, trx)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
		fmt.Println("Answer : ", resp)
	}
    } else if p == 2 {
        ans, err := client.API.GetAccounts(acc.Login)
	if err != nil {
		fmt.Println("Error : ", err)
	} else {
        fmt.Println("Баланс: ", ans[0].Balance)
    }

        client.SetKeys(&viz.Keys{AKey: []string{acc.ActiveKey}})

    fmt.Println("Укажите получателя перевода")
    var to string
    fmt.Scanln(&to)
    fmt.Println("Сумма")
var amount float64
fmt.Scanln(&amount)
fmt.Println("Укажите memo (заметку к платежу)")
var memo string
fmt.Scanln(&memo)

tx := &operations.TransferOperation{
        From:      acc.Login,
        To:       to,
		Amount:         viz.SetAsset(amount, "VIZ"),
		Memo:           string(memo),
	}
	trx = append(trx, tx)

    resp, err := client.SendTrx(acc.Login, trx)
	if err != nil {
        fmt.Println("Error : ", err)
        } else {
		fmt.Println("Answer : ", resp)
	}

}

fmt.Println("Для закрытия приложения нажмите Enter")
var conferm string
fmt.Scanln(&conferm)
}

func accountMenu(acc User) {
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

func selectAccount(u []User) {
	var choice int
fmt.Scanln(&choice)
if choice == 0 {
	os.Exit(2)
} else if choice > 0 && len(u) >= 1 {
accountMenu(u[choice-1])
}
}

func main() {
    var users []User

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