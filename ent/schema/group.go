package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique(),
		field.String("groupname").Default(""),
		field.String("bio").Default(""),
		// field.Strings("pics").Default([]string{""}),
		field.Time("createdAt").Default(time.Now),
		field.Time("updatedAt").Default(time.Now),
		field.Time("readAt").Default(time.Now),
	} // Uid        string        `json:"uid"`
	// GroupName  string        `json:"groupName"`
	// Pics       []interface{} `json:"pics"`
	// CreateTime time.Time     `json:"createTime"`
	// Bio        string        `json:"bio"`
	// Members    []User        `json:"members"`
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", User.Type),
		edge.From("like_from_user", User.Type).
			Ref("like_to"),
		edge.From("saved", User.Type).
			Ref("save"),
		edge.To("like_to", Group.Type).
			From("like_from_group"),
		edge.To("pics", Pic.Type),
	}
}
