package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Pic holds the schema definition for the Pic entity.
type Pic struct {
	ent.Schema
}

// Fields of the Pic.
func (Pic) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID { return uuid.New() }).Immutable().Unique().StorageKey("uid"),
		field.String("user_id").Optional(),
		field.UUID("group_id", uuid.UUID{}).Optional(),
		field.String("url").NotEmpty(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("read_at").Default(time.Now),
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
