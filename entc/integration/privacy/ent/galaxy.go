// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/galaxy"
)

// Galaxy is the model entity for the Galaxy schema.
type Galaxy struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Type holds the value of the "type" field.
	Type galaxy.Type `json:"type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GalaxyQuery when eager-loading is set.
	Edges GalaxyEdges `json:"edges"`
}

// GalaxyEdges holds the relations/edges for other nodes in the graph.
type GalaxyEdges struct {
	// Planets holds the value of the planets edge.
	Planets []*Planet
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PlanetsOrErr returns the Planets value or an error if the edge
// was not loaded in eager-loading.
func (e GalaxyEdges) PlanetsOrErr() ([]*Planet, error) {
	if e.loadedTypes[0] {
		return e.Planets, nil
	}
	return nil, &NotLoadedError{edge: "planets"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Galaxy) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
		&sql.NullString{}, // type
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Galaxy fields.
func (ga *Galaxy) assignValues(values ...interface{}) error {
	if m, n := len(values), len(galaxy.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ga.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		ga.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field type", values[1])
	} else if value.Valid {
		ga.Type = galaxy.Type(value.String)
	}
	return nil
}

// QueryPlanets queries the planets edge of the Galaxy.
func (ga *Galaxy) QueryPlanets() *PlanetQuery {
	return (&GalaxyClient{config: ga.config}).QueryPlanets(ga)
}

// Update returns a builder for updating this Galaxy.
// Note that, you need to call Galaxy.Unwrap() before calling this method, if this Galaxy
// was returned from a transaction, and the transaction was committed or rolled back.
func (ga *Galaxy) Update() *GalaxyUpdateOne {
	return (&GalaxyClient{config: ga.config}).UpdateOne(ga)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ga *Galaxy) Unwrap() *Galaxy {
	tx, ok := ga.config.driver.(*txDriver)
	if !ok {
		panic("ent: Galaxy is not a transactional entity")
	}
	ga.config.driver = tx.drv
	return ga
}

// String implements the fmt.Stringer.
func (ga *Galaxy) String() string {
	var builder strings.Builder
	builder.WriteString("Galaxy(")
	builder.WriteString(fmt.Sprintf("id=%v", ga.ID))
	builder.WriteString(", name=")
	builder.WriteString(ga.Name)
	builder.WriteString(", type=")
	builder.WriteString(fmt.Sprintf("%v", ga.Type))
	builder.WriteByte(')')
	return builder.String()
}

// Galaxies is a parsable slice of Galaxy.
type Galaxies []*Galaxy

func (ga Galaxies) config(cfg config) {
	for _i := range ga {
		ga[_i].config = cfg
	}
}