package req

type ProjectCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Password    string `json:"password"`
}
