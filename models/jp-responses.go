package models

type JP_Todo_Resp struct {
	UserId    int16  `json:"userId,omitempty"`
	Id        int16  `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Completed bool   `json:"completed,omitempty"`
}
