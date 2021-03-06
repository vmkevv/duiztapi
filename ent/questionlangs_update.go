// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vmkevv/duiztapi/ent/i18n"
	"github.com/vmkevv/duiztapi/ent/predicate"
	"github.com/vmkevv/duiztapi/ent/question"
	"github.com/vmkevv/duiztapi/ent/questionlangs"
)

// QuestionLangsUpdate is the builder for updating QuestionLangs entities.
type QuestionLangsUpdate struct {
	config
	hooks    []Hook
	mutation *QuestionLangsMutation
}

// Where adds a new predicate for the QuestionLangsUpdate builder.
func (qlu *QuestionLangsUpdate) Where(ps ...predicate.QuestionLangs) *QuestionLangsUpdate {
	qlu.mutation.predicates = append(qlu.mutation.predicates, ps...)
	return qlu
}

// SetTitle sets the "title" field.
func (qlu *QuestionLangsUpdate) SetTitle(s string) *QuestionLangsUpdate {
	qlu.mutation.SetTitle(s)
	return qlu
}

// SetBody sets the "body" field.
func (qlu *QuestionLangsUpdate) SetBody(s string) *QuestionLangsUpdate {
	qlu.mutation.SetBody(s)
	return qlu
}

// SetExplanation sets the "explanation" field.
func (qlu *QuestionLangsUpdate) SetExplanation(s string) *QuestionLangsUpdate {
	qlu.mutation.SetExplanation(s)
	return qlu
}

// SetI18nID sets the "i18n" edge to the I18n entity by ID.
func (qlu *QuestionLangsUpdate) SetI18nID(id int) *QuestionLangsUpdate {
	qlu.mutation.SetI18nID(id)
	return qlu
}

// SetNillableI18nID sets the "i18n" edge to the I18n entity by ID if the given value is not nil.
func (qlu *QuestionLangsUpdate) SetNillableI18nID(id *int) *QuestionLangsUpdate {
	if id != nil {
		qlu = qlu.SetI18nID(*id)
	}
	return qlu
}

// SetI18n sets the "i18n" edge to the I18n entity.
func (qlu *QuestionLangsUpdate) SetI18n(i *I18n) *QuestionLangsUpdate {
	return qlu.SetI18nID(i.ID)
}

// SetQuestionID sets the "question" edge to the Question entity by ID.
func (qlu *QuestionLangsUpdate) SetQuestionID(id int) *QuestionLangsUpdate {
	qlu.mutation.SetQuestionID(id)
	return qlu
}

// SetNillableQuestionID sets the "question" edge to the Question entity by ID if the given value is not nil.
func (qlu *QuestionLangsUpdate) SetNillableQuestionID(id *int) *QuestionLangsUpdate {
	if id != nil {
		qlu = qlu.SetQuestionID(*id)
	}
	return qlu
}

// SetQuestion sets the "question" edge to the Question entity.
func (qlu *QuestionLangsUpdate) SetQuestion(q *Question) *QuestionLangsUpdate {
	return qlu.SetQuestionID(q.ID)
}

// Mutation returns the QuestionLangsMutation object of the builder.
func (qlu *QuestionLangsUpdate) Mutation() *QuestionLangsMutation {
	return qlu.mutation
}

// ClearI18n clears the "i18n" edge to the I18n entity.
func (qlu *QuestionLangsUpdate) ClearI18n() *QuestionLangsUpdate {
	qlu.mutation.ClearI18n()
	return qlu
}

