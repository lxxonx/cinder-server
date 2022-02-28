// Code generated by entc, DO NOT EDIT.

package pic

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/lxxonx/cinder-server/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupID), v))
	})
}

// Adress applies equality check predicate on the "adress" field. It's identical to AdressEQ.
func Adress(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdress), v))
	})
}

// CreatedAt applies equality check predicate on the "createdAt" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updatedAt" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// ReadAt applies equality check predicate on the "readAt" field. It's identical to ReadAtEQ.
func ReadAt(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReadAt), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Pic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Pic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUserID), v))
	})
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUserID), v))
	})
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUserID), v))
	})
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUserID), v))
	})
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldUserID), v))
	})
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldUserID), v))
	})
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldUserID), v))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldUserID), v))
	})
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldUserID), v))
	})
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupID), v))
	})
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGroupID), v))
	})
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...string) predicate.Pic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGroupID), v...))
	})
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...string) predicate.Pic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGroupID), v...))
	})
}

// GroupIDGT applies the GT predicate on the "group_id" field.
func GroupIDGT(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGroupID), v))
	})
}

// GroupIDGTE applies the GTE predicate on the "group_id" field.
func GroupIDGTE(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGroupID), v))
	})
}

// GroupIDLT applies the LT predicate on the "group_id" field.
func GroupIDLT(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGroupID), v))
	})
}

// GroupIDLTE applies the LTE predicate on the "group_id" field.
func GroupIDLTE(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGroupID), v))
	})
}

// GroupIDContains applies the Contains predicate on the "group_id" field.
func GroupIDContains(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGroupID), v))
	})
}

// GroupIDHasPrefix applies the HasPrefix predicate on the "group_id" field.
func GroupIDHasPrefix(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGroupID), v))
	})
}

// GroupIDHasSuffix applies the HasSuffix predicate on the "group_id" field.
func GroupIDHasSuffix(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGroupID), v))
	})
}

// GroupIDIsNil applies the IsNil predicate on the "group_id" field.
func GroupIDIsNil() predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldGroupID)))
	})
}

// GroupIDNotNil applies the NotNil predicate on the "group_id" field.
func GroupIDNotNil() predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldGroupID)))
	})
}

// GroupIDEqualFold applies the EqualFold predicate on the "group_id" field.
func GroupIDEqualFold(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGroupID), v))
	})
}

// GroupIDContainsFold applies the ContainsFold predicate on the "group_id" field.
func GroupIDContainsFold(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGroupID), v))
	})
}

// AdressEQ applies the EQ predicate on the "adress" field.
func AdressEQ(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAdress), v))
	})
}

// AdressNEQ applies the NEQ predicate on the "adress" field.
func AdressNEQ(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAdress), v))
	})
}

// AdressIn applies the In predicate on the "adress" field.
func AdressIn(vs ...string) predicate.Pic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAdress), v...))
	})
}

// AdressNotIn applies the NotIn predicate on the "adress" field.
func AdressNotIn(vs ...string) predicate.Pic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAdress), v...))
	})
}

// AdressGT applies the GT predicate on the "adress" field.
func AdressGT(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAdress), v))
	})
}

// AdressGTE applies the GTE predicate on the "adress" field.
func AdressGTE(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAdress), v))
	})
}

// AdressLT applies the LT predicate on the "adress" field.
func AdressLT(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAdress), v))
	})
}

// AdressLTE applies the LTE predicate on the "adress" field.
func AdressLTE(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAdress), v))
	})
}

// AdressContains applies the Contains predicate on the "adress" field.
func AdressContains(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAdress), v))
	})
}

// AdressHasPrefix applies the HasPrefix predicate on the "adress" field.
func AdressHasPrefix(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAdress), v))
	})
}

// AdressHasSuffix applies the HasSuffix predicate on the "adress" field.
func AdressHasSuffix(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAdress), v))
	})
}

// AdressEqualFold applies the EqualFold predicate on the "adress" field.
func AdressEqualFold(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAdress), v))
	})
}

// AdressContainsFold applies the ContainsFold predicate on the "adress" field.
func AdressContainsFold(v string) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAdress), v))
	})
}

// CreatedAtEQ applies the EQ predicate on the "createdAt" field.
func CreatedAtEQ(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "createdAt" field.
func CreatedAtNEQ(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "createdAt" field.
func CreatedAtIn(vs ...time.Time) predicate.Pic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "createdAt" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Pic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "createdAt" field.
func CreatedAtGT(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "createdAt" field.
func CreatedAtGTE(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "createdAt" field.
func CreatedAtLT(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "createdAt" field.
func CreatedAtLTE(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updatedAt" field.
func UpdatedAtEQ(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updatedAt" field.
func UpdatedAtNEQ(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updatedAt" field.
func UpdatedAtIn(vs ...time.Time) predicate.Pic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updatedAt" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Pic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updatedAt" field.
func UpdatedAtGT(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updatedAt" field.
func UpdatedAtGTE(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updatedAt" field.
func UpdatedAtLT(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updatedAt" field.
func UpdatedAtLTE(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// ReadAtEQ applies the EQ predicate on the "readAt" field.
func ReadAtEQ(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldReadAt), v))
	})
}

// ReadAtNEQ applies the NEQ predicate on the "readAt" field.
func ReadAtNEQ(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldReadAt), v))
	})
}

// ReadAtIn applies the In predicate on the "readAt" field.
func ReadAtIn(vs ...time.Time) predicate.Pic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldReadAt), v...))
	})
}

// ReadAtNotIn applies the NotIn predicate on the "readAt" field.
func ReadAtNotIn(vs ...time.Time) predicate.Pic {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pic(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldReadAt), v...))
	})
}

// ReadAtGT applies the GT predicate on the "readAt" field.
func ReadAtGT(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldReadAt), v))
	})
}

// ReadAtGTE applies the GTE predicate on the "readAt" field.
func ReadAtGTE(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldReadAt), v))
	})
}

// ReadAtLT applies the LT predicate on the "readAt" field.
func ReadAtLT(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldReadAt), v))
	})
}

// ReadAtLTE applies the LTE predicate on the "readAt" field.
func ReadAtLTE(v time.Time) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldReadAt), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroup applies the HasEdge predicate on the "group" edge.
func HasGroup() predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupWith applies the HasEdge predicate on the "group" edge with a given conditions (other predicates).
func HasGroupWith(preds ...predicate.Group) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Pic) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Pic) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Pic) predicate.Pic {
	return predicate.Pic(func(s *sql.Selector) {
		p(s.Not())
	})
}