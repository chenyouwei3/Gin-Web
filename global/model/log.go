package model

type Log struct {
	Id        int64       `bson:"_id,omitempty" json:"_id,omitempty"`
	User      interface{} `bson:"user" json:"user"`
	Path      string      `bson:"path" json:"path"`
	Method    string      `bson:"method" json:"method"`
	Status    int         `bson:"status" json:"status"`
	Query     string      `bson:"query" json:"query"`
	Body      interface{} `bson:"body" json:"body"`
	Ip        string      `bson:"ip" json:"ip"`
	UserAgent string      `bson:"userAgent" json:"userAgent"`
	Errors    string      `bson:"errors" json:"errors"`
	Cost      string      `bson:"cost" json:"cost"`
}
