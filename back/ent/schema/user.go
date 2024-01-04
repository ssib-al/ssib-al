package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable().Unique(),
		field.String("username").SchemaType(map[string]string{
			"type":        "varchar(32)",
			"primary_key": "true",
		}).Unique(),
		field.String("email").SchemaType(map[string]string{
			"type": "varchar(256)",
		}).Optional().Unique(),
		field.String("password_hash").SchemaType(map[string]string{
			"type": "varchar(512)",
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("links", Link.Type),
	}
}
