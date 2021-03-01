package models

// User struct model
type User struct {
	UserID      string `json:"user_id"`
	UserName    string `json:"username"`
	DateOfBirth string `json:"date_of_birth"`
	NoKtp       string `json:"no_ktp"`
	Job         string `json:"job"`
	Education   string `json:"education"`
}

// Job struct model
type Job struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Education struct model
type Education struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
