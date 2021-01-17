// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vmkevv/duiztapi/ent/question"
	"github.com/vmkevv/duiztapi/ent/quiz"
	"github.com/vmkevv/duiztapi/ent/quizlangs"
	"github.com/vmkevv/duiztapi/ent/user"
)

// QuizCreate is the builder for creating a Quiz entity.
type QuizCreate struct {
	config
	mutation *QuizMutation
	hooks    []Hook
}

// SetURLImg sets the "url_img" field.
func (qc *QuizCreate) SetURLImg(s string) *QuizCreate {
	qc.mutation.SetURLImg(s)
	return qc
}

// AddQuestionIDs adds the "questions" edge to the Question entity by IDs.
func (qc *QuizCreate) AddQuestionIDs(ids ...int) *QuizCreate {
	qc.mutation.AddQuestionIDs(ids...)
	return qc
}

// AddQuestions adds the "questions" edges to the Question entity.
func (qc *QuizCreate) AddQuestions(q ...*Question) *QuizCreate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return qc.AddQuestionIDs(ids...)
}

// AddLangIDs adds the "langs" edge to the QuizLangs entity by IDs.
func (qc *QuizCreate) AddLangIDs(ids ...int) *QuizCreate {
	qc.mutation.AddLangIDs(ids...)
	return qc
}

// AddLangs adds the "langs" edges to the QuizLangs entity.
func (qc *QuizCreate) AddLangs(q ...*QuizLangs) *QuizCreate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return qc.AddLangIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (qc *QuizCreate) AddUserIDs(ids ...int) *QuizCreate {
	qc.mutation.AddUserIDs(ids...)
	return qc
}

// AddUsers adds the "users" edges to the User entity.
func (qc *QuizCreate) AddUsers(u ...*User) *QuizCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return qc.AddUserIDs(ids...)
}

// Mutation returns the QuizMutation object of the builder.
func (qc *QuizCreate) Mutation() *QuizMutation {
	return qc.mutation
}

// Save creates the Quiz in the database.
func (qc *QuizCreate) Save(ctx context.Context) (*Quiz, error) {
	var (
		err  error
		node *Quiz
	)
	if len(qc.hooks) == 0 {
		if err = qc.check(); err != nil {
			return nil, err
		}
		node, err = qc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuizMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = qc.check(); err != nil {
				return nil, err
			}
			qc.mutation = mutation
			node, err = qc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(qc.hooks) - 1; i >= 0; i-- {
			mut = qc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (qc *QuizCreate) SaveX(ctx context.Context) *Quiz {
	v, err := qc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (qc *QuizCreate) check() error {
	if _, ok := qc.mutation.URLImg(); !ok {
		return &ValidationError{Name: "url_img", err: errors.New("ent: missing required field \"url_img\"")}
	}
	return nil
}

func (qc *QuizCreate) sqlSave(ctx context.Context) (*Quiz, error) {
	_node, _spec := qc.createSpec()
	if err := sqlgraph.CreateNode(ctx, qc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (qc *QuizCreate) createSpec() (*Quiz, *sqlgraph.CreateSpec) {
	var (
		_node = &Quiz{config: qc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: quiz.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: quiz.FieldID,
			},
		}
	)
	if value, ok := qc.mutation.URLImg(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: quiz.FieldURLImg,
		})
		_node.URLImg = value
	}
	if nodes := qc.mutation.QuestionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   quiz.QuestionsTable,
			Columns: []string{quiz.QuestionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: question.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qc.mutation.LangsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   quiz.LangsTable,
			Columns: []string{quiz.LangsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: quizlangs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := qc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   quiz.UsersTable,
			Columns: quiz.UsersPrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// QuizCreateBulk is the builder for creating many Quiz entities in bulk.
type QuizCreateBulk struct {
	config
	builders []*QuizCreate
}

// Save creates the Quiz entities in the database.
func (qcb *QuizCreateBulk) Save(ctx context.Context) ([]*Quiz, error) {
	specs := make([]*sqlgraph.CreateSpec, len(qcb.builders))
	nodes := make([]*Quiz, len(qcb.builders))
	mutators := make([]Mutator, len(qcb.builders))
	for i := range qcb.builders {
		func(i int, root context.Context) {
			builder := qcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*QuizMutation)
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
					_, err = mutators[i+1].Mutate(root, qcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, qcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, qcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (qcb *QuizCreateBulk) SaveX(ctx context.Context) []*Quiz {
	v, err := qcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
