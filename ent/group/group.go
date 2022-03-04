// Code generated by entc, DO NOT EDIT.

package group

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the group type in the database.
	Label = "group"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uid"
	// FieldGroupname holds the string denoting the groupname field in the database.
	FieldGroupname = "groupname"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldReadAt holds the string denoting the read_at field in the database.
	FieldReadAt = "read_at"
	// EdgeMembers holds the string denoting the members edge name in mutations.
	EdgeMembers = "members"
	// EdgeLikeFromUser holds the string denoting the like_from_user edge name in mutations.
	EdgeLikeFromUser = "like_from_user"
	// EdgeSaved holds the string denoting the saved edge name in mutations.
	EdgeSaved = "saved"
	// EdgeLikeFromGroup holds the string denoting the like_from_group edge name in mutations.
	EdgeLikeFromGroup = "like_from_group"
	// EdgeLikeTo holds the string denoting the like_to edge name in mutations.
	EdgeLikeTo = "like_to"
	// EdgePics holds the string denoting the pics edge name in mutations.
	EdgePics = "pics"
	// Table holds the table name of the group in the database.
	Table = "groups"
	// MembersTable is the table that holds the members relation/edge. The primary key declared below.
	MembersTable = "group_members"
	// MembersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	MembersInverseTable = "users"
	// LikeFromUserTable is the table that holds the like_from_user relation/edge. The primary key declared below.
	LikeFromUserTable = "user_like_to"
	// LikeFromUserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	LikeFromUserInverseTable = "users"
	// SavedTable is the table that holds the saved relation/edge. The primary key declared below.
	SavedTable = "user_save"
	// SavedInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	SavedInverseTable = "users"
	// LikeFromGroupTable is the table that holds the like_from_group relation/edge. The primary key declared below.
	LikeFromGroupTable = "group_like_to"
	// LikeToTable is the table that holds the like_to relation/edge. The primary key declared below.
	LikeToTable = "group_like_to"
	// PicsTable is the table that holds the pics relation/edge.
	PicsTable = "pics"
	// PicsInverseTable is the table name for the Pic entity.
	// It exists in this package in order to avoid circular dependency with the "pic" package.
	PicsInverseTable = "pics"
	// PicsColumn is the table column denoting the pics relation/edge.
	PicsColumn = "group_id"
)

// Columns holds all SQL columns for group fields.
var Columns = []string{
	FieldID,
	FieldGroupname,
	FieldBio,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldReadAt,
}

var (
	// MembersPrimaryKey and MembersColumn2 are the table columns denoting the
	// primary key for the members relation (M2M).
	MembersPrimaryKey = []string{"group_id", "user_id"}
	// LikeFromUserPrimaryKey and LikeFromUserColumn2 are the table columns denoting the
	// primary key for the like_from_user relation (M2M).
	LikeFromUserPrimaryKey = []string{"user_id", "group_id"}
	// SavedPrimaryKey and SavedColumn2 are the table columns denoting the
	// primary key for the saved relation (M2M).
	SavedPrimaryKey = []string{"user_id", "group_id"}
	// LikeFromGroupPrimaryKey and LikeFromGroupColumn2 are the table columns denoting the
	// primary key for the like_from_group relation (M2M).
	LikeFromGroupPrimaryKey = []string{"group_id", "like_from_group_id"}
	// LikeToPrimaryKey and LikeToColumn2 are the table columns denoting the
	// primary key for the like_to relation (M2M).
	LikeToPrimaryKey = []string{"group_id", "like_from_group_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultGroupname holds the default value on creation for the "groupname" field.
	DefaultGroupname string
	// DefaultBio holds the default value on creation for the "bio" field.
	DefaultBio string
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultReadAt holds the default value on creation for the "read_at" field.
	DefaultReadAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
