package entities

type Segment struct {
	tableName struct{} `sql:"segments,alias:segments" pg:",discard_unknown_columns"`
	Id        int      `sql:"id" pg:"id" db:"id" json:",omitempty"`
	Slug      string   `sql:"slug,notnull" pg:"slug" db:"slug" json:"slug"`
}
