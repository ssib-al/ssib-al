package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
)

// Link holds the schema definition for the Link entity.
type Link struct {
	ent.Schema
}

// Fields of the Link.
func (Link) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Unique(),
		field.String("domain").SchemaType(map[string]string{
			"type": "varchar(32)",
		}),
		field.String("uri").SchemaType(
			map[string]string{
				"type": "varchar(256)",
			},
		),
		field.String("target_url").SchemaType(
			map[string]string{
				"type": "varchar(256)",
			},
		),
		field.UUID("owner_id", uuid.UUID{}).SchemaType(
			map[string]string{
				"type": "uuid",
			},
		).Optional(),
	}
}

// Edges of the Link.
func (Link) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("links").
			Unique().
			Field("owner_id"),
	}
}

func (Link) Indexes() []ent.Index {
	return []ent.Index{
		// unique index.
		index.Fields("domain", "uri").
			Unique(),
	}
}
