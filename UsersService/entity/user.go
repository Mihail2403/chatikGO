package entity

type User struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ImgBase64 string `json:"img"`
	UserId    int    `json:"user_id"`
}
