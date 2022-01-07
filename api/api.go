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
	"github.com/gin-gonic/gin"
)

type allPersonTimes []utils.PersonTimes

func initPersonTimes() allPersonTimes {
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

var personTimes = initPersonTimes()

func GetAllPersonTimes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, personTimes)
}

func CreatePersonTime(c *gin.Context) {
	var newPersonTimes utils.PersonTimes
	reqBody, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.Writer.WriteString("Double check to make sure data was sent properly")
	}

	json.Unmarshal(reqBody, &newPersonTimes)

	alreadyExists := false

	for i := 0; i < len(personTimes); i++ {
		if personTimes[i].UserName == newPersonTimes.UserName && personTimes[i].Week == newPersonTimes.Week {
			alreadyExists = true
		}
	}

	newPersonTimes.Id = strconv.Itoa(rand.Intn(9999999))

	if !alreadyExists {
		personTimes = append(personTimes, newPersonTimes)
		c.IndentedJSON(http.StatusCreated, personTimes)
	} else {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(c.Writer, "Time for this user for this week alreay exists\n")
	}
}

func FindPersonsTime(c *gin.Context) {
	var personsTime []utils.PersonTimes

	var err = true

	for i := 0; i < len(personTimes); i++ {
		if personTimes[i].UserName == c.Param("username") {
			personsTime = append(personsTime, personTimes[i])
			err = false
		}
	}

	if err {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(c.Writer, "There are no records for the user specified")
	} else {
		c.IndentedJSON(http.StatusOK, personsTime)
	}
}
