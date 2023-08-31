package entities

type Record struct {
	tableName struct{} `sql:"users_segments,alias:users_segments" pg:",discard_unknown_columns"`
	UserId    int      `sql:"user_id,notnull" pg:"user_id" db:"user_id" json:"user_id"`
	Segment   int      `sql:"segment_id" pg:"segment_id" db:"segment_id" json:",omitempty"`
}
