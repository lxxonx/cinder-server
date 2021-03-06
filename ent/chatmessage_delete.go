// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/lxxonx/cinder-server/ent/chatmessage"
	"github.com/lxxonx/cinder-server/ent/predicate"
)

// ChatMessageDelete is the builder for deleting a ChatMessage entity.
type ChatMessageDelete struct {
	config
	hooks    []Hook
	mutation *ChatMessageMutation
}

// Where appends a list predicates to the ChatMessageDelete builder.
func (cmd *ChatMessageDelete) Where(ps ...predicate.ChatMessage) *ChatMessageDelete {
	cmd.mutation.Where(ps...)
	return cmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cmd *ChatMessageDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cmd.hooks) == 0 {
		affected, err = cmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChatMessageMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cmd.mutation = mutation
			affected, err = cmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cmd.hooks) - 1; i >= 0; i-- {
			if cmd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cmd *ChatMessageDelete) ExecX(ctx context.Context) int {
	n, err := cmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cmd *ChatMessageDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: chatmessage.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: chatmessage.FieldID,
			},
		},
	}
	if ps := cmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cmd.driver, _spec)
}

// ChatMessageDeleteOne is the builder for deleting a single ChatMessage entity.
type ChatMessageDeleteOne struct {
	cmd *ChatMessageDelete
}

// Exec executes the deletion query.
func (cmdo *ChatMessageDeleteOne) Exec(ctx context.Context) error {
	n, err := cmdo.cmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{chatmessage.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cmdo *ChatMessageDeleteOne) ExecX(ctx context.Context) {
	cmdo.cmd.ExecX(ctx)
}
