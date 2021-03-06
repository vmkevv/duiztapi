// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/vmkevv/duiztapi/ent/answerlangs"
	"github.com/vmkevv/duiztapi/ent/i18n"
	"github.com/vmkevv/duiztapi/ent/predicate"
	"github.com/vmkevv/duiztapi/ent/questionlangs"
	"github.com/vmkevv/duiztapi/ent/quizlangs"
)

// I18nUpdate is the builder for updating I18n entities.
type I18nUpdate struct {
	config
	hooks    []Hook
	mutation *I18nMutation
}

// Where adds a new predicate for the I18nUpdate builder.
func (iu *I18nUpdate) Where(ps ...predicate.I18n) *I18nUpdate {
	iu.mutation.predicates = append(iu.mutation.predicates, ps...)
	return iu
}

// SetCode sets the "code" field.
func (iu *I18nUpdate) SetCode(s string) *I18nUpdate {
	iu.mutation.SetCode(s)
	return iu
}

// SetLanguage sets the "language" field.
func (iu *I18nUpdate) SetLanguage(s string) *I18nUpdate {
	iu.mutation.SetLanguage(s)
	return iu
}

// AddAnswerLangIDs adds the "answer_langs" edge to the AnswerLangs entity by IDs.
func (iu *I18nUpdate) AddAnswerLangIDs(ids ...int) *I18nUpdate {
	iu.mutation.AddAnswerLangIDs(ids...)
	return iu
}

// AddAnswerLangs adds the "answer_langs" edges to the AnswerLangs entity.
func (iu *I18nUpdate) AddAnswerLangs(a ...*AnswerLangs) *I18nUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iu.AddAnswerLangIDs(ids...)
}

// AddQuestionLangIDs adds the "question_langs" edge to the QuestionLangs entity by IDs.
func (iu *I18nUpdate) AddQuestionLangIDs(ids ...int) *I18nUpdate {
	iu.mutation.AddQuestionLangIDs(ids...)
	return iu
}

// AddQuestionLangs adds the "question_langs" edges to the QuestionLangs entity.
func (iu *I18nUpdate) AddQuestionLangs(q ...*QuestionLangs) *I18nUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return iu.AddQuestionLangIDs(ids...)
}

// AddQuizLangIDs adds the "quiz_langs" edge to the QuizLangs entity by IDs.
func (iu *I18nUpdate) AddQuizLangIDs(ids ...int) *I18nUpdate {
	iu.mutation.AddQuizLangIDs(ids...)
	return iu
}

// AddQuizLangs adds the "quiz_langs" edges to the QuizLangs entity.
func (iu *I18nUpdate) AddQuizLangs(q ...*QuizLangs) *I18nUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return iu.AddQuizLangIDs(ids...)
}

// Mutation returns the I18nMutation object of the builder.
func (iu *I18nUpdate) Mutation() *I18nMutation {
	return iu.mutation
}

// ClearAnswerLangs clears all "answer_langs" edges to the AnswerLangs entity.
func (iu *I18nUpdate) ClearAnswerLangs() *I18nUpdate {
	iu.mutation.ClearAnswerLangs()
	return iu
}

// RemoveAnswerLangIDs removes the "answer_langs" edge to AnswerLangs entities by IDs.
func (iu *I18nUpdate) RemoveAnswerLangIDs(ids ...int) *I18nUpdate {
	iu.mutation.RemoveAnswerLangIDs(ids...)
	return iu
}

// RemoveAnswerLangs removes "answer_langs" edges to AnswerLangs entities.
func (iu *I18nUpdate) RemoveAnswerLangs(a ...*AnswerLangs) *I18nUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iu.RemoveAnswerLangIDs(ids...)
}

// ClearQuestionLangs clears all "question_langs" edges to the QuestionLangs entity.
func (iu *I18nUpdate) ClearQuestionLangs() *I18nUpdate {
	iu.mutation.ClearQuestionLangs()
	return iu
}

// RemoveQuestionLangIDs removes the "question_langs" edge to QuestionLangs entities by IDs.
func (iu *I18nUpdate) RemoveQuestionLangIDs(ids ...int) *I18nUpdate {
	iu.mutation.RemoveQuestionLangIDs(ids...)
	return iu
}

// RemoveQuestionLangs removes "question_langs" edges to QuestionLangs entities.
func (iu *I18nUpdate) RemoveQuestionLangs(q ...*QuestionLangs) *I18nUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return iu.RemoveQuestionLangIDs(ids...)
}

// ClearQuizLangs clears all "quiz_langs" edges to the QuizLangs entity.
func (iu *I18nUpdate) ClearQuizLangs() *I18nUpdate {
	iu.mutation.ClearQuizLangs()
	return iu
}