// ClearQuestion clears the "question" edge to the Question entity.
func (qlu *QuestionLangsUpdate) ClearQuestion() *QuestionLangsUpdate {
	qlu.mutation.ClearQuestion()
	return qlu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (qlu *QuestionLangsUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(qlu.hooks) == 0 {
		affected, err = qlu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuestionLangsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qlu.mutation = mutation
			affected, err = qlu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(qlu.hooks) - 1; i >= 0; i-- {
			mut = qlu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qlu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (qlu *QuestionLangsUpdate) SaveX(ctx context.Context) int {
	affected, err := qlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (qlu *QuestionLangsUpdate) Exec(ctx context.Context) error {
	_, err := qlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qlu *QuestionLangsUpdate) ExecX(ctx context.Context) {
	if err := qlu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (qlu *QuestionLangsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   questionlangs.Table,
			Columns: questionlangs.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: questionlangs.FieldID,
			},
		},
	}
	if ps := qlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := qlu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: questionlangs.FieldTitle,
		})
	}
	if value, ok := qlu.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: questionlangs.FieldBody,
		})
	}
	if value, ok := qlu.mutation.Explanation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: questionlangs.FieldExplanation,
		})
	}
	if qlu.mutation.I18nCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionlangs.I18nTable,
			Columns: []string{questionlangs.I18nColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: i18n.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qlu.mutation.I18nIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionlangs.I18nTable,
			Columns: []string{questionlangs.I18nColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: i18n.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qlu.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionlangs.QuestionTable,
			Columns: []string{questionlangs.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: question.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qlu.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionlangs.QuestionTable,
			Columns: []string{questionlangs.QuestionColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, qlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{questionlangs.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// QuestionLangsUpdateOne is the builder for updating a single QuestionLangs entity.
type QuestionLangsUpdateOne struct {
	config
	hooks    []Hook
	mutation *QuestionLangsMutation
}

// SetTitle sets the "title" field.
func (qluo *QuestionLangsUpdateOne) SetTitle(s string) *QuestionLangsUpdateOne {
	qluo.mutation.SetTitle(s)
	return qluo
}

// SetBody sets the "body" field.
func (qluo *QuestionLangsUpdateOne) SetBody(s string) *QuestionLangsUpdateOne {
	qluo.mutation.SetBody(s)
	return qluo
}

// SetExplanation sets the "explanation" field.
func (qluo *QuestionLangsUpdateOne) SetExplanation(s string) *QuestionLangsUpdateOne {
	qluo.mutation.SetExplanation(s)
	return qluo
}

// SetI18nID sets the "i18n" edge to the I18n entity by ID.
func (qluo *QuestionLangsUpdateOne) SetI18nID(id int) *QuestionLangsUpdateOne {
	qluo.mutation.SetI18nID(id)
	return qluo
}

// SetNillableI18nID sets the "i18n" edge to the I18n entity by ID if the given value is not nil.
func (qluo *QuestionLangsUpdateOne) SetNillableI18nID(id *int) *QuestionLangsUpdateOne {
	if id != nil {
		qluo = qluo.SetI18nID(*id)
	}
	return qluo
}

// SetI18n sets the "i18n" edge to the I18n entity.
func (qluo *QuestionLangsUpdateOne) SetI18n(i *I18n) *QuestionLangsUpdateOne {
	return qluo.SetI18nID(i.ID)
}

// SetQuestionID sets the "question" edge to the Question entity by ID.
func (qluo *QuestionLangsUpdateOne) SetQuestionID(id int) *QuestionLangsUpdateOne {
	qluo.mutation.SetQuestionID(id)
	return qluo
}

// SetNillableQuestionID sets the "question" edge to the Question entity by ID if the given value is not nil.
func (qluo *QuestionLangsUpdateOne) SetNillableQuestionID(id *int) *QuestionLangsUpdateOne {
	if id != nil {
		qluo = qluo.SetQuestionID(*id)
	}
	return qluo
}

// SetQuestion sets the "question" edge to the Question entity.
func (qluo *QuestionLangsUpdateOne) SetQuestion(q *Question) *QuestionLangsUpdateOne {
	return qluo.SetQuestionID(q.ID)
}

// Mutation returns the QuestionLangsMutation object of the builder.
func (qluo *QuestionLangsUpdateOne) Mutation() *QuestionLangsMutation {
	return qluo.mutation
}

// ClearI18n clears the "i18n" edge to the I18n entity.
func (qluo *QuestionLangsUpdateOne) ClearI18n() *QuestionLangsUpdateOne {
	qluo.mutation.ClearI18n()
	return qluo
}

// ClearQuestion clears the "question" edge to the Question entity.
func (qluo *QuestionLangsUpdateOne) ClearQuestion() *QuestionLangsUpdateOne {
	qluo.mutation.ClearQuestion()
	return qluo
}

// Save executes the query and returns the updated QuestionLangs entity.
func (qluo *QuestionLangsUpdateOne) Save(ctx context.Context) (*QuestionLangs, error) {
	var (
		err  error
		node *QuestionLangs
	)
	if len(qluo.hooks) == 0 {
		node, err = qluo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*QuestionLangsMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			qluo.mutation = mutation
			node, err = qluo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(qluo.hooks) - 1; i >= 0; i-- {
			mut = qluo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, qluo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (qluo *QuestionLangsUpdateOne) SaveX(ctx context.Context) *QuestionLangs {
	node, err := qluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (qluo *QuestionLangsUpdateOne) Exec(ctx context.Context) error {
	_, err := qluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (qluo *QuestionLangsUpdateOne) ExecX(ctx context.Context) {
	if err := qluo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (qluo *QuestionLangsUpdateOne) sqlSave(ctx context.Context) (_node *QuestionLangs, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   questionlangs.Table,
			Columns: questionlangs.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: questionlangs.FieldID,
			},
		},
	}
	id, ok := qluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing QuestionLangs.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := qluo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: questionlangs.FieldTitle,
		})
	}
	if value, ok := qluo.mutation.Body(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: questionlangs.FieldBody,
		})
	}
	if value, ok := qluo.mutation.Explanation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: questionlangs.FieldExplanation,
		})
	}
	if qluo.mutation.I18nCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionlangs.I18nTable,
			Columns: []string{questionlangs.I18nColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: i18n.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qluo.mutation.I18nIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionlangs.I18nTable,
			Columns: []string{questionlangs.I18nColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: i18n.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if qluo.mutation.QuestionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionlangs.QuestionTable,
			Columns: []string{questionlangs.QuestionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: question.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := qluo.mutation.QuestionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   questionlangs.QuestionTable,
			Columns: []string{questionlangs.QuestionColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &QuestionLangs{config: qluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, qluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{questionlangs.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
