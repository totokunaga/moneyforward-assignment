package types

type GetUserType struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	AccountIds []int  `json:"account_ids"`
}

type GetUserAccountType struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}
