// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lxxonx/cinder-server/ent/chatmessage"
	"github.com/lxxonx/cinder-server/ent/chatroom"
	"github.com/lxxonx/cinder-server/ent/user"
)

// ChatRoomCreate is the builder for creating a ChatRoom entity.
type ChatRoomCreate struct {
	config
	mutation *ChatRoomMutation
	hooks    []Hook
}

// SetCreatedAt sets the "createdAt" field.
func (crc *ChatRoomCreate) SetCreatedAt(t time.Time) *ChatRoomCreate {
	crc.mutation.SetCreatedAt(t)
	return crc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (crc *ChatRoomCreate) SetNillableCreatedAt(t *time.Time) *ChatRoomCreate {
	if t != nil {
		crc.SetCreatedAt(*t)
	}
	return crc
}

// SetUpdatedAt sets the "updatedAt" field.
func (crc *ChatRoomCreate) SetUpdatedAt(t time.Time) *ChatRoomCreate {
	crc.mutation.SetUpdatedAt(t)
	return crc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (crc *ChatRoomCreate) SetNillableUpdatedAt(t *time.Time) *ChatRoomCreate {
	if t != nil {
		crc.SetUpdatedAt(*t)
	}
	return crc
}

// SetReadAt sets the "readAt" field.
func (crc *ChatRoomCreate) SetReadAt(t time.Time) *ChatRoomCreate {
	crc.mutation.SetReadAt(t)
	return crc
}

// SetNillableReadAt sets the "readAt" field if the given value is not nil.
func (crc *ChatRoomCreate) SetNillableReadAt(t *time.Time) *ChatRoomCreate {
	if t != nil {
		crc.SetReadAt(*t)
	}
	return crc
}

// SetID sets the "id" field.
func (crc *ChatRoomCreate) SetID(s string) *ChatRoomCreate {
	crc.mutation.SetID(s)
	return crc
}

// AddParticipantIDs adds the "participants" edge to the User entity by IDs.
func (crc *ChatRoomCreate) AddParticipantIDs(ids ...string) *ChatRoomCreate {
	crc.mutation.AddParticipantIDs(ids...)
	return crc
}

// AddParticipants adds the "participants" edges to the User entity.
func (crc *ChatRoomCreate) AddParticipants(u ...*User) *ChatRoomCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return crc.AddParticipantIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the ChatMessage entity by IDs.
func (crc *ChatRoomCreate) AddMessageIDs(ids ...string) *ChatRoomCreate {
	crc.mutation.AddMessageIDs(ids...)
	return crc
}

// AddMessages adds the "messages" edges to the ChatMessage entity.
func (crc *ChatRoomCreate) AddMessages(c ...*ChatMessage) *ChatRoomCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return crc.AddMessageIDs(ids...)
}

// Mutation returns the ChatRoomMutation object of the builder.
func (crc *ChatRoomCreate) Mutation() *ChatRoomMutation {
	return crc.mutation
}

// Save creates the ChatRoom in the database.
func (crc *ChatRoomCreate) Save(ctx context.Context) (*ChatRoom, error) {
	var (
		err  error
		node *ChatRoom
	)
	crc.defaults()
	if len(crc.hooks) == 0 {
		if err = crc.check(); err != nil {
			return nil, err
		}
		node, err = crc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChatRoomMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = crc.check(); err != nil {
				return nil, err
			}
			crc.mutation = mutation
			if node, err = crc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(crc.hooks) - 1; i >= 0; i-- {
			if crc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = crc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, crc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (crc *ChatRoomCreate) SaveX(ctx context.Context) *ChatRoom {
	v, err := crc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (crc *ChatRoomCreate) Exec(ctx context.Context) error {
	_, err := crc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crc *ChatRoomCreate) ExecX(ctx context.Context) {
	if err := crc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (crc *ChatRoomCreate) defaults() {
	if _, ok := crc.mutation.CreatedAt(); !ok {
		v := chatroom.DefaultCreatedAt()
		crc.mutation.SetCreatedAt(v)
	}
	if _, ok := crc.mutation.UpdatedAt(); !ok {
		v := chatroom.DefaultUpdatedAt()
		crc.mutation.SetUpdatedAt(v)
	}
	if _, ok := crc.mutation.ReadAt(); !ok {
		v := chatroom.DefaultReadAt()
		crc.mutation.SetReadAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (crc *ChatRoomCreate) check() error {
	if _, ok := crc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "ChatRoom.createdAt"`)}
	}
	if _, ok := crc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "ChatRoom.updatedAt"`)}
	}
	if _, ok := crc.mutation.ReadAt(); !ok {
		return &ValidationError{Name: "readAt", err: errors.New(`ent: missing required field "ChatRoom.readAt"`)}
	}
	return nil
}

func (crc *ChatRoomCreate) sqlSave(ctx context.Context) (*ChatRoom, error) {
	_node, _spec := crc.createSpec()
	if err := sqlgraph.CreateNode(ctx, crc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected ChatRoom.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (crc *ChatRoomCreate) createSpec() (*ChatRoom, *sqlgraph.CreateSpec) {
	var (
		_node = &ChatRoom{config: crc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: chatroom.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: chatroom.FieldID,
			},
		}
	)
	if id, ok := crc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := crc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatroom.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := crc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatroom.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := crc.mutation.ReadAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: chatroom.FieldReadAt,
		})
		_node.ReadAt = value
	}
	if nodes := crc.mutation.ParticipantsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   chatroom.ParticipantsTable,
			Columns: chatroom.ParticipantsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := crc.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   chatroom.MessagesTable,
			Columns: []string{chatroom.MessagesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: chatmessage.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ChatRoomCreateBulk is the builder for creating many ChatRoom entities in bulk.
type ChatRoomCreateBulk struct {
	config
	builders []*ChatRoomCreate
}

// Save creates the ChatRoom entities in the database.
func (crcb *ChatRoomCreateBulk) Save(ctx context.Context) ([]*ChatRoom, error) {
	specs := make([]*sqlgraph.CreateSpec, len(crcb.builders))
	nodes := make([]*ChatRoom, len(crcb.builders))
	mutators := make([]Mutator, len(crcb.builders))
	for i := range crcb.builders {
		func(i int, root context.Context) {
			builder := crcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ChatRoomMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, crcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, crcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, crcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (crcb *ChatRoomCreateBulk) SaveX(ctx context.Context) []*ChatRoom {
	v, err := crcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (crcb *ChatRoomCreateBulk) Exec(ctx context.Context) error {
	_, err := crcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (crcb *ChatRoomCreateBulk) ExecX(ctx context.Context) {
	if err := crcb.Exec(ctx); err != nil {
		panic(err)
	}
}