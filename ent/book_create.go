// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ritarock/manga/ent/book"
)

// BookCreate is the builder for creating a Book entity.
type BookCreate struct {
	config
	mutation *BookMutation
	hooks    []Hook
}

// SetIsbn sets the "isbn" field.
func (bc *BookCreate) SetIsbn(s string) *BookCreate {
	bc.mutation.SetIsbn(s)
	return bc
}

// SetNillableIsbn sets the "isbn" field if the given value is not nil.
func (bc *BookCreate) SetNillableIsbn(s *string) *BookCreate {
	if s != nil {
		bc.SetIsbn(*s)
	}
	return bc
}

// SetTitle sets the "title" field.
func (bc *BookCreate) SetTitle(s string) *BookCreate {
	bc.mutation.SetTitle(s)
	return bc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (bc *BookCreate) SetNillableTitle(s *string) *BookCreate {
	if s != nil {
		bc.SetTitle(*s)
	}
	return bc
}

// SetPublisher sets the "publisher" field.
func (bc *BookCreate) SetPublisher(s string) *BookCreate {
	bc.mutation.SetPublisher(s)
	return bc
}

// SetNillablePublisher sets the "publisher" field if the given value is not nil.
func (bc *BookCreate) SetNillablePublisher(s *string) *BookCreate {
	if s != nil {
		bc.SetPublisher(*s)
	}
	return bc
}

// SetPubdate sets the "pubdate" field.
func (bc *BookCreate) SetPubdate(s string) *BookCreate {
	bc.mutation.SetPubdate(s)
	return bc
}

// SetNillablePubdate sets the "pubdate" field if the given value is not nil.
func (bc *BookCreate) SetNillablePubdate(s *string) *BookCreate {
	if s != nil {
		bc.SetPubdate(*s)
	}
	return bc
}

// SetCover sets the "cover" field.
func (bc *BookCreate) SetCover(s string) *BookCreate {
	bc.mutation.SetCover(s)
	return bc
}

// SetNillableCover sets the "cover" field if the given value is not nil.
func (bc *BookCreate) SetNillableCover(s *string) *BookCreate {
	if s != nil {
		bc.SetCover(*s)
	}
	return bc
}

// SetAuthor sets the "author" field.
func (bc *BookCreate) SetAuthor(s string) *BookCreate {
	bc.mutation.SetAuthor(s)
	return bc
}

// SetNillableAuthor sets the "author" field if the given value is not nil.
func (bc *BookCreate) SetNillableAuthor(s *string) *BookCreate {
	if s != nil {
		bc.SetAuthor(*s)
	}
	return bc
}

// SetSubjectCode sets the "subject_code" field.
func (bc *BookCreate) SetSubjectCode(s string) *BookCreate {
	bc.mutation.SetSubjectCode(s)
	return bc
}

// SetNillableSubjectCode sets the "subject_code" field if the given value is not nil.
func (bc *BookCreate) SetNillableSubjectCode(s *string) *BookCreate {
	if s != nil {
		bc.SetSubjectCode(*s)
	}
	return bc
}

// Mutation returns the BookMutation object of the builder.
func (bc *BookCreate) Mutation() *BookMutation {
	return bc.mutation
}

// Save creates the Book in the database.
func (bc *BookCreate) Save(ctx context.Context) (*Book, error) {
	return withHooks[*Book, BookMutation](ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookCreate) SaveX(ctx context.Context) *Book {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BookCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BookCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BookCreate) check() error {
	return nil
}

func (bc *BookCreate) sqlSave(ctx context.Context) (*Book, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BookCreate) createSpec() (*Book, *sqlgraph.CreateSpec) {
	var (
		_node = &Book{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: book.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: book.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.Isbn(); ok {
		_spec.SetField(book.FieldIsbn, field.TypeString, value)
		_node.Isbn = value
	}
	if value, ok := bc.mutation.Title(); ok {
		_spec.SetField(book.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := bc.mutation.Publisher(); ok {
		_spec.SetField(book.FieldPublisher, field.TypeString, value)
		_node.Publisher = value
	}
	if value, ok := bc.mutation.Pubdate(); ok {
		_spec.SetField(book.FieldPubdate, field.TypeString, value)
		_node.Pubdate = value
	}
	if value, ok := bc.mutation.Cover(); ok {
		_spec.SetField(book.FieldCover, field.TypeString, value)
		_node.Cover = value
	}
	if value, ok := bc.mutation.Author(); ok {
		_spec.SetField(book.FieldAuthor, field.TypeString, value)
		_node.Author = value
	}
	if value, ok := bc.mutation.SubjectCode(); ok {
		_spec.SetField(book.FieldSubjectCode, field.TypeString, value)
		_node.SubjectCode = value
	}
	return _node, _spec
}

// BookCreateBulk is the builder for creating many Book entities in bulk.
type BookCreateBulk struct {
	config
	builders []*BookCreate
}

// Save creates the Book entities in the database.
func (bcb *BookCreateBulk) Save(ctx context.Context) ([]*Book, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Book, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookMutation)
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
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BookCreateBulk) SaveX(ctx context.Context) []*Book {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BookCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BookCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
