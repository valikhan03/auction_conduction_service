package delivery

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"auction_conduction/errors"
)

type RemoteService struct{
	Host string
	
}

func NewRemoteService(host string) *RemoteService{
	return &RemoteService{
		Host: host,
	}
}


type AccessResp struct{
	Result string `json:"result"`
}

func(r *RemoteService) CheckAccess(auction_id string, user_id string) (bool, error){
	check := fmt.Sprintf("%s/%s/%s", r.Host, auction_id, user_id)
	res, err := http.Get(check)
	if err != nil{
		log.Println(err)
		return false, err
	}

	var accessRes AccessResp

	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&accessRes)
	if err != nil{
		log.Println(err)
		return false, errors.CHECK_ACCESS_ERR
	}

	if accessRes.Result == "true"{
		return true, nil
	}else{
		return false, nil
	}
}