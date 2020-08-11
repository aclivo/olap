package olap

import (
	"context"
	"errors"
)

var (
	ErrCellNotFound = errors.New("cell not found")

	ErrCubeNotFound      = errors.New("cube not found")
	ErrCubeAlreadyExists = errors.New("cube already exists")

	ErrElementNotFound      = errors.New("element not found")
	ErrElementAlreadyExists = errors.New("element already exists")

	ErrComponentNotFound      = errors.New("component not found")
	ErrComponentAlreadyExists = errors.New("component already exists")

	ErrDimensionNotFound      = errors.New("dimension not found")
	ErrDimensionAlreadyExists = errors.New("dimension already exists")
)

type Attr struct {
	Name  string
	Value string
}

type Cube struct {
	Name       string
	Dimensions []string
	Attributes []Attr
}

type Dimension struct {
	Name       string
	Attributes []Attr
}

type Element struct {
	Name       string
	Dimension  string
	Weight     float64
	Attributes []Attr
}

type Cell struct {
	Cube     string
	Elements []string
	Value    float64
}

type View struct {
	Cube   string
	Slices map[string][]string
}

type CubeRules struct {
	Cube string
}

type Process struct {
	Name string
}

type Storage interface {
	AddCube(ctx context.Context, cube Cube) error
	GetCube(ctx context.Context, name string) (Cube, error)

	AddDimension(ctx context.Context, dim Dimension) error
	GetDimension(ctx context.Context, name string) (Dimension, error)

	AddElement(ctx context.Context, el Element) error
	GetElement(ctx context.Context, dim, name string) (Element, error)

	AddComponent(ctx context.Context, tot, el Element) error
	Children(ctx context.Context, dim, name string) ([]Element, error)

	AddCell(ctx context.Context, cell Cell) error
	GetCell(ctx context.Context, cube string, elements ...string) (Cell, error)

	AddCubeRules(ctx context.Context, cube string, rules CubeRules) error
	GetCubeRules(ctx context.Context, cube string) (CubeRules, error)

	AddProcess(ctx context.Context, process Process) error
	GetProcess(ctx context.Context, process string) (Process, error)
}

type Server interface {
	Storage

	Get(ctx context.Context, cube string, elements ...string) float64
	Put(ctx context.Context, value float64, cube string, elements ...string)

	Query(ctx context.Context, view View) (Rows, error)
	NewView(ctx context.Context, cube string, elements ...[]string) (View, error)

	ExecuteProcess(ctx context.Context, name string) error
}

type Rows interface {
	Columns() []string
	Next() bool
	Scan(...interface{})
}
