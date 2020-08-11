package olap

import (
	"context"
	"errors"
)

var (
	ErrCellNotFound           = errors.New("cell not found")
	ErrCubeNotFound           = errors.New("cube not found")
	ErrElementNotFound        = errors.New("element not found")
	ErrComponentNotFound      = errors.New("component not found")
	ErrDimensionNotFound      = errors.New("dimension not found")
	ErrCubeAlreadyExists      = errors.New("cube already exists")
	ErrElementAlreadyExists   = errors.New("element already exists")
	ErrComponentAlreadyExists = errors.New("component already exists")
	ErrDimensionAlreadyExists = errors.New("dimension already exists")
)

type Cube struct {
	Name       string
	Dimensions []string
}

type Dimension struct {
	Name string
}

type Element struct {
	Name      string
	Dimension string
	Weight    float64
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

type Storage interface {
	// Cube methods
	AddCube(ctx context.Context, cube Cube) error
	GetCube(ctx context.Context, name string) (Cube, error)

	// Dimension methods
	AddDimension(ctx context.Context, dim Dimension) error
	GetDimension(ctx context.Context, name string) (Dimension, error)

	// Element methods
	AddElement(ctx context.Context, el Element) error
	GetElement(ctx context.Context, dim, name string) (Element, error)

	// Component methods
	AddComponent(ctx context.Context, tot, el Element) error
	GetComponent(ctx context.Context, dim, name string) (Element, error)
	Children(ctx context.Context, dim, name string) ([]Element, error)

	// Cell methods
	AddCell(ctx context.Context, cell Cell) error
	GetCell(ctx context.Context, cube string, elements ...string) (Cell, error)
}

type Server interface {
	Storage // Server extends the storage

	// Get & Put Cell
	Get(ctx context.Context, cube string, elements ...string) float64
	Put(ctx context.Context, value float64, cube string, elements ...string)

	Query(ctx context.Context, view View) (Rows, error)
	NewView(ctx context.Context, cube string, elements ...[]string) (View, error)
}

type Rows interface {
	Columns() []string
	Next() bool
	Scan(...interface{})
}
