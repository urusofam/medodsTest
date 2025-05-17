package models

type RefreshToken struct {
	ID        int    `db:"id"`
	UserGUID  string `db:"user_guid"`
	TokenHash string `db:"token_hash"`
	UserAgent string `db:"user_agent"`
	IP        string `db:"ip"`
	IsValid   bool   `db:"is_valid"`
}
