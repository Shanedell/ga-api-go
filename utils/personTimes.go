package utils

import (
	"encoding/json"
	"math/rand"
	"strconv"
)

type PersonTimes struct {
	Id        string `json:"id"`
	Week      string `json:"week"`
	UserName  string `json:"username"`
	Monday    string `json:"monday"`
	Tuesday   string `json:"tuesday"`
	Wednesday string `json:"wednesday"`
	Thursday  string `json:"thursday"`
	Friday    string `json:"friday"`
	Status    string `json:"status"`
}

func CreatePersonTimes(week string, userName string, monday string, tuesday string, wednesday string, thursday string, friday string) PersonTimes {
	var pt_id = strconv.Itoa(rand.Intn(9999999))
	var status = ""

	pt := PersonTimes{pt_id, week, userName, monday, tuesday, wednesday, thursday, friday, status}
	return pt
}

func (personTime PersonTimes) DaysOff(other PersonTimes) bool {
	return (personTime.Week == other.Week) && (personTime.UserName == other.UserName) && ((personTime.Monday != other.Monday) || (personTime.Tuesday != other.Tuesday) || (personTime.Wednesday != other.Wednesday) || (personTime.Thursday != other.Thursday) || (personTime.Friday != other.Friday))
}

func (personTime PersonTimes) PersonTimesToJson() string {
	j, _ := json.MarshalIndent(personTime, "", " ")
	return string(j)
}
