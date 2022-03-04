package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ChatRoom holds the schema definition for the ChatRoom entity.
type ChatRoom struct {
	ent.Schema
}

// Fields of the ChatRoom.
func (ChatRoom) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID { return uuid.New() }).Immutable().Unique().StorageKey("uid"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("read_at").Default(time.Now),
	}
}

// Edges of the ChatRoom.
func (ChatRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("participants", User.Type),
		edge.To("messages", ChatMessage.Type),
	}
}
