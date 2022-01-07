package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/Shanedell/ga-api-go/utils"
)

type allPersonTimes2 []utils.PersonTimes

func initPersonTimes2() allPersonTimes {
	var personTimesData = allPersonTimes{}

	jsonFile, err := os.Open("data/example.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &personTimesData)

	for i := 0; i < len(personTimesData); i++ {
		var t_pt_id = strconv.Itoa(rand.Intn(9999999))
		personTimesData[i].Id = t_pt_id
	}

	return personTimesData
}

var personTimes2 = initPersonTimes()

func GetPersonTimes2(w http.ResponseWriter, r *http.Request) {
	json.MarshalIndent(json.NewEncoder(w).Encode(personTimes), "", " ")
}

func CreatePersonTime2(w http.ResponseWriter, r *http.Request) {
	var newPersonTimes utils.PersonTimes
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Check data data entered")
	}

	json.Unmarshal(reqBody, &newPersonTimes)

	var alreadyExists = false

	for i := 0; i < len(personTimes); i++ {
		if personTimes[i].UserName == newPersonTimes.UserName && personTimes[i].Week == newPersonTimes.UserName {
			alreadyExists = true
		}
	}

	if !alreadyExists {
		personTimes = append(personTimes, newPersonTimes)
		w.WriteHeader(http.StatusCreated)
	} else {
		fmt.Fprintf(w, "Time for this user for this week alreay exists")
		w.WriteHeader(http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(personTimes)
}
