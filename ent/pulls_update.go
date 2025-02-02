// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github-bot/ent/predicate"
	"github-bot/ent/pulls"
	"github-bot/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PullsUpdate is the builder for updating Pulls entities.
type PullsUpdate struct {
	config
	hooks    []Hook
	mutation *PullsMutation
}

// Where appends a list predicates to the PullsUpdate builder.
func (pu *PullsUpdate) Where(ps ...predicate.Pulls) *PullsUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetOwner sets the "owner" field.
func (pu *PullsUpdate) SetOwner(s string) *PullsUpdate {
	pu.mutation.SetOwner(s)
	return pu
}

// SetPrID sets the "prID" field.
func (pu *PullsUpdate) SetPrID(i int64) *PullsUpdate {
	pu.mutation.ResetPrID()
	pu.mutation.SetPrID(i)
	return pu
}

// SetNillablePrID sets the "prID" field if the given value is not nil.
func (pu *PullsUpdate) SetNillablePrID(i *int64) *PullsUpdate {
	if i != nil {
		pu.SetPrID(*i)
	}
	return pu
}

// AddPrID adds i to the "prID" field.
func (pu *PullsUpdate) AddPrID(i int64) *PullsUpdate {
	pu.mutation.AddPrID(i)
	return pu
}

// SetRepo sets the "repo" field.
func (pu *PullsUpdate) SetRepo(s string) *PullsUpdate {
	pu.mutation.SetRepo(s)
	return pu
}

// SetRepoID sets the "repoID" field.
func (pu *PullsUpdate) SetRepoID(i int64) *PullsUpdate {
	pu.mutation.ResetRepoID()
	pu.mutation.SetRepoID(i)
	return pu
}

// SetNillableRepoID sets the "repoID" field if the given value is not nil.
func (pu *PullsUpdate) SetNillableRepoID(i *int64) *PullsUpdate {
	if i != nil {
		pu.SetRepoID(*i)
	}
	return pu
}

// AddRepoID adds i to the "repoID" field.
func (pu *PullsUpdate) AddRepoID(i int64) *PullsUpdate {
	pu.mutation.AddRepoID(i)
	return pu
}

// SetNumber sets the "number" field.
func (pu *PullsUpdate) SetNumber(i int) *PullsUpdate {
	pu.mutation.ResetNumber()
	pu.mutation.SetNumber(i)
	return pu
}

// AddNumber adds i to the "number" field.
func (pu *PullsUpdate) AddNumber(i int) *PullsUpdate {
	pu.mutation.AddNumber(i)
	return pu
}

