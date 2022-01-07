package utils

type Scheduler struct {
	Generation       int
	BestSchedule     Schedule
	TimeOptions      []string
	Meeting_Required bool
}
