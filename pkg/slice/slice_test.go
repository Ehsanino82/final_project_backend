package slice_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"final_project_backend/pkg/slice"
)

type sliceUtilsTestSuit struct {
	suite.Suite
}

func TestSliceUtilsTestSuit(t *testing.T) {
	t.Parallel()

	suite.Run(t, new(sliceUtilsTestSuit))
}

func (t *sliceUtilsTestSuit) TestIntSliceDifferenceBothFilled() {
	first := []int{1, 2, 3}
	second := []int{1, 2}

	diffA := slice.Difference(first, second)
	diffB := slice.Difference(second, first)

	assert.Equal(t.T(), diffA, []int{3})
	assert.True(t.T(), diffB == nil)

	// diffB is nil and above test case is true but we can not use assert.Equal
	// to check null as it checks whether their bytes arrays are equal or not.
	// just nil has nil type so []int([]int(nil)) != <nil>(<nil>):
	assert.NotEqual(t.T(), diffB, nil)

	// null is nil with type []int
	var null []int

	assert.Equal(t.T(), diffB, null)
}

func (t *sliceUtilsTestSuit) TestIntSliceDifferenceOneEmpty() {
	first := []int{1, 2, 3}
	second := []int{}

	diffA := slice.Difference(first, second)
	diffB := slice.Difference(second, first)

	assert.Equal(t.T(), diffA, []int{1, 2, 3})
	assert.True(t.T(), diffB == nil)
}

func (t *sliceUtilsTestSuit) TestIntSliceDifferenceOneNil() {
	first := []int{1, 2, 3}

	var second []int

	diffA := slice.Difference(first, second)
	diffB := slice.Difference(second, first)

	assert.Equal(t.T(), diffA, []int{1, 2, 3})
	assert.True(t.T(), diffB == nil)
}

func (t *sliceUtilsTestSuit) TestIntSliceDifferenceBothEmpty() {
	first := []int{}
	second := []int{}

	diffA := slice.Difference(first, second)
	diffB := slice.Difference(second, first)

	assert.True(t.T(), diffA == nil)
	assert.True(t.T(), diffB == nil)
}

func (t *sliceUtilsTestSuit) TestIntSliceDifferenceBothNil() {
	var (
		first  []int
		second []int
	)

	diffA := slice.Difference(first, second)
	diffB := slice.Difference(second, first)

	assert.True(t.T(), diffA == nil)
	assert.True(t.T(), diffB == nil)
}

func (t *sliceUtilsTestSuit) TestSelfDeclaredStructSliceDifference() {
	type Person struct {
		Name string
	}

	first := []Person{{"ALI"}, {"REZA"}}
	second := []Person{{Name: "ALI"}}

	diffA := slice.Difference(first, second)
	diffB := slice.Difference(second, first)

	assert.Equal(t.T(), diffA, []Person{{"REZA"}})
	assert.True(t.T(), diffB == nil)

	// null is nil with type []Person
	var null []Person

	assert.Equal(t.T(), diffB, null)
}

func (t *sliceUtilsTestSuit) TestPointerDifference() {
	common := 1
	inFirst := 2
	inSecond := 2
	first := []*int{&common, &inFirst}
	second := []*int{&common, &inSecond}
	assert.Equal(t.T(), slice.Difference(first, second), []*int{&inFirst})
	assert.Equal(t.T(), slice.Difference(second, first), []*int{&inSecond})
}
