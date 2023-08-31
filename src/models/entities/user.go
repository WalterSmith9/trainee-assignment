package entities

type User struct {
	tableName struct{} `sql:"users,alias:users" pg:",discard_unknown_columns"`
	Id        int      `sql:"id" pg:"id" db:"id" json:",omitempty"`
	UserId    int      `sql:"user_id,notnull" pg:"user_id" db:"user_id" json:"user_id"`
}
