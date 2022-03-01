package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Pic holds the schema definition for the Pic entity.
type Pic struct {
	ent.Schema
}

// Fields of the Pic.
func (Pic) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").Unique(),
		field.Int("user_id").Optional(),
		field.Int("group_id").Optional(),
		field.String("url").NotEmpty(),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now),
		field.Time("readAt").Default(time.Now),
	} // Uid
}

// Edges of the Pic.
func (Pic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("pics").
			Unique().Field("user_id"),
		edge.From("group", Group.Type).
			Ref("pics").
			Unique().Field("group_id"),
	}
}
