package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ChatMessage holds the schema definition for the ChatMessage entity.
type ChatMessage struct {
	ent.Schema
}

// Fields of the ChatMessage.
func (ChatMessage) Fields() []ent.Field {
	return []ent.Field{
		field.String("uid").Unique(),
		field.String("message").Default(""),
		field.Int("room_id").Optional(),
		field.Int("user_id").Optional(),
		field.Time("createdAt").Default(time.Now),
		field.Time("readAt").Default(time.Now),
	}
}

// Edges of the ChatMessage.
func (ChatMessage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("message").
			Field("user_id").
			Unique(),
		edge.From("room", ChatRoom.Type).
			Ref("messages").
			Field("room_id").
			Unique(),
	}
}
