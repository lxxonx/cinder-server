// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lxxonx/cinder-server/ent/chatmessage"
	"github.com/lxxonx/cinder-server/ent/chatroom"
	"github.com/lxxonx/cinder-server/ent/predicate"
	"github.com/lxxonx/cinder-server/ent/user"
)

// ChatMessageUpdate is the builder for updating ChatMessage entities.
type ChatMessageUpdate struct {
	config
	hooks    []Hook
	mutation *ChatMessageMutation
}

// Where appends a list predicates to the ChatMessageUpdate builder.
func (cmu *ChatMessageUpdate) Where(ps ...predicate.ChatMessage) *ChatMessageUpdate {
	cmu.mutation.Where(ps...)
	return cmu
}

// SetUID sets the "uid" field.
func (cmu *ChatMessageUpdate) SetUID(s string) *ChatMessageUpdate {
	cmu.mutation.SetUID(s)
	return cmu
}

// SetMessage sets the "message" field.
func (cmu *ChatMessageUpdate) SetMessage(s string) *ChatMessageUpdate {
	cmu.mutation.SetMessage(s)
	return cmu
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (cmu *ChatMessageUpdate) SetNillableMessage(s *string) *ChatMessageUpdate {
	if s != nil {
		cmu.SetMessage(*s)
	}
	return cmu
}

// SetRoomID sets the "room_id" field.
func (cmu *ChatMessageUpdate) SetRoomID(i int) *ChatMessageUpdate {
	cmu.mutation.SetRoomID(i)
	return cmu
}

// SetNillableRoomID sets the "room_id" field if the given value is not nil.
func (cmu *ChatMessageUpdate) SetNillableRoomID(i *int) *ChatMessageUpdate {
	if i != nil {
		cmu.SetRoomID(*i)
	}
	return cmu
}

// ClearRoomID clears the value of the "room_id" field.
func (cmu *ChatMessageUpdate) ClearRoomID() *ChatMessageUpdate {
	cmu.mutation.ClearRoomID()
	return cmu
}

// SetUserID sets the "user_id" field.
func (cmu *ChatMessageUpdate) SetUserID(i int) *ChatMessageUpdate {
	cmu.mutation.SetUserID(i)
	return cmu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cmu *ChatMessageUpdate) SetNillableUserID(i *int) *ChatMessageUpdate {
	if i != nil {
		cmu.SetUserID(*i)
	}
	return cmu
}

// ClearUserID clears the value of the "user_id" field.
func (cmu *ChatMessageUpdate) ClearUserID() *ChatMessageUpdate {
	cmu.mutation.ClearUserID()
	return cmu
}

// SetCreatedAt sets the "createdAt" field.
func (cmu *ChatMessageUpdate) SetCreatedAt(t time.Time) *ChatMessageUpdate {
	cmu.mutation.SetCreatedAt(t)
	return cmu
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (cmu *ChatMessageUpdate) SetNillableCreatedAt(t *time.Time) *ChatMessageUpdate {
	if t != nil {
		cmu.SetCreatedAt(*t)
	}
	return cmu
}

// SetReadAt sets the "readAt" field.
func (cmu *ChatMessageUpdate) SetReadAt(t time.Time) *ChatMessageUpdate {
	cmu.mutation.SetReadAt(t)
	return cmu
}

// SetNillableReadAt sets the "readAt" field if the given value is not nil.
func (cmu *ChatMessageUpdate) SetNillableReadAt(t *time.Time) *ChatMessageUpdate {
	if t != nil {
		cmu.SetReadAt(*t)
	}
	return cmu
}

// SetUser sets the "user" edge to the User entity.
func (cmu *ChatMessageUpdate) SetUser(u *User) *ChatMessageUpdate {
	return cmu.SetUserID(u.ID)
}

// SetRoom sets the "room" edge to the ChatRoom entity.
func (cmu *ChatMessageUpdate) SetRoom(c *ChatRoom) *ChatMessageUpdate {
	return cmu.SetRoomID(c.ID)
}

// Mutation returns the ChatMessageMutation object of the builder.
func (cmu *ChatMessageUpdate) Mutation() *ChatMessageMutation {
	return cmu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cmu *ChatMessageUpdate) ClearUser() *ChatMessageUpdate {
	cmu.mutation.ClearUser()
	return cmu
}

// ClearRoom clears the "room" edge to the ChatRoom entity.
func (cmu *ChatMessageUpdate) ClearRoom() *ChatMessageUpdate {
	cmu.mutation.ClearRoom()
	return cmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cmu *ChatMessageUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cmu.hooks) == 0 {
		affected, err = cmu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChatMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cmu.mutation = mutation
			affected, err = cmu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cmu.hooks) - 1; i >= 0; i-- {
			if cmu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cmu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cmu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cmu *ChatMessageUpdate) SaveX(ctx context.Context) int {
	affected, err := cmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cmu *ChatMessageUpdate) Exec(ctx context.Context) error {
	_, err := cmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmu *ChatMessageUpdate) ExecX(ctx context.Context) {
	if err := cmu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cmu *ChatMessageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chatmessage.Table,
			Columns: chatmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: chatmessage.FieldID,
			},
		},
	}
	if ps := cmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cmu.mutation.UID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatmessage.FieldUID,
		})
	}
	if value, ok := cmu.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatmessage.FieldMessage,
		})
	}
	if value, ok := cmu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatmessage.FieldCreatedAt,
		})
	}
	if value, ok := cmu.mutation.ReadAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatmessage.FieldReadAt,
		})
	}
	if cmu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatmessage.UserTable,
			Columns: []string{chatmessage.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatmessage.UserTable,
			Columns: []string{chatmessage.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cmu.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatmessage.RoomTable,
			Columns: []string{chatmessage.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: chatroom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmu.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatmessage.RoomTable,
			Columns: []string{chatmessage.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: chatroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chatmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ChatMessageUpdateOne is the builder for updating a single ChatMessage entity.
type ChatMessageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChatMessageMutation
}

// SetUID sets the "uid" field.
func (cmuo *ChatMessageUpdateOne) SetUID(s string) *ChatMessageUpdateOne {
	cmuo.mutation.SetUID(s)
	return cmuo
}

// SetMessage sets the "message" field.
func (cmuo *ChatMessageUpdateOne) SetMessage(s string) *ChatMessageUpdateOne {
	cmuo.mutation.SetMessage(s)
	return cmuo
}

// SetNillableMessage sets the "message" field if the given value is not nil.
func (cmuo *ChatMessageUpdateOne) SetNillableMessage(s *string) *ChatMessageUpdateOne {
	if s != nil {
		cmuo.SetMessage(*s)
	}
	return cmuo
}

// SetRoomID sets the "room_id" field.
func (cmuo *ChatMessageUpdateOne) SetRoomID(i int) *ChatMessageUpdateOne {
	cmuo.mutation.SetRoomID(i)
	return cmuo
}

// SetNillableRoomID sets the "room_id" field if the given value is not nil.
func (cmuo *ChatMessageUpdateOne) SetNillableRoomID(i *int) *ChatMessageUpdateOne {
	if i != nil {
		cmuo.SetRoomID(*i)
	}
	return cmuo
}

// ClearRoomID clears the value of the "room_id" field.
func (cmuo *ChatMessageUpdateOne) ClearRoomID() *ChatMessageUpdateOne {
	cmuo.mutation.ClearRoomID()
	return cmuo
}

// SetUserID sets the "user_id" field.
func (cmuo *ChatMessageUpdateOne) SetUserID(i int) *ChatMessageUpdateOne {
	cmuo.mutation.SetUserID(i)
	return cmuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (cmuo *ChatMessageUpdateOne) SetNillableUserID(i *int) *ChatMessageUpdateOne {
	if i != nil {
		cmuo.SetUserID(*i)
	}
	return cmuo
}

// ClearUserID clears the value of the "user_id" field.
func (cmuo *ChatMessageUpdateOne) ClearUserID() *ChatMessageUpdateOne {
	cmuo.mutation.ClearUserID()
	return cmuo
}

// SetCreatedAt sets the "createdAt" field.
func (cmuo *ChatMessageUpdateOne) SetCreatedAt(t time.Time) *ChatMessageUpdateOne {
	cmuo.mutation.SetCreatedAt(t)
	return cmuo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (cmuo *ChatMessageUpdateOne) SetNillableCreatedAt(t *time.Time) *ChatMessageUpdateOne {
	if t != nil {
		cmuo.SetCreatedAt(*t)
	}
	return cmuo
}

// SetReadAt sets the "readAt" field.
func (cmuo *ChatMessageUpdateOne) SetReadAt(t time.Time) *ChatMessageUpdateOne {
	cmuo.mutation.SetReadAt(t)
	return cmuo
}

// SetNillableReadAt sets the "readAt" field if the given value is not nil.
func (cmuo *ChatMessageUpdateOne) SetNillableReadAt(t *time.Time) *ChatMessageUpdateOne {
	if t != nil {
		cmuo.SetReadAt(*t)
	}
	return cmuo
}

// SetUser sets the "user" edge to the User entity.
func (cmuo *ChatMessageUpdateOne) SetUser(u *User) *ChatMessageUpdateOne {
	return cmuo.SetUserID(u.ID)
}

// SetRoom sets the "room" edge to the ChatRoom entity.
func (cmuo *ChatMessageUpdateOne) SetRoom(c *ChatRoom) *ChatMessageUpdateOne {
	return cmuo.SetRoomID(c.ID)
}

// Mutation returns the ChatMessageMutation object of the builder.
func (cmuo *ChatMessageUpdateOne) Mutation() *ChatMessageMutation {
	return cmuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (cmuo *ChatMessageUpdateOne) ClearUser() *ChatMessageUpdateOne {
	cmuo.mutation.ClearUser()
	return cmuo
}

// ClearRoom clears the "room" edge to the ChatRoom entity.
func (cmuo *ChatMessageUpdateOne) ClearRoom() *ChatMessageUpdateOne {
	cmuo.mutation.ClearRoom()
	return cmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cmuo *ChatMessageUpdateOne) Select(field string, fields ...string) *ChatMessageUpdateOne {
	cmuo.fields = append([]string{field}, fields...)
	return cmuo
}

// Save executes the query and returns the updated ChatMessage entity.
func (cmuo *ChatMessageUpdateOne) Save(ctx context.Context) (*ChatMessage, error) {
	var (
		err  error
		node *ChatMessage
	)
	if len(cmuo.hooks) == 0 {
		node, err = cmuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChatMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cmuo.mutation = mutation
			node, err = cmuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cmuo.hooks) - 1; i >= 0; i-- {
			if cmuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cmuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cmuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cmuo *ChatMessageUpdateOne) SaveX(ctx context.Context) *ChatMessage {
	node, err := cmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cmuo *ChatMessageUpdateOne) Exec(ctx context.Context) error {
	_, err := cmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmuo *ChatMessageUpdateOne) ExecX(ctx context.Context) {
	if err := cmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cmuo *ChatMessageUpdateOne) sqlSave(ctx context.Context) (_node *ChatMessage, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chatmessage.Table,
			Columns: chatmessage.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: chatmessage.FieldID,
			},
		},
	}
	id, ok := cmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ChatMessage.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chatmessage.FieldID)
		for _, f := range fields {
			if !chatmessage.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chatmessage.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cmuo.mutation.UID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatmessage.FieldUID,
		})
	}
	if value, ok := cmuo.mutation.Message(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chatmessage.FieldMessage,
		})
	}
	if value, ok := cmuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatmessage.FieldCreatedAt,
		})
	}
	if value, ok := cmuo.mutation.ReadAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatmessage.FieldReadAt,
		})
	}
	if cmuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatmessage.UserTable,
			Columns: []string{chatmessage.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatmessage.UserTable,
			Columns: []string{chatmessage.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cmuo.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatmessage.RoomTable,
			Columns: []string{chatmessage.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: chatroom.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cmuo.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   chatmessage.RoomTable,
			Columns: []string{chatmessage.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: chatroom.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ChatMessage{config: cmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chatmessage.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
