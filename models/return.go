package models

type (
  // Result struct {
  //   Msg string  `json:"msg" bson:"msg"`
  //   Data Auth `json:"data" bson:"data"`
  // }

  Recs interface {}

	Result struct {
		Msg   string `json:"msg" bson:"msg"`
		Data  []Recs `json:"data,omitempty" bson:"data"`
	}

  Auth struct {
    Token string `json:"token" bson:"token"`
    RefreshToken string `json:"refresh_token" bson:"refresh_token"`
    AccountID string `json:"aid" bson:"aid"`
  }
)
