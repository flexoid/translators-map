// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TranslatorsColumns holds the columns for the "translators" table.
	TranslatorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "language", Type: field.TypeString},
		{Name: "address", Type: field.TypeString},
		{Name: "contacts", Type: field.TypeString},
		{Name: "details_url", Type: field.TypeString},
		{Name: "latitude", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "numeric"}},
		{Name: "longitude", Type: field.TypeFloat64, Nullable: true, SchemaType: map[string]string{"postgres": "numeric"}},
	}
	// TranslatorsTable holds the schema information for the "translators" table.
	TranslatorsTable = &schema.Table{
		Name:       "translators",
		Columns:    TranslatorsColumns,
		PrimaryKey: []*schema.Column{TranslatorsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "translator_name_language",
				Unique:  true,
				Columns: []*schema.Column{TranslatorsColumns[1], TranslatorsColumns[2]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TranslatorsTable,
	}
)

func init() {
}
