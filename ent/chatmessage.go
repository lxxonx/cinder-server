// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/lxxonx/cinder-server/ent/chatmessage"
	"github.com/lxxonx/cinder-server/ent/chatroom"
	"github.com/lxxonx/cinder-server/ent/user"
)

// ChatMessage is the model entity for the ChatMessage schema.
type ChatMessage struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Message holds the value of the "message" field.
	Message string `json:"message,omitempty"`
	// RoomID holds the value of the "room_id" field.
	RoomID uuid.UUID `json:"room_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// ReadAt holds the value of the "read_at" field.
	ReadAt time.Time `json:"read_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ChatMessageQuery when eager-loading is set.
	Edges ChatMessageEdges `json:"edges"`
}

// ChatMessageEdges holds the relations/edges for other nodes in the graph.
type ChatMessageEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Room holds the value of the room edge.
	Room *ChatRoom `json:"room,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChatMessageEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// RoomOrErr returns the Room value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ChatMessageEdges) RoomOrErr() (*ChatRoom, error) {
	if e.loadedTypes[1] {
		if e.Room == nil {
			// The edge room was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: chatroom.Label}
		}
		return e.Room, nil
	}
	return nil, &NotLoadedError{edge: "room"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*ChatMessage) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case chatmessage.FieldMessage, chatmessage.FieldUserID:
			values[i] = new(sql.NullString)
		case chatmessage.FieldCreatedAt, chatmessage.FieldUpdatedAt, chatmessage.FieldReadAt:
			values[i] = new(sql.NullTime)
		case chatmessage.FieldID, chatmessage.FieldRoomID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type ChatMessage", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the ChatMessage fields.
func (cm *ChatMessage) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case chatmessage.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				cm.ID = *value
			}
		case chatmessage.FieldMessage:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field message", values[i])
			} else if value.Valid {
				cm.Message = value.String
			}
		case chatmessage.FieldRoomID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field room_id", values[i])
			} else if value != nil {
				cm.RoomID = *value
			}
		case chatmessage.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				cm.UserID = value.String
			}
		case chatmessage.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				cm.CreatedAt = value.Time
			}
		case chatmessage.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				cm.UpdatedAt = value.Time
			}
		case chatmessage.FieldReadAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field read_at", values[i])
			} else if value.Valid {
				cm.ReadAt = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the ChatMessage entity.
func (cm *ChatMessage) QueryUser() *UserQuery {
	return (&ChatMessageClient{config: cm.config}).QueryUser(cm)
}

// QueryRoom queries the "room" edge of the ChatMessage entity.
func (cm *ChatMessage) QueryRoom() *ChatRoomQuery {
	return (&ChatMessageClient{config: cm.config}).QueryRoom(cm)
}

// Update returns a builder for updating this ChatMessage.
// Note that you need to call ChatMessage.Unwrap() before calling this method if this ChatMessage
// was returned from a transaction, and the transaction was committed or rolled back.
func (cm *ChatMessage) Update() *ChatMessageUpdateOne {
	return (&ChatMessageClient{config: cm.config}).UpdateOne(cm)
}

// Unwrap unwraps the ChatMessage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cm *ChatMessage) Unwrap() *ChatMessage {
	tx, ok := cm.config.driver.(*txDriver)
	if !ok {
		panic("ent: ChatMessage is not a transactional entity")
	}
	cm.config.driver = tx.drv
	return cm
}

// String implements the fmt.Stringer.
func (cm *ChatMessage) String() string {
	var builder strings.Builder
	builder.WriteString("ChatMessage(")
	builder.WriteString(fmt.Sprintf("id=%v", cm.ID))
	builder.WriteString(", message=")
	builder.WriteString(cm.Message)
	builder.WriteString(", room_id=")
	builder.WriteString(fmt.Sprintf("%v", cm.RoomID))
	builder.WriteString(", user_id=")
	builder.WriteString(cm.UserID)
	builder.WriteString(", created_at=")
	builder.WriteString(cm.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(cm.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", read_at=")
	builder.WriteString(cm.ReadAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// ChatMessages is a parsable slice of ChatMessage.
type ChatMessages []*ChatMessage

func (cm ChatMessages) config(cfg config) {
	for _i := range cm {
		cm[_i].config = cfg
	}
}
