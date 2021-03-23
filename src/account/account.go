package account

import (
	"fmt"
	"math"
    "github.com/VIZ-Blockchain/viz-go-lib"
	"github.com/VIZ-Blockchain/viz-go-lib/operations"
    "github.com/VIZ-Blockchain/viz-go-lib/types"
)

type User struct {
    Login         string
    RegularKey       string
    ActiveKey string
}

func Award(acc User) {
	client, _ := viz.NewClient("https://viz-node.dpos.space")
	defer client.Close()

	var trx []operations.Operation

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
}

func Transfer(acc User) {
	client, _ := viz.NewClient("https://viz-node.dpos.space")
	defer client.Close()

	var trx []operations.Operation

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