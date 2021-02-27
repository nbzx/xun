package grammar

import (
	"time"

	"github.com/jmoiron/sqlx"
)

// Grammar the Grammar inteface
type Grammar interface {
	DBName() string
	SchemaName() string

	Exists(name string, db *sqlx.DB) bool
	Get(table *Table, db *sqlx.DB) error
	Create(table *Table, db *sqlx.DB) error
	Alter(table *Table, db *sqlx.DB) error
	Drop(name string, db *sqlx.DB) error
	DropIfExists(name string, db *sqlx.DB) error
	Rename(old string, new string, db *sqlx.DB) error
	GetColumnListing(dbName string, tableName string, db *sqlx.DB) ([]*Column, error)
}

// Quoter the database quoting query text intrface
type Quoter interface {
	ID(name string, db *sqlx.DB) string
	VAL(v interface{}, db *sqlx.DB) string // operates on both string and []byte and int or other types.
}

// Table the table struct
type Table struct {
	DBName        string    `db:"db_name"`
	SchemaName    string    `db:"schema_name"`
	Name          string    `db:"table_name"`
	Comment       string    `db:"table_comment"`
	Type          string    `db:"table_type"`
	Engine        string    `db:"engine"`
	CreateTime    time.Time `db:"create_time"`
	CreateOptions string    `db:"create_options"`
	Collation     string    `db:"collation"`
	Charset       string    `db:"charset"`
	Rows          int       `db:"table_rows"`
	RowLength     int       `db:"avg_row_length"`
	IndexLength   int       `db:"index_length"`
	AutoIncrement int       `db:"auto_increment"`
	Primary       *Primary
	ColumnMap     map[string]*Column
	IndexMap      map[string]*Index
	Columns       []*Column
	Indexes       []*Index
	Commands      []*Command
}

// Column the table Column
type Column struct {
	DBName            string      `db:"db_name"`
	TableName         string      `db:"table_name"`
	Name              string      `db:"name"`
	Position          int         `db:"position"`
	Default           interface{} `db:"default"`
	Nullable          bool        `db:"nullable"`
	IsUnsigned        bool        `db:"unsigned"`
	Type              string      `db:"type"`
	Length            *int        `db:"length"`
	OctetLength       *int        `db:"octet_length"`
	Precision         *int        `db:"precision"`
	Scale             *int        `db:"scale"`
	DatetimePrecision *int        `db:"datetime_precision"`
	Charset           *string     `db:"charset"`
	Collation         *string     `db:"collation"`
	Key               *string     `db:"key"`
	Extra             *string     `db:"extra"`
	Comment           *string     `db:"comment"`
	Primary           bool        `db:"primary"`
	Table             *Table
	Indexes           []*Index
}

// Index the talbe index
type Index struct {
	DBName       string  `db:"db_name"`
	TableName    string  `db:"table_name"`
	ColumnName   string  `db:"column_name"`
	Name         string  `db:"index_name"`
	SEQ          int     `db:"seq_in_index"`
	SeqColumn    int     `db:"seq_in_column"`
	Collation    string  `db:"collation"`
	Nullable     bool    `db:"nullable"`
	Unique       bool    `db:"unique"`
	Primary      bool    `db:"primary"`
	SubPart      int     `db:"sub_part"`
	Type         string  `db:"type"`
	IndexType    string  `db:"index_type"`
	Comment      *string `db:"comment"`
	IndexComment *string `db:"index_comment"`
	Table        *Table
	Columns      []*Column
}

// Primary the table primary key
type Primary struct {
	DBName    string `db:"db_name"`
	TableName string `db:"table_name"`
	Name      string `db:"primary_name"`
	Table     *Table
	Columns   []*Column
}

// Command The Command that should be run for the table.
type Command struct {
	Name   string        // The command name
	Params []interface{} // The command parameters
}
