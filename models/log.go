package models

import "gopkg.in/mgo.v2/bson"

type (
	// Log is
	Log struct {
		ID   bson.ObjectId `json:"id" bson:"_id,omitempty"`
		Type string        `json:"type" bson:"type"`
		Info string        `json:"info" bson:"info"`
	}
)