// RemoveQuizLangIDs removes the "quiz_langs" edge to QuizLangs entities by IDs.
func (iu *I18nUpdate) RemoveQuizLangIDs(ids ...int) *I18nUpdate {
	iu.mutation.RemoveQuizLangIDs(ids...)
	return iu
}

// RemoveQuizLangs removes "quiz_langs" edges to QuizLangs entities.
func (iu *I18nUpdate) RemoveQuizLangs(q ...*QuizLangs) *I18nUpdate {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return iu.RemoveQuizLangIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *I18nUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*I18nMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *I18nUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *I18nUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *I18nUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *I18nUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   i18n.Table,
			Columns: i18n.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: i18n.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: i18n.FieldCode,
		})
	}
	if value, ok := iu.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: i18n.FieldLanguage,
		})
	}
	if iu.mutation.AnswerLangsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.AnswerLangsTable,
			Columns: []string{i18n.AnswerLangsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: answerlangs.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedAnswerLangsIDs(); len(nodes) > 0 && !iu.mutation.AnswerLangsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.AnswerLangsTable,
			Columns: []string{i18n.AnswerLangsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: answerlangs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.AnswerLangsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.AnswerLangsTable,
			Columns: []string{i18n.AnswerLangsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: answerlangs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.QuestionLangsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.QuestionLangsTable,
			Columns: []string{i18n.QuestionLangsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: questionlangs.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedQuestionLangsIDs(); len(nodes) > 0 && !iu.mutation.QuestionLangsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.QuestionLangsTable,
			Columns: []string{i18n.QuestionLangsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: questionlangs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.QuestionLangsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.QuestionLangsTable,
			Columns: []string{i18n.QuestionLangsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: questionlangs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iu.mutation.QuizLangsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.QuizLangsTable,
			Columns: []string{i18n.QuizLangsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: quizlangs.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.RemovedQuizLangsIDs(); len(nodes) > 0 && !iu.mutation.QuizLangsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.QuizLangsTable,
			Columns: []string{i18n.QuizLangsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.QuizLangsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.QuizLangsTable,
			Columns: []string{i18n.QuizLangsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{i18n.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// I18nUpdateOne is the builder for updating a single I18n entity.
type I18nUpdateOne struct {
	config
	hooks    []Hook
	mutation *I18nMutation
}

// SetCode sets the "code" field.
func (iuo *I18nUpdateOne) SetCode(s string) *I18nUpdateOne {
	iuo.mutation.SetCode(s)
	return iuo
}

// SetLanguage sets the "language" field.
func (iuo *I18nUpdateOne) SetLanguage(s string) *I18nUpdateOne {
	iuo.mutation.SetLanguage(s)
	return iuo
}

// AddAnswerLangIDs adds the "answer_langs" edge to the AnswerLangs entity by IDs.
func (iuo *I18nUpdateOne) AddAnswerLangIDs(ids ...int) *I18nUpdateOne {
	iuo.mutation.AddAnswerLangIDs(ids...)
	return iuo
}

// AddAnswerLangs adds the "answer_langs" edges to the AnswerLangs entity.
func (iuo *I18nUpdateOne) AddAnswerLangs(a ...*AnswerLangs) *I18nUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iuo.AddAnswerLangIDs(ids...)
}

// AddQuestionLangIDs adds the "question_langs" edge to the QuestionLangs entity by IDs.
func (iuo *I18nUpdateOne) AddQuestionLangIDs(ids ...int) *I18nUpdateOne {
	iuo.mutation.AddQuestionLangIDs(ids...)
	return iuo
}

// AddQuestionLangs adds the "question_langs" edges to the QuestionLangs entity.
func (iuo *I18nUpdateOne) AddQuestionLangs(q ...*QuestionLangs) *I18nUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return iuo.AddQuestionLangIDs(ids...)
}

// AddQuizLangIDs adds the "quiz_langs" edge to the QuizLangs entity by IDs.
func (iuo *I18nUpdateOne) AddQuizLangIDs(ids ...int) *I18nUpdateOne {
	iuo.mutation.AddQuizLangIDs(ids...)
	return iuo
}

// AddQuizLangs adds the "quiz_langs" edges to the QuizLangs entity.
func (iuo *I18nUpdateOne) AddQuizLangs(q ...*QuizLangs) *I18nUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return iuo.AddQuizLangIDs(ids...)
}

// Mutation returns the I18nMutation object of the builder.
func (iuo *I18nUpdateOne) Mutation() *I18nMutation {
	return iuo.mutation
}

// ClearAnswerLangs clears all "answer_langs" edges to the AnswerLangs entity.
func (iuo *I18nUpdateOne) ClearAnswerLangs() *I18nUpdateOne {
	iuo.mutation.ClearAnswerLangs()
	return iuo
}

// RemoveAnswerLangIDs removes the "answer_langs" edge to AnswerLangs entities by IDs.
func (iuo *I18nUpdateOne) RemoveAnswerLangIDs(ids ...int) *I18nUpdateOne {
	iuo.mutation.RemoveAnswerLangIDs(ids...)
	return iuo
}

// RemoveAnswerLangs removes "answer_langs" edges to AnswerLangs entities.
func (iuo *I18nUpdateOne) RemoveAnswerLangs(a ...*AnswerLangs) *I18nUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return iuo.RemoveAnswerLangIDs(ids...)
}

// ClearQuestionLangs clears all "question_langs" edges to the QuestionLangs entity.
func (iuo *I18nUpdateOne) ClearQuestionLangs() *I18nUpdateOne {
	iuo.mutation.ClearQuestionLangs()
	return iuo
}

// RemoveQuestionLangIDs removes the "question_langs" edge to QuestionLangs entities by IDs.
func (iuo *I18nUpdateOne) RemoveQuestionLangIDs(ids ...int) *I18nUpdateOne {
	iuo.mutation.RemoveQuestionLangIDs(ids...)
	return iuo
}

// RemoveQuestionLangs removes "question_langs" edges to QuestionLangs entities.
func (iuo *I18nUpdateOne) RemoveQuestionLangs(q ...*QuestionLangs) *I18nUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return iuo.RemoveQuestionLangIDs(ids...)
}

// ClearQuizLangs clears all "quiz_langs" edges to the QuizLangs entity.
func (iuo *I18nUpdateOne) ClearQuizLangs() *I18nUpdateOne {
	iuo.mutation.ClearQuizLangs()
	return iuo
}

// RemoveQuizLangIDs removes the "quiz_langs" edge to QuizLangs entities by IDs.
func (iuo *I18nUpdateOne) RemoveQuizLangIDs(ids ...int) *I18nUpdateOne {
	iuo.mutation.RemoveQuizLangIDs(ids...)
	return iuo
}

// RemoveQuizLangs removes "quiz_langs" edges to QuizLangs entities.
func (iuo *I18nUpdateOne) RemoveQuizLangs(q ...*QuizLangs) *I18nUpdateOne {
	ids := make([]int, len(q))
	for i := range q {
		ids[i] = q[i].ID
	}
	return iuo.RemoveQuizLangIDs(ids...)
}

// Save executes the query and returns the updated I18n entity.
func (iuo *I18nUpdateOne) Save(ctx context.Context) (*I18n, error) {
	var (
		err  error
		node *I18n
	)
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*I18nMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *I18nUpdateOne) SaveX(ctx context.Context) *I18n {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *I18nUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *I18nUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *I18nUpdateOne) sqlSave(ctx context.Context) (_node *I18n, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   i18n.Table,
			Columns: i18n.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: i18n.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing I18n.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := iuo.mutation.Code(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: i18n.FieldCode,
		})
	}
	if value, ok := iuo.mutation.Language(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: i18n.FieldLanguage,
		})
	}
	if iuo.mutation.AnswerLangsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.AnswerLangsTable,
			Columns: []string{i18n.AnswerLangsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: answerlangs.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedAnswerLangsIDs(); len(nodes) > 0 && !iuo.mutation.AnswerLangsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.AnswerLangsTable,
			Columns: []string{i18n.AnswerLangsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: answerlangs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.AnswerLangsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.AnswerLangsTable,
			Columns: []string{i18n.AnswerLangsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: answerlangs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.QuestionLangsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.QuestionLangsTable,
			Columns: []string{i18n.QuestionLangsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: questionlangs.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedQuestionLangsIDs(); len(nodes) > 0 && !iuo.mutation.QuestionLangsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.QuestionLangsTable,
			Columns: []string{i18n.QuestionLangsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: questionlangs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.QuestionLangsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.QuestionLangsTable,
			Columns: []string{i18n.QuestionLangsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: questionlangs.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if iuo.mutation.QuizLangsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.QuizLangsTable,
			Columns: []string{i18n.QuizLangsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: quizlangs.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.RemovedQuizLangsIDs(); len(nodes) > 0 && !iuo.mutation.QuizLangsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.QuizLangsTable,
			Columns: []string{i18n.QuizLangsColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.QuizLangsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   i18n.QuizLangsTable,
			Columns: []string{i18n.QuizLangsColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &I18n{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{i18n.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
