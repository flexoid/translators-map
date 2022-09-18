package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
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
		field.Int("external_id"),
		field.String("language"),
		field.String("address").Optional(),
		field.Bytes("address_sha"),
		field.String("details_url"),
		field.Float("latitude").Optional().SchemaType(map[string]string{
			dialect.Postgres: "numeric",
		}),
		field.Float("longitude").Optional().SchemaType(map[string]string{
			dialect.Postgres: "numeric",
		}),
		field.Time("created_at").Default(time.Now).Immutable().Optional(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now).Optional(),
	}
}

func (Translator) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("external_id", "language").Unique(),
	}
}

// Edges of the Translator.
func (Translator) Edges() []ent.Edge {
	return nil
}
