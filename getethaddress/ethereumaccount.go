package getethaddress

import (
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"fmt"
	"deposit/accounts"
	"net/http"
	"deposit/logger"
)

var acc = &accounts.AccountSchema{}

func GetAddress(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	acc.Init()
	acc, err := acc.ByID(ps.ByName("id"))
	if err != nil {
		logger.Log.Panic(err)
	}
	values := map[string]string{"account": acc.EthAddress}
	jsonValue, _ := json.Marshal(values)
	fmt.Fprint(w, string(jsonValue))
	logger.Log.Printf("Request address: %s", acc.EthAddress)
	logger.Log.Printf("Request host: %s Request method: %s", r.Host, r.Method)
}
