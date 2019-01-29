// Copyright (c) 2019, The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package etensor

import "github.com/apache/arrow/go/arrow"

//go:generate tmpl -i -data=numeric.tmpldata numeric.gen.go.tmpl

// Tensor is the general interface for n-dimensional tensors
type Tensor interface {
	// Len returns the number of elements in the tensor.
	Len() int

	// DataType returns the type of data, using arrow.DataType (ID() is the arrow.Type enum value)
	DataType() arrow.DataType

	// Shapes returns the size in each dimension of the tensor. (Shape is the full Shape struct)
	Shapes() []int

	// Strides returns the number of elements to step in each dimension when traversing the tensor.
	Strides() []int

	// Shape64 returns the size in each dimension using int64 (arrow compatbile)
	Shape64() []int64

	// Strides64 returns the strides in each dimension using int64 (arrow compatbile)
	Strides64() []int64

	// NumDims returns the number of dimensions of the tensor.
	NumDims() int

	// Dim returns the size of the given dimension
	Dim(i int) int

	// DimNames returns the string slice of dimension names
	DimNames() []string

	// DimName returns the name of the i-th dimension.
	DimName(i int) string

	IsContiguous() bool
	IsRowMajor() bool
	IsColMajor() bool

	// Offset returns the flat 1D array / slice index into an element at the given n-dimensional index.
	// No checking is done on the length or size of the index values relative to the shape of the tensor.
	Offset(i []int) int

	// Float64Val returns the value of given index as a float64
	Float64Val(i []int) float64

	// SetFloat64 sets the value of given index as a float64
	SetFloat64(i []int, val float64)

	// StringVal returns the value of given index as a string
	StringVal(i []int) string

	// SetString sets the value of given index as a string
	SetString(i []int, val string)

	// AggFloat64 applies given aggregation function to each element in the tensor, using float64
	// conversions of the values.  init is the initial value for the agg variable.  returns final
	// aggregate value
	AggFloat64(fun func(val float64, agg float64) float64, init float64) float64

	// EvalFloat64 applies given function to each element in the tensor, using float64
	// conversions of the values, and puts the results into given float64 slice, which is
	// ensured to be of the proper length
	EvalFloat64(fun func(val float64) float64, res *[]float64)

	// UpdtFloat64 applies given function to each element in the tensor, using float64
	// conversions of the values, and writes the results back into the same tensor values
	UpdtFloat64(fun func(val float64) float64)

	// CloneTensor clones this tensor returning a Tensor interface.
	// There is a type-specific Clone() method as well for each tensor.
	CloneTensor() Tensor

	// SetShape sets the shape parameters of the tensor, and resizes backing storage appropriately.
	// existing RowMajor or ColMajor stride preference will be used if strides is nil, and
	// existing names will be preserved if nil
	SetShape(shape, strides []int, names []string)

	// AddRows adds n rows (outer-most dimension) to RowMajor organized tensor.
	// Does nothing for other stride layouts
	AddRows(n int)

	// SetNumRows sets the number of rows (outer-most dimension) in a RowMajor organized tensor.
	// Does nothing for other stride layouts
	SetNumRows(rows int)
}

// Check impl
var _ Tensor = (*Float32)(nil)
