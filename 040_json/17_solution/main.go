package main

import (
	json "encoding/json"
	"fmt"
	"log"
)

type status struct {
	Code    string `json:"Code"`
	Descrip int    `json:"Descrip"`
}

type statuses []status

func main() {
	rcvd := `[{"Code":"StatusOK","Descrip":200},{"Code":"StatusMovedPermanently","Descrip":301},{"Code":"StatusFound","Descrip":302},{"Code":"StatusSeeOther","Descrip":303},{"Code":"StatusTemporaryRedirect","Descrip":307},{"Code":"StatusBadRequest","Descrip":400},{"Code":"StatusUnauthorized","Descrip":401},{"Code":"StatusPaymentRequired","Descrip":402},{"Code":"StatusForbidden","Descrip":403},{"Code":"StatusNotFound","Descrip":404},{"Code":"StatusMethodNotAllowed","Descrip":405},{"Code":"StatusTeapot","Descrip":418},{"Code":"StatusInternalServerError","Descrip":500}]`
	var data statuses
	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range data {
		fmt.Println(v.Code, "-", v.Descrip)
	}
}
