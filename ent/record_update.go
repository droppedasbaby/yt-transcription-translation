// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/GrewalAS/yt-transcription-translation/ent/predicate"
	"github.com/GrewalAS/yt-transcription-translation/ent/record"
)

// RecordUpdate is the builder for updating Record entities.
type RecordUpdate struct {
	config
	hooks    []Hook
	mutation *RecordMutation
}

// Where appends a list predicates to the RecordUpdate builder.
func (ru *RecordUpdate) Where(ps ...predicate.Record) *RecordUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetVideoURL sets the "video_url" field.
func (ru *RecordUpdate) SetVideoURL(s string) *RecordUpdate {
	ru.mutation.SetVideoURL(s)
	return ru
}

// SetVideoID sets the "video_id" field.
func (ru *RecordUpdate) SetVideoID(s string) *RecordUpdate {
	ru.mutation.SetVideoID(s)
	return ru
}

// SetFileLocation sets the "file_location" field.
func (ru *RecordUpdate) SetFileLocation(s string) *RecordUpdate {
	ru.mutation.SetFileLocation(s)
	return ru
}

// SetNillableFileLocation sets the "file_location" field if the given value is not nil.
func (ru *RecordUpdate) SetNillableFileLocation(s *string) *RecordUpdate {
	if s != nil {
		ru.SetFileLocation(*s)
	}
	return ru
}

// ClearFileLocation clears the value of the "file_location" field.
func (ru *RecordUpdate) ClearFileLocation() *RecordUpdate {
	ru.mutation.ClearFileLocation()
	return ru
}

// SetStatus sets the "status" field.
func (ru *RecordUpdate) SetStatus(r record.Status) *RecordUpdate {
	ru.mutation.SetStatus(r)
	return ru
}

// Mutation returns the RecordMutation object of the builder.
func (ru *RecordUpdate) Mutation() *RecordMutation {
	return ru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RecordUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RecordUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RecordUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RecordUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RecordUpdate) check() error {
	if v, ok := ru.mutation.VideoURL(); ok {
		if err := record.VideoURLValidator(v); err != nil {
			return &ValidationError{Name: "video_url", err: fmt.Errorf(`ent: validator failed for field "Record.video_url": %w`, err)}
		}
	}
	if v, ok := ru.mutation.VideoID(); ok {
		if err := record.VideoIDValidator(v); err != nil {
			return &ValidationError{Name: "video_id", err: fmt.Errorf(`ent: validator failed for field "Record.video_id": %w`, err)}
		}
	}
	if v, ok := ru.mutation.FileLocation(); ok {
		if err := record.FileLocationValidator(v); err != nil {
			return &ValidationError{Name: "file_location", err: fmt.Errorf(`ent: validator failed for field "Record.file_location": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Status(); ok {
		if err := record.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Record.status": %w`, err)}
		}
	}
	return nil
}

func (ru *RecordUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(record.Table, record.Columns, sqlgraph.NewFieldSpec(record.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.VideoURL(); ok {
		_spec.SetField(record.FieldVideoURL, field.TypeString, value)
	}
	if value, ok := ru.mutation.VideoID(); ok {
		_spec.SetField(record.FieldVideoID, field.TypeString, value)
	}
	if value, ok := ru.mutation.FileLocation(); ok {
		_spec.SetField(record.FieldFileLocation, field.TypeString, value)
	}
	if ru.mutation.FileLocationCleared() {
		_spec.ClearField(record.FieldFileLocation, field.TypeString)
	}
	if value, ok := ru.mutation.Status(); ok {
		_spec.SetField(record.FieldStatus, field.TypeEnum, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{record.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RecordUpdateOne is the builder for updating a single Record entity.
type RecordUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RecordMutation
}

// SetVideoURL sets the "video_url" field.
func (ruo *RecordUpdateOne) SetVideoURL(s string) *RecordUpdateOne {
	ruo.mutation.SetVideoURL(s)
	return ruo
}

// SetVideoID sets the "video_id" field.
func (ruo *RecordUpdateOne) SetVideoID(s string) *RecordUpdateOne {
	ruo.mutation.SetVideoID(s)
	return ruo
}

// SetFileLocation sets the "file_location" field.
func (ruo *RecordUpdateOne) SetFileLocation(s string) *RecordUpdateOne {
	ruo.mutation.SetFileLocation(s)
	return ruo
}

// SetNillableFileLocation sets the "file_location" field if the given value is not nil.
func (ruo *RecordUpdateOne) SetNillableFileLocation(s *string) *RecordUpdateOne {
	if s != nil {
		ruo.SetFileLocation(*s)
	}
	return ruo
}

// ClearFileLocation clears the value of the "file_location" field.
func (ruo *RecordUpdateOne) ClearFileLocation() *RecordUpdateOne {
	ruo.mutation.ClearFileLocation()
	return ruo
}

// SetStatus sets the "status" field.
func (ruo *RecordUpdateOne) SetStatus(r record.Status) *RecordUpdateOne {
	ruo.mutation.SetStatus(r)
	return ruo
}

// Mutation returns the RecordMutation object of the builder.
func (ruo *RecordUpdateOne) Mutation() *RecordMutation {
	return ruo.mutation
}

// Where appends a list predicates to the RecordUpdate builder.
func (ruo *RecordUpdateOne) Where(ps ...predicate.Record) *RecordUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RecordUpdateOne) Select(field string, fields ...string) *RecordUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Record entity.
func (ruo *RecordUpdateOne) Save(ctx context.Context) (*Record, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RecordUpdateOne) SaveX(ctx context.Context) *Record {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RecordUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RecordUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RecordUpdateOne) check() error {
	if v, ok := ruo.mutation.VideoURL(); ok {
		if err := record.VideoURLValidator(v); err != nil {
			return &ValidationError{Name: "video_url", err: fmt.Errorf(`ent: validator failed for field "Record.video_url": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.VideoID(); ok {
		if err := record.VideoIDValidator(v); err != nil {
			return &ValidationError{Name: "video_id", err: fmt.Errorf(`ent: validator failed for field "Record.video_id": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.FileLocation(); ok {
		if err := record.FileLocationValidator(v); err != nil {
			return &ValidationError{Name: "file_location", err: fmt.Errorf(`ent: validator failed for field "Record.file_location": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Status(); ok {
		if err := record.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Record.status": %w`, err)}
		}
	}
	return nil
}

func (ruo *RecordUpdateOne) sqlSave(ctx context.Context) (_node *Record, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(record.Table, record.Columns, sqlgraph.NewFieldSpec(record.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Record.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, record.FieldID)
		for _, f := range fields {
			if !record.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != record.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.VideoURL(); ok {
		_spec.SetField(record.FieldVideoURL, field.TypeString, value)
	}
	if value, ok := ruo.mutation.VideoID(); ok {
		_spec.SetField(record.FieldVideoID, field.TypeString, value)
	}
	if value, ok := ruo.mutation.FileLocation(); ok {
		_spec.SetField(record.FieldFileLocation, field.TypeString, value)
	}
	if ruo.mutation.FileLocationCleared() {
		_spec.ClearField(record.FieldFileLocation, field.TypeString)
	}
	if value, ok := ruo.mutation.Status(); ok {
		_spec.SetField(record.FieldStatus, field.TypeEnum, value)
	}
	_node = &Record{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{record.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}