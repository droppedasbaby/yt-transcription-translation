// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/GrewalAS/yt-transcription-translation/ent/record"
	"github.com/google/uuid"
)

// RecordCreate is the builder for creating a Record entity.
type RecordCreate struct {
	config
	mutation *RecordMutation
	hooks    []Hook
}

// SetVideoURL sets the "video_url" field.
func (rc *RecordCreate) SetVideoURL(s string) *RecordCreate {
	rc.mutation.SetVideoURL(s)
	return rc
}

// SetVideoID sets the "video_id" field.
func (rc *RecordCreate) SetVideoID(s string) *RecordCreate {
	rc.mutation.SetVideoID(s)
	return rc
}

// SetFileLocation sets the "file_location" field.
func (rc *RecordCreate) SetFileLocation(s string) *RecordCreate {
	rc.mutation.SetFileLocation(s)
	return rc
}

// SetNillableFileLocation sets the "file_location" field if the given value is not nil.
func (rc *RecordCreate) SetNillableFileLocation(s *string) *RecordCreate {
	if s != nil {
		rc.SetFileLocation(*s)
	}
	return rc
}

// SetStatus sets the "status" field.
func (rc *RecordCreate) SetStatus(r record.Status) *RecordCreate {
	rc.mutation.SetStatus(r)
	return rc
}

// SetRunID sets the "run_id" field.
func (rc *RecordCreate) SetRunID(u uuid.UUID) *RecordCreate {
	rc.mutation.SetRunID(u)
	return rc
}

// SetNillableRunID sets the "run_id" field if the given value is not nil.
func (rc *RecordCreate) SetNillableRunID(u *uuid.UUID) *RecordCreate {
	if u != nil {
		rc.SetRunID(*u)
	}
	return rc
}

// Mutation returns the RecordMutation object of the builder.
func (rc *RecordCreate) Mutation() *RecordMutation {
	return rc.mutation
}

// Save creates the Record in the database.
func (rc *RecordCreate) Save(ctx context.Context) (*Record, error) {
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RecordCreate) SaveX(ctx context.Context) *Record {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RecordCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RecordCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RecordCreate) check() error {
	if _, ok := rc.mutation.VideoURL(); !ok {
		return &ValidationError{Name: "video_url", err: errors.New(`ent: missing required field "Record.video_url"`)}
	}
	if v, ok := rc.mutation.VideoURL(); ok {
		if err := record.VideoURLValidator(v); err != nil {
			return &ValidationError{Name: "video_url", err: fmt.Errorf(`ent: validator failed for field "Record.video_url": %w`, err)}
		}
	}
	if _, ok := rc.mutation.VideoID(); !ok {
		return &ValidationError{Name: "video_id", err: errors.New(`ent: missing required field "Record.video_id"`)}
	}
	if v, ok := rc.mutation.VideoID(); ok {
		if err := record.VideoIDValidator(v); err != nil {
			return &ValidationError{Name: "video_id", err: fmt.Errorf(`ent: validator failed for field "Record.video_id": %w`, err)}
		}
	}
	if v, ok := rc.mutation.FileLocation(); ok {
		if err := record.FileLocationValidator(v); err != nil {
			return &ValidationError{Name: "file_location", err: fmt.Errorf(`ent: validator failed for field "Record.file_location": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Record.status"`)}
	}
	if v, ok := rc.mutation.Status(); ok {
		if err := record.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Record.status": %w`, err)}
		}
	}
	return nil
}

func (rc *RecordCreate) sqlSave(ctx context.Context) (*Record, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RecordCreate) createSpec() (*Record, *sqlgraph.CreateSpec) {
	var (
		_node = &Record{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(record.Table, sqlgraph.NewFieldSpec(record.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.VideoURL(); ok {
		_spec.SetField(record.FieldVideoURL, field.TypeString, value)
		_node.VideoURL = value
	}
	if value, ok := rc.mutation.VideoID(); ok {
		_spec.SetField(record.FieldVideoID, field.TypeString, value)
		_node.VideoID = value
	}
	if value, ok := rc.mutation.FileLocation(); ok {
		_spec.SetField(record.FieldFileLocation, field.TypeString, value)
		_node.FileLocation = value
	}
	if value, ok := rc.mutation.Status(); ok {
		_spec.SetField(record.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := rc.mutation.RunID(); ok {
		_spec.SetField(record.FieldRunID, field.TypeUUID, value)
		_node.RunID = value
	}
	return _node, _spec
}

// RecordCreateBulk is the builder for creating many Record entities in bulk.
type RecordCreateBulk struct {
	config
	err      error
	builders []*RecordCreate
}

// Save creates the Record entities in the database.
func (rcb *RecordCreateBulk) Save(ctx context.Context) ([]*Record, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Record, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecordMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RecordCreateBulk) SaveX(ctx context.Context) []*Record {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RecordCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RecordCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
