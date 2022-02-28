// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/lxxonx/cinder-server/ent/group"
	"github.com/lxxonx/cinder-server/ent/pic"
	"github.com/lxxonx/cinder-server/ent/user"
)

// Pic is the model entity for the Pic schema.
type Pic struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID string `json:"group_id,omitempty"`
	// Adress holds the value of the "adress" field.
	Adress string `json:"adress,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// ReadAt holds the value of the "readAt" field.
	ReadAt time.Time `json:"readAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PicQuery when eager-loading is set.
	Edges PicEdges `json:"edges"`
}

// PicEdges holds the relations/edges for other nodes in the graph.
type PicEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PicEdges) UserOrErr() (*User, error) {
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

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PicEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[1] {
		if e.Group == nil {
			// The edge group was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Pic) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case pic.FieldID:
			values[i] = new(sql.NullInt64)
		case pic.FieldUserID, pic.FieldGroupID, pic.FieldAdress:
			values[i] = new(sql.NullString)
		case pic.FieldCreatedAt, pic.FieldUpdatedAt, pic.FieldReadAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Pic", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Pic fields.
func (pi *Pic) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case pic.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pi.ID = int(value.Int64)
		case pic.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				pi.UserID = value.String
			}
		case pic.FieldGroupID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				pi.GroupID = value.String
			}
		case pic.FieldAdress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field adress", values[i])
			} else if value.Valid {
				pi.Adress = value.String
			}
		case pic.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				pi.CreatedAt = value.Time
			}
		case pic.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				pi.UpdatedAt = value.Time
			}
		case pic.FieldReadAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field readAt", values[i])
			} else if value.Valid {
				pi.ReadAt = value.Time
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Pic entity.
func (pi *Pic) QueryUser() *UserQuery {
	return (&PicClient{config: pi.config}).QueryUser(pi)
}

// QueryGroup queries the "group" edge of the Pic entity.
func (pi *Pic) QueryGroup() *GroupQuery {
	return (&PicClient{config: pi.config}).QueryGroup(pi)
}

// Update returns a builder for updating this Pic.
// Note that you need to call Pic.Unwrap() before calling this method if this Pic
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *Pic) Update() *PicUpdateOne {
	return (&PicClient{config: pi.config}).UpdateOne(pi)
}

// Unwrap unwraps the Pic entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pi *Pic) Unwrap() *Pic {
	tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("ent: Pic is not a transactional entity")
	}
	pi.config.driver = tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *Pic) String() string {
	var builder strings.Builder
	builder.WriteString("Pic(")
	builder.WriteString(fmt.Sprintf("id=%v", pi.ID))
	builder.WriteString(", user_id=")
	builder.WriteString(pi.UserID)
	builder.WriteString(", group_id=")
	builder.WriteString(pi.GroupID)
	builder.WriteString(", adress=")
	builder.WriteString(pi.Adress)
	builder.WriteString(", createdAt=")
	builder.WriteString(pi.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", updatedAt=")
	builder.WriteString(pi.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", readAt=")
	builder.WriteString(pi.ReadAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Pics is a parsable slice of Pic.
type Pics []*Pic

func (pi Pics) config(cfg config) {
	for _i := range pi {
		pi[_i].config = cfg
	}
}
