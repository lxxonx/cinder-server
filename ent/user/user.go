// Code generated by entc, DO NOT EDIT.

package user

import (
	"time"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "uid"
	// FieldActualName holds the string denoting the actual_name field in the database.
	FieldActualName = "actual_name"
	// FieldPhoneNumber holds the string denoting the phone_number field in the database.
	FieldPhoneNumber = "phone_number"
	// FieldUsername holds the string denoting the username field in the database.
	FieldUsername = "username"
	// FieldGender holds the string denoting the gender field in the database.
	FieldGender = "gender"
	// FieldUni holds the string denoting the uni field in the database.
	FieldUni = "uni"
	// FieldDep holds the string denoting the dep field in the database.
	FieldDep = "dep"
	// FieldBio holds the string denoting the bio field in the database.
	FieldBio = "bio"
	// FieldBirthYear holds the string denoting the birth_year field in the database.
	FieldBirthYear = "birth_year"
	// FieldIsVerified holds the string denoting the is_verified field in the database.
	FieldIsVerified = "is_verified"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldMaxGroup holds the string denoting the max_group field in the database.
	FieldMaxGroup = "max_group"
	// FieldAvatar holds the string denoting the avatar field in the database.
	FieldAvatar = "avatar"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldReadAt holds the string denoting the read_at field in the database.
	FieldReadAt = "read_at"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// EdgeRequests holds the string denoting the requests edge name in mutations.
	EdgeRequests = "requests"
	// EdgeFriendsReq holds the string denoting the friends_req edge name in mutations.
	EdgeFriendsReq = "friends_req"
	// EdgeLikeTo holds the string denoting the like_to edge name in mutations.
	EdgeLikeTo = "like_to"
	// EdgeSave holds the string denoting the save edge name in mutations.
	EdgeSave = "save"
	// EdgeGroup holds the string denoting the group edge name in mutations.
	EdgeGroup = "group"
	// EdgeChatroom holds the string denoting the chatroom edge name in mutations.
	EdgeChatroom = "chatroom"
	// EdgeMessage holds the string denoting the message edge name in mutations.
	EdgeMessage = "message"
	// EdgePics holds the string denoting the pics edge name in mutations.
	EdgePics = "pics"
	// Table holds the table name of the user in the database.
	Table = "users"
	// FriendsTable is the table that holds the friends relation/edge. The primary key declared below.
	FriendsTable = "user_friends"
	// RequestsTable is the table that holds the requests relation/edge. The primary key declared below.
	RequestsTable = "user_friends_req"
	// FriendsReqTable is the table that holds the friends_req relation/edge. The primary key declared below.
	FriendsReqTable = "user_friends_req"
	// LikeToTable is the table that holds the like_to relation/edge. The primary key declared below.
	LikeToTable = "user_like_to"
	// LikeToInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	LikeToInverseTable = "groups"
	// SaveTable is the table that holds the save relation/edge. The primary key declared below.
	SaveTable = "user_save"
	// SaveInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	SaveInverseTable = "groups"
	// GroupTable is the table that holds the group relation/edge. The primary key declared below.
	GroupTable = "group_members"
	// GroupInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupInverseTable = "groups"
	// ChatroomTable is the table that holds the chatroom relation/edge. The primary key declared below.
	ChatroomTable = "chat_room_participants"
	// ChatroomInverseTable is the table name for the ChatRoom entity.
	// It exists in this package in order to avoid circular dependency with the "chatroom" package.
	ChatroomInverseTable = "chat_rooms"
	// MessageTable is the table that holds the message relation/edge.
	MessageTable = "chat_messages"
	// MessageInverseTable is the table name for the ChatMessage entity.
	// It exists in this package in order to avoid circular dependency with the "chatmessage" package.
	MessageInverseTable = "chat_messages"
	// MessageColumn is the table column denoting the message relation/edge.
	MessageColumn = "user_id"
	// PicsTable is the table that holds the pics relation/edge.
	PicsTable = "pics"
	// PicsInverseTable is the table name for the Pic entity.
	// It exists in this package in order to avoid circular dependency with the "pic" package.
	PicsInverseTable = "pics"
	// PicsColumn is the table column denoting the pics relation/edge.
	PicsColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldActualName,
	FieldPhoneNumber,
	FieldUsername,
	FieldGender,
	FieldUni,
	FieldDep,
	FieldBio,
	FieldBirthYear,
	FieldIsVerified,
	FieldStatus,
	FieldMaxGroup,
	FieldAvatar,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldReadAt,
}

var (
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"user_id", "friend_id"}
	// RequestsPrimaryKey and RequestsColumn2 are the table columns denoting the
	// primary key for the requests relation (M2M).
	RequestsPrimaryKey = []string{"user_id", "request_id"}
	// FriendsReqPrimaryKey and FriendsReqColumn2 are the table columns denoting the
	// primary key for the friends_req relation (M2M).
	FriendsReqPrimaryKey = []string{"user_id", "request_id"}
	// LikeToPrimaryKey and LikeToColumn2 are the table columns denoting the
	// primary key for the like_to relation (M2M).
	LikeToPrimaryKey = []string{"user_id", "group_id"}
	// SavePrimaryKey and SaveColumn2 are the table columns denoting the
	// primary key for the save relation (M2M).
	SavePrimaryKey = []string{"user_id", "group_id"}
	// GroupPrimaryKey and GroupColumn2 are the table columns denoting the
	// primary key for the group relation (M2M).
	GroupPrimaryKey = []string{"group_id", "user_id"}
	// ChatroomPrimaryKey and ChatroomColumn2 are the table columns denoting the
	// primary key for the chatroom relation (M2M).
	ChatroomPrimaryKey = []string{"chat_room_id", "user_id"}
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
	// ActualNameValidator is a validator for the "actual_name" field. It is called by the builders before save.
	ActualNameValidator func(string) error
	// UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	UsernameValidator func(string) error
	// UniValidator is a validator for the "uni" field. It is called by the builders before save.
	UniValidator func(string) error
	// DepValidator is a validator for the "dep" field. It is called by the builders before save.
	DepValidator func(string) error
	// BirthYearValidator is a validator for the "birth_year" field. It is called by the builders before save.
	BirthYearValidator func(int) error
	// DefaultIsVerified holds the default value on creation for the "is_verified" field.
	DefaultIsVerified bool
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus string
	// DefaultMaxGroup holds the default value on creation for the "max_group" field.
	DefaultMaxGroup int
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultReadAt holds the default value on creation for the "read_at" field.
	DefaultReadAt func() time.Time
)
