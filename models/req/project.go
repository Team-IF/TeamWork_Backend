package req

type Project struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Password    string `json:"password"`
}

type ProjectCreate struct {
	Project
}

type ProjectUpdate struct {
	ID uint `json:"id"`
	Project
}