// SetComment sets the "comment" field.
func (pu *PullsUpdate) SetComment(i int64) *PullsUpdate {
	pu.mutation.ResetComment()
	pu.mutation.SetComment(i)
	return pu
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (pu *PullsUpdate) SetNillableComment(i *int64) *PullsUpdate {
	if i != nil {
		pu.SetComment(*i)
	}
	return pu
}

// AddComment adds i to the "comment" field.
func (pu *PullsUpdate) AddComment(i int64) *PullsUpdate {
	pu.mutation.AddComment(i)
	return pu
}

// SetCreatedAt sets the "createdAt" field.
func (pu *PullsUpdate) SetCreatedAt(t time.Time) *PullsUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetMergedAt sets the "mergedAt" field.
func (pu *PullsUpdate) SetMergedAt(t time.Time) *PullsUpdate {
	pu.mutation.SetMergedAt(t)
	return pu
}

// SetNillableMergedAt sets the "mergedAt" field if the given value is not nil.
func (pu *PullsUpdate) SetNillableMergedAt(t *time.Time) *PullsUpdate {
	if t != nil {
		pu.SetMergedAt(*t)
	}
	return pu
}

// ClearMergedAt clears the value of the "mergedAt" field.
func (pu *PullsUpdate) ClearMergedAt() *PullsUpdate {
	pu.mutation.ClearMergedAt()
	return pu
}

// SetCheckRunID sets the "checkRunID" field.
func (pu *PullsUpdate) SetCheckRunID(i int64) *PullsUpdate {
	pu.mutation.ResetCheckRunID()
	pu.mutation.SetCheckRunID(i)
	return pu
}

// SetNillableCheckRunID sets the "checkRunID" field if the given value is not nil.
func (pu *PullsUpdate) SetNillableCheckRunID(i *int64) *PullsUpdate {
	if i != nil {
		pu.SetCheckRunID(*i)
	}
	return pu
}

// AddCheckRunID adds i to the "checkRunID" field.
func (pu *PullsUpdate) AddCheckRunID(i int64) *PullsUpdate {
	pu.mutation.AddCheckRunID(i)
	return pu
}

// SetCheckRunResult sets the "checkRunResult" field.
func (pu *PullsUpdate) SetCheckRunResult(s string) *PullsUpdate {
	pu.mutation.SetCheckRunResult(s)
	return pu
}

// SetNillableCheckRunResult sets the "checkRunResult" field if the given value is not nil.
func (pu *PullsUpdate) SetNillableCheckRunResult(s *string) *PullsUpdate {
	if s != nil {
		pu.SetCheckRunResult(*s)
	}
	return pu
}

// SetHeadSha sets the "headSha" field.
func (pu *PullsUpdate) SetHeadSha(s string) *PullsUpdate {
	pu.mutation.SetHeadSha(s)
	return pu
}

// SetNillableHeadSha sets the "headSha" field if the given value is not nil.
func (pu *PullsUpdate) SetNillableHeadSha(s *string) *PullsUpdate {
	if s != nil {
		pu.SetHeadSha(*s)
	}
	return pu
}

// SetCreatorID sets the "Creator" edge to the User entity by ID.
func (pu *PullsUpdate) SetCreatorID(id int) *PullsUpdate {
	pu.mutation.SetCreatorID(id)
	return pu
}

// SetCreator sets the "Creator" edge to the User entity.
func (pu *PullsUpdate) SetCreator(u *User) *PullsUpdate {
	return pu.SetCreatorID(u.ID)
}

// Mutation returns the PullsMutation object of the builder.
func (pu *PullsUpdate) Mutation() *PullsMutation {
	return pu.mutation
}

// ClearCreator clears the "Creator" edge to the User entity.
func (pu *PullsUpdate) ClearCreator() *PullsUpdate {
	pu.mutation.ClearCreator()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PullsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PullsMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PullsUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PullsUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PullsUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PullsUpdate) check() error {
	if _, ok := pu.mutation.CreatorID(); pu.mutation.CreatorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Pulls.Creator"`)
	}
	return nil
}

