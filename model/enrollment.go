package model

import "time"

type Enrollment struct {
	Id                string `json:"id"`
	Course            Course `json:"course"`
	EnrollmentDetails []EnrollmentDetail
	Status            string    `json:"status"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

type EnrollmentDetail struct {
	Id           string    `json:"id"`
	EnrollmentId string    `json:"enrollmentId"`
	User         User      `json:"user"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
