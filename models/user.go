package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	User struct {
		ID           bson.ObjectId `json:"id" bson:"_id"`
    Avatar       string         `json:"avatar" bson:"avatar"`
    UserID       string         `json:"user_id" bson:"user_id"`
    UserName      string        `json:"user_name" bson:"user_name"`
    Password    string    `json:"password" bson:"password"`
    MobileNumber string `json:"mobile_number" bson:"mobile_number"`
    UpdateTime  time.Time `json:"updatetime" bson:"updatetime"`
    CreateTime  time.Time `json:"createtime" bson:"createtime"`
    UpdateBy  string  `json:"updateuserid" bson:"updateuserid"`
    CreateBy  string  `json:"createuserid" bson:"createuserid"`
    Status    string  `json:"status" bson:"status"`
    ConverPicUrl  string  `json:"cover_picture_url" bson:"cover_picture_url"`
    Gender  bool    `json:"gender" bson:"gender"`
    Type    string  `json:"type" bson:"type"`
	}
)
