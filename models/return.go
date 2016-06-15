package models

type (
	Recs interface {
	}

	Result struct {
		Error bool   `json:"err" bson:"err"`
		Msg   string `json:"msg" bson:"msg"`
		Data  []Recs `json:"data" bson:"data"`
	}

  Auth struct {
    Key string  `json:"key" bson:"key"`
    UserID string `json:"user_id" bson:"user_id"`
    Msg string `json:"msg" bson:"msg"`
  }
)
