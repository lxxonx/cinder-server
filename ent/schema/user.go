package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("username").NotEmpty(),
		field.Bytes("password").Sensitive(),
		field.String("uni").NotEmpty(),
		field.String("dep").NotEmpty(),
		field.String("bio").NotEmpty(),
		field.Strings("pics").Default([]string{""}),
		field.String("group_id").Optional(),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now),
		field.Time("readAt").Default(time.Now),
	}
	// Username string        `json:"username"`
	// Uni      string        `json:"uni"`
	// Dep      string        `json:"dep"`
	// Bio      string        `json:"bio"`
	// Pics     []interface{} `json:"pics"`
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("friends", User.Type),
		edge.To("like_to", Group.Type),
		edge.To("save", Group.Type),
		edge.From("group", Group.Type).
			Ref("members").
			Unique().Field("group_id"),
		edge.From("chatroom", ChatRoom.Type).
			Ref("participants"),
		edge.To("message", ChatMessage.Type),
	}
}
