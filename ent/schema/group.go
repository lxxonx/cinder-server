package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(func() uuid.UUID { return uuid.New() }).Immutable().Unique().StorageKey("uid"),
		field.String("groupname").Default("").MaxLen(10).Unique(),
		field.String("bio").Default(""),
		// field.Strings("pics").Default([]string{""}),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("read_at").Default(time.Now),
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
