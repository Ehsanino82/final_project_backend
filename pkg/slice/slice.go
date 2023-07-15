package slice

// Difference function takes 2 slice of the same type and returns all members of
// first one which is not in second one.
//
// Slices passed to this function must be comparable as they are used in map key.
//
// exp1
//
//	first := []int{1, 2, 3}
//	second := []int{1, 2}
//	fmt.Println(Difference(first, second)) // [3]
//	fmt.Println(Difference(second, first)) // []
//
// There was an important part in above example. the second Difference will not return
// an empty slice but nil. but the nil with type of []int so when printing to stdout
// we will see an empty slice.
// in golang nil has type and have different behaviors in different places. for example,
// we can not range over a nil with type int, but we can range over a nil with type slice.
// in our case we can always check if there is no difference like this:
//
//	result := Difference(second, first)
//	if result == nil {
//	    fmt.Println("result is empty") // this wil be printed
//	}
//
// Consider this that output is not interface so result can be nil. (if we were working
// with interface we had an interface with a type filled and value nil so this interface
// is not equal to nil)
//
// exp2
//
//	type Person struct {
//	  	Name string
//	}
//
//	first := []Person{{"ALI"}, {"REZA"}}
//	second := []Person{{"ALI"}}
//	fmt.Println(Difference(first, second)) // [{REZA}]
//	fmt.Println(Difference(second, first)) // [{}]
//
// exp3
//
//	// be careful with pointers. pointers are comparable but we will not check its value.
//	common := 1
//	inFirst := 2
//	inSecond := 2
//	first := []*int{&common, &inFirst}
//	second := []*int{&common, &inSecond}
//	fmt.Println(Difference(first, second)) // [0xFFFFF]
//	fmt.Println(Difference(second, first)) // [0x00000]
//
// Check testcases.
//
//nolint:gofumpt,gofmt
func Difference[T comparable](first, second []T) []T {
	secondValues := make(map[T]struct{})

	for _, t := range second {
		secondValues[t] = struct{}{}
	}

	var result []T

	for _, t := range first {
		_, ok := secondValues[t]
		if !ok {
			result = append(result, t)
		}
	}

	return result
}

func Contains[T comparable](list []T, target T) bool {
	for id := range list {
		if list[id] == target {
			return true
		}
	}

	return false
}
