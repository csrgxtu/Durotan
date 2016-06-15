package models

type (
	Recs interface {
	}

	Result struct {
		Error bool   `json:"err" bson:"err"`
		Msg   string `json:"msg" bson:"msg"`
		Data  []Recs `json:"data" bson:"data"`
	}
)
