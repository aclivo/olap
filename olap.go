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

type Process struct {
	Name       string
	Run        func(ctx context.Context) error
	Attributes []Attr
}

type Rule struct {
	Cube  string
	Match func(ctx context.Context, elements ...string) bool
	Eval  func(ctx context.Context, elements ...string) float64
}

type Storage interface {
	AddCube(ctx context.Context, cube Cube) error
	GetCubeByName(ctx context.Context, name string) (Cube, error)
	CubeExists(ctx context.Context, name string) (bool, error)

	AddDimension(ctx context.Context, dim Dimension) error
	GetDimensionByName(ctx context.Context, name string) (Dimension, error)
	DimensionExists(ctx context.Context, name string) (bool, error)

	AddElement(ctx context.Context, el Element) error
	GetElementByName(ctx context.Context, dim, name string) (Element, error)
	ElementExists(ctx context.Context, dim, name string) (bool, error)

	AddComponent(ctx context.Context, tot, el Element) error
	ComponentExists(ctx context.Context, dim, name string) (bool, error)
	Children(ctx context.Context, dim, name string) ([]Element, error)

	AddCell(ctx context.Context, cell Cell) error
	GetCellByName(ctx context.Context, cube string, elements ...string) (Cell, error)
}

type Server interface {
	Storage

	Get(ctx context.Context, cube string, elements ...string) (float64, error)
	Put(ctx context.Context, value float64, cube string, elements ...string) error

	AddProcess(ctx context.Context, p Process) error
	ExecuteProcess(ctx context.Context, name string) error
	AddRules(ctx context.Context, cube string, rules []Rule) error
	GetStorage(ctx context.Context) (Storage, error)

	NewView(ctx context.Context, cube string, elements ...[]string) (View, error)
	Query(ctx context.Context, view View) (Rows, error)
}

type Rows interface {
	Columns() []string
	Next() bool
	Scan(...interface{})
}

type RuleExpr interface{}
type RuleList []RuleExpr
