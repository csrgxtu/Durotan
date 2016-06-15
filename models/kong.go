package models

type (
  KeyAuthKey struct {
    Key string  `json:"key" bson:"key"`
    ConsumerID  string `json:"consumer_id" bson:"consumer_id"`
    ID string `json:"id" bson:"id"`
    CreateAt int64 `json:"created_at" bson:"created_at"`
  }
)
