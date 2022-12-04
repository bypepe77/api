// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "age", Type: field.TypeInt},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
		{Name: "username", Type: field.TypeString, Unique: true, Size: 25},
		{Name: "password", Type: field.TypeString, Default: "unknown"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "active", Type: field.TypeBool, Default: false},
		{Name: "follows_count", Type: field.TypeInt, Default: 0},
		{Name: "following_count", Type: field.TypeInt, Default: 0},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		UsersTable,
	}
)

func init() {
}