func (pu *PullsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(pulls.Table, pulls.Columns, sqlgraph.NewFieldSpec(pulls.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Owner(); ok {
		_spec.SetField(pulls.FieldOwner, field.TypeString, value)
	}
	if value, ok := pu.mutation.PrID(); ok {
		_spec.SetField(pulls.FieldPrID, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedPrID(); ok {
		_spec.AddField(pulls.FieldPrID, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.Repo(); ok {
		_spec.SetField(pulls.FieldRepo, field.TypeString, value)
	}
	if value, ok := pu.mutation.RepoID(); ok {
		_spec.SetField(pulls.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedRepoID(); ok {
		_spec.AddField(pulls.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.Number(); ok {
		_spec.SetField(pulls.FieldNumber, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedNumber(); ok {
		_spec.AddField(pulls.FieldNumber, field.TypeInt, value)
	}
	if value, ok := pu.mutation.Comment(); ok {
		_spec.SetField(pulls.FieldComment, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedComment(); ok {
		_spec.AddField(pulls.FieldComment, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(pulls.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.MergedAt(); ok {
		_spec.SetField(pulls.FieldMergedAt, field.TypeTime, value)
	}
	if pu.mutation.MergedAtCleared() {
		_spec.ClearField(pulls.FieldMergedAt, field.TypeTime)
	}
	if value, ok := pu.mutation.CheckRunID(); ok {
		_spec.SetField(pulls.FieldCheckRunID, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedCheckRunID(); ok {
		_spec.AddField(pulls.FieldCheckRunID, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.CheckRunResult(); ok {
		_spec.SetField(pulls.FieldCheckRunResult, field.TypeString, value)
	}
	if value, ok := pu.mutation.HeadSha(); ok {
		_spec.SetField(pulls.FieldHeadSha, field.TypeString, value)
	}
	if pu.mutation.CreatorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pulls.CreatorTable,
			Columns: []string{pulls.CreatorColumn},
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
	if nodes := pu.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pulls.CreatorTable,
			Columns: []string{pulls.CreatorColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pulls.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PullsUpdateOne is the builder for updating a single Pulls entity.
type PullsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PullsMutation
}

// SetOwner sets the "owner" field.
func (puo *PullsUpdateOne) SetOwner(s string) *PullsUpdateOne {
	puo.mutation.SetOwner(s)
	return puo
}

// SetPrID sets the "prID" field.
func (puo *PullsUpdateOne) SetPrID(i int64) *PullsUpdateOne {
	puo.mutation.ResetPrID()
	puo.mutation.SetPrID(i)
	return puo
}

// SetNillablePrID sets the "prID" field if the given value is not nil.
func (puo *PullsUpdateOne) SetNillablePrID(i *int64) *PullsUpdateOne {
	if i != nil {
		puo.SetPrID(*i)
	}
	return puo
}

// AddPrID adds i to the "prID" field.
func (puo *PullsUpdateOne) AddPrID(i int64) *PullsUpdateOne {
	puo.mutation.AddPrID(i)
	return puo
}

// SetRepo sets the "repo" field.
func (puo *PullsUpdateOne) SetRepo(s string) *PullsUpdateOne {
	puo.mutation.SetRepo(s)
	return puo
}

// SetRepoID sets the "repoID" field.
func (puo *PullsUpdateOne) SetRepoID(i int64) *PullsUpdateOne {
	puo.mutation.ResetRepoID()
	puo.mutation.SetRepoID(i)
	return puo
}

// SetNillableRepoID sets the "repoID" field if the given value is not nil.
func (puo *PullsUpdateOne) SetNillableRepoID(i *int64) *PullsUpdateOne {
	if i != nil {
		puo.SetRepoID(*i)
	}
	return puo
}

// AddRepoID adds i to the "repoID" field.
func (puo *PullsUpdateOne) AddRepoID(i int64) *PullsUpdateOne {
	puo.mutation.AddRepoID(i)
	return puo
}

// SetNumber sets the "number" field.
func (puo *PullsUpdateOne) SetNumber(i int) *PullsUpdateOne {
	puo.mutation.ResetNumber()
	puo.mutation.SetNumber(i)
	return puo
}

// AddNumber adds i to the "number" field.
func (puo *PullsUpdateOne) AddNumber(i int) *PullsUpdateOne {
	puo.mutation.AddNumber(i)
	return puo
}

// SetComment sets the "comment" field.
func (puo *PullsUpdateOne) SetComment(i int64) *PullsUpdateOne {
	puo.mutation.ResetComment()
	puo.mutation.SetComment(i)
	return puo
}

// SetNillableComment sets the "comment" field if the given value is not nil.
func (puo *PullsUpdateOne) SetNillableComment(i *int64) *PullsUpdateOne {
	if i != nil {
		puo.SetComment(*i)
	}
	return puo
}

// AddComment adds i to the "comment" field.
func (puo *PullsUpdateOne) AddComment(i int64) *PullsUpdateOne {
	puo.mutation.AddComment(i)
	return puo
}

// SetCreatedAt sets the "createdAt" field.
func (puo *PullsUpdateOne) SetCreatedAt(t time.Time) *PullsUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetMergedAt sets the "mergedAt" field.
func (puo *PullsUpdateOne) SetMergedAt(t time.Time) *PullsUpdateOne {
	puo.mutation.SetMergedAt(t)
	return puo
}

// SetNillableMergedAt sets the "mergedAt" field if the given value is not nil.
func (puo *PullsUpdateOne) SetNillableMergedAt(t *time.Time) *PullsUpdateOne {
	if t != nil {
		puo.SetMergedAt(*t)
	}
	return puo
}

// ClearMergedAt clears the value of the "mergedAt" field.
func (puo *PullsUpdateOne) ClearMergedAt() *PullsUpdateOne {
	puo.mutation.ClearMergedAt()
	return puo
}

// SetCheckRunID sets the "checkRunID" field.
func (puo *PullsUpdateOne) SetCheckRunID(i int64) *PullsUpdateOne {
	puo.mutation.ResetCheckRunID()
	puo.mutation.SetCheckRunID(i)
	return puo
}

// SetNillableCheckRunID sets the "checkRunID" field if the given value is not nil.
func (puo *PullsUpdateOne) SetNillableCheckRunID(i *int64) *PullsUpdateOne {
	if i != nil {
		puo.SetCheckRunID(*i)
	}
	return puo
}

// AddCheckRunID adds i to the "checkRunID" field.
func (puo *PullsUpdateOne) AddCheckRunID(i int64) *PullsUpdateOne {
	puo.mutation.AddCheckRunID(i)
	return puo
}

// SetCheckRunResult sets the "checkRunResult" field.
func (puo *PullsUpdateOne) SetCheckRunResult(s string) *PullsUpdateOne {
	puo.mutation.SetCheckRunResult(s)
	return puo
}

// SetNillableCheckRunResult sets the "checkRunResult" field if the given value is not nil.
func (puo *PullsUpdateOne) SetNillableCheckRunResult(s *string) *PullsUpdateOne {
	if s != nil {
		puo.SetCheckRunResult(*s)
	}
	return puo
}

// SetHeadSha sets the "headSha" field.
func (puo *PullsUpdateOne) SetHeadSha(s string) *PullsUpdateOne {
	puo.mutation.SetHeadSha(s)
	return puo
}

// SetNillableHeadSha sets the "headSha" field if the given value is not nil.
func (puo *PullsUpdateOne) SetNillableHeadSha(s *string) *PullsUpdateOne {
	if s != nil {
		puo.SetHeadSha(*s)
	}
	return puo
}

// SetCreatorID sets the "Creator" edge to the User entity by ID.
func (puo *PullsUpdateOne) SetCreatorID(id int) *PullsUpdateOne {
	puo.mutation.SetCreatorID(id)
	return puo
}

// SetCreator sets the "Creator" edge to the User entity.
func (puo *PullsUpdateOne) SetCreator(u *User) *PullsUpdateOne {
	return puo.SetCreatorID(u.ID)
}

// Mutation returns the PullsMutation object of the builder.
func (puo *PullsUpdateOne) Mutation() *PullsMutation {
	return puo.mutation
}

// ClearCreator clears the "Creator" edge to the User entity.
func (puo *PullsUpdateOne) ClearCreator() *PullsUpdateOne {
	puo.mutation.ClearCreator()
	return puo
}

// Where appends a list predicates to the PullsUpdate builder.
func (puo *PullsUpdateOne) Where(ps ...predicate.Pulls) *PullsUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PullsUpdateOne) Select(field string, fields ...string) *PullsUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pulls entity.
func (puo *PullsUpdateOne) Save(ctx context.Context) (*Pulls, error) {
	return withHooks[*Pulls, PullsMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PullsUpdateOne) SaveX(ctx context.Context) *Pulls {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PullsUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PullsUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PullsUpdateOne) check() error {
	if _, ok := puo.mutation.CreatorID(); puo.mutation.CreatorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Pulls.Creator"`)
	}
	return nil
}

func (puo *PullsUpdateOne) sqlSave(ctx context.Context) (_node *Pulls, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(pulls.Table, pulls.Columns, sqlgraph.NewFieldSpec(pulls.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Pulls.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pulls.FieldID)
		for _, f := range fields {
			if !pulls.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pulls.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Owner(); ok {
		_spec.SetField(pulls.FieldOwner, field.TypeString, value)
	}
	if value, ok := puo.mutation.PrID(); ok {
		_spec.SetField(pulls.FieldPrID, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedPrID(); ok {
		_spec.AddField(pulls.FieldPrID, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.Repo(); ok {
		_spec.SetField(pulls.FieldRepo, field.TypeString, value)
	}
	if value, ok := puo.mutation.RepoID(); ok {
		_spec.SetField(pulls.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedRepoID(); ok {
		_spec.AddField(pulls.FieldRepoID, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.Number(); ok {
		_spec.SetField(pulls.FieldNumber, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedNumber(); ok {
		_spec.AddField(pulls.FieldNumber, field.TypeInt, value)
	}
	if value, ok := puo.mutation.Comment(); ok {
		_spec.SetField(pulls.FieldComment, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedComment(); ok {
		_spec.AddField(pulls.FieldComment, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(pulls.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.MergedAt(); ok {
		_spec.SetField(pulls.FieldMergedAt, field.TypeTime, value)
	}
	if puo.mutation.MergedAtCleared() {
		_spec.ClearField(pulls.FieldMergedAt, field.TypeTime)
	}
	if value, ok := puo.mutation.CheckRunID(); ok {
		_spec.SetField(pulls.FieldCheckRunID, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedCheckRunID(); ok {
		_spec.AddField(pulls.FieldCheckRunID, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.CheckRunResult(); ok {
		_spec.SetField(pulls.FieldCheckRunResult, field.TypeString, value)
	}
	if value, ok := puo.mutation.HeadSha(); ok {
		_spec.SetField(pulls.FieldHeadSha, field.TypeString, value)
	}
	if puo.mutation.CreatorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pulls.CreatorTable,
			Columns: []string{pulls.CreatorColumn},
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
	if nodes := puo.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pulls.CreatorTable,
			Columns: []string{pulls.CreatorColumn},
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
	_node = &Pulls{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pulls.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
