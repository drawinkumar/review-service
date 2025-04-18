package jobs

import "gorm.io/gorm"

type JobStatus string

const (
	Unprocessed JobStatus = "unprocessed"
	Inprogress  JobStatus = "inprogress"
	Processed   JobStatus = "processed"
)

type Jobs struct {
	gorm.Model
	FileName string    `gorm:"size:255;index"`
	Status   JobStatus `gorm:"type:enum('unprocessed','inprogress','processed');default:'unprocessed';index"`
}
