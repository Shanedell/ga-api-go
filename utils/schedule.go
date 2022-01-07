package utils

import (
	"math/rand"
	"strconv"
)

type Schedule struct {
	Id                string        `json:"id"`
	FirstDayOfTheWeek string        `json:"firstDayOfTheWeek"`
	Schedule          []PersonTimes `json:"schedule"`
	Fitness           int           `json:"fitness"`
	Meeting           string        `json:"meeting"`
	MeetingNeeded     bool          `json:"meetingNeeded"`
}

func NewSchedule(firstDayOfTheWeek string, schedule []PersonTimes, fitness int, meeting string, meetingNeed bool, id string) Schedule {
	var s_id = strconv.Itoa(rand.Intn(9999999))

	s := Schedule{s_id, firstDayOfTheWeek, schedule, fitness, meeting, meetingNeed}
	return s
}

func (schedule Schedule) CalculateFitness(requestTimes []PersonTimes) {
	var fitness = 0

	for _, actualTime := range schedule.Schedule {
		for _, requestedTime := range requestTimes {
			if requestedTime == actualTime {
				fitness += 100
				break
			}

			if requestedTime.DaysOff(actualTime) {
				fitness -= 200
				break
			}
		}
	}

	if schedule.MeetingNeeded && (schedule.Meeting == "2:00pm-2:30pm" || schedule.Meeting == "2:30pm-3:00pm" || schedule.Meeting == "3:00pm-3:30pm") {
		fitness += 111
	} else if schedule.MeetingNeeded && schedule.Meeting != "2:00pm-2:30pm" && schedule.Meeting != "2:30pm-3:00pm" && schedule.Meeting != "3:00pm-3:30pm" {
		fitness -= 333
	}
}
