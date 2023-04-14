// Code generated by ent, DO NOT EDIT.

package testcasesuite

import (
	"galileo/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldEQ(FieldName, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v uint32) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldEQ(FieldCreatedBy, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldContainsFold(FieldName, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v uint32) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v uint32) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...uint32) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...uint32) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v uint32) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v uint32) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v uint32) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v uint32) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(sql.FieldLTE(FieldCreatedBy, v))
}

// HasTestcase applies the HasEdge predicate on the "testcase" edge.
func HasTestcase() predicate.TestCaseSuite {
	return predicate.TestCaseSuite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TestcaseTable, TestcaseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestcaseWith applies the HasEdge predicate on the "testcase" edge with a given conditions (other predicates).
func HasTestcaseWith(preds ...predicate.TestCase) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TestcaseInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TestcaseTable, TestcaseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.TestCaseSuite) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.TestCaseSuite) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(func(s *sql.Selector) {
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
func Not(p predicate.TestCaseSuite) predicate.TestCaseSuite {
	return predicate.TestCaseSuite(func(s *sql.Selector) {
		p(s.Not())
	})
}