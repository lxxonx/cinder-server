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
		field.String("id").Unique().Immutable().StorageKey("uid"),
		field.String("actual_name").NotEmpty().Comment("본명").Immutable(),
		field.Int("phone_number").Comment("전화번호").Immutable().Optional().Unique(),
		field.String("username").NotEmpty().Comment("아이디").Immutable().Unique(),
		field.String("gender").Immutable(),
		field.String("uni").NotEmpty().Comment("대학교"),
		field.String("dep").NotEmpty().Comment("단과대"),
		field.String("bio").Optional(),
		field.Int("birth_year").Min(1900).Max(time.Now().Year() - 18),
		field.Bool("is_verified").Default(false),
		field.String("status").Default("phone_number_verified"),
		field.Int("max_group").Default(3),
		field.String("avatar").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("read_at").Default(time.Now),
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
		edge.To("friends_req", User.Type).
			From("requests"),
		edge.To("like_to", Group.Type),
		edge.To("save", Group.Type),
		edge.From("group", Group.Type).
			Ref("members"),
		edge.From("chatroom", ChatRoom.Type).
			Ref("participants"),
		edge.To("message", ChatMessage.Type),
		edge.To("pics", Pic.Type),
	}
}
