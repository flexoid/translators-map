package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Translator holds the schema definition for the Translator entity.
type Translator struct {
	ent.Schema
}

// Fields of the Translator.
func (Translator) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("language"),
		field.String("address"),
		field.String("contacts"),
		field.String("details_url"),
	}
}

func (Translator) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name", "language").Unique(),
	}
}

// Edges of the Translator.
func (Translator) Edges() []ent.Edge {
	return nil
}
