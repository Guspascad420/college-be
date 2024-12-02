package models

type LoginResponse struct {
	Token string `json:"token"`
}

type UserProfileResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	NIM      string `json:"nim"`
	Angkatan int    `json:"angkatan"`
	Major    string `json:"prodi"`
}
