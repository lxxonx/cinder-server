package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ChatRoom holds the schema definition for the ChatRoom entity.
type ChatRoom struct {
	ent.Schema
}

// Fields of the ChatRoom.
func (ChatRoom) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").Unique(),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now),
		field.Time("readAt").Default(time.Now),
	}
}

// Edges of the ChatRoom.
func (ChatRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("participants", User.Type),
		edge.To("messages", ChatMessage.Type),
	}
}
