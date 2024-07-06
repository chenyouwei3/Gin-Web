package model

import "time"

type OperationLog struct {
	Id        int64
	Username  string
	Ip        string
	Method    string
	Query     string
	Path      string
	Status    int
	StartTime time.Time
	TimeCost  int64
	UserAgent string
	Errors    string
}
