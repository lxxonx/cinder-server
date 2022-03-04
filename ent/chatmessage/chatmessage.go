// Code generated by entc, DO NOT EDIT.

package chatmessage

import (
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the chatmessage type in the database.
	Label = "chat_message"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uid"
	// FieldMessage holds the string denoting the message field in the database.
	FieldMessage = "message"
	// FieldRoomID holds the string denoting the room_id field in the database.
	FieldRoomID = "room_id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldReadAt holds the string denoting the read_at field in the database.
	FieldReadAt = "read_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeRoom holds the string denoting the room edge name in mutations.
	EdgeRoom = "room"
	// Table holds the table name of the chatmessage in the database.
	Table = "chat_messages"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "chat_messages"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// RoomTable is the table that holds the room relation/edge.
	RoomTable = "chat_messages"
	// RoomInverseTable is the table name for the ChatRoom entity.
	// It exists in this package in order to avoid circular dependency with the "chatroom" package.
	RoomInverseTable = "chat_rooms"
	// RoomColumn is the table column denoting the room relation/edge.
	RoomColumn = "room_id"
)

// Columns holds all SQL columns for chatmessage fields.
var Columns = []string{
	FieldID,
	FieldMessage,
	FieldRoomID,
	FieldUserID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldReadAt,
}

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
	// DefaultMessage holds the default value on creation for the "message" field.
	DefaultMessage string
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
