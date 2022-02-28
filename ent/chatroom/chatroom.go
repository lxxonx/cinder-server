// Code generated by entc, DO NOT EDIT.

package chatroom

import (
	"time"
)

const (
	// Label holds the string label denoting the chatroom type in the database.
	Label = "chat_room"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the createdat field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updatedat field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldReadAt holds the string denoting the readat field in the database.
	FieldReadAt = "read_at"
	// EdgeParticipants holds the string denoting the participants edge name in mutations.
	EdgeParticipants = "participants"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// Table holds the table name of the chatroom in the database.
	Table = "chat_rooms"
	// ParticipantsTable is the table that holds the participants relation/edge. The primary key declared below.
	ParticipantsTable = "chat_room_participants"
	// ParticipantsInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	ParticipantsInverseTable = "users"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "chat_messages"
	// MessagesInverseTable is the table name for the ChatMessage entity.
	// It exists in this package in order to avoid circular dependency with the "chatmessage" package.
	MessagesInverseTable = "chat_messages"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "room_id"
)

// Columns holds all SQL columns for chatroom fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldReadAt,
}

var (
	// ParticipantsPrimaryKey and ParticipantsColumn2 are the table columns denoting the
	// primary key for the participants relation (M2M).
	ParticipantsPrimaryKey = []string{"chat_room_id", "user_id"}
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
	// DefaultCreatedAt holds the default value on creation for the "createdAt" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updatedAt" field.
	DefaultUpdatedAt func() time.Time
	// DefaultReadAt holds the default value on creation for the "readAt" field.
	DefaultReadAt func() time.Time
)