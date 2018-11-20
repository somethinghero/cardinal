package numerical

// //FloatValue float numerical value
// type FloatValue struct {
// 	value float64
// 	e     float64
// }

// // NewFloatValue new a float value
// func NewFloatValue(e float64) *FloatValue {
// 	ret := &FloatValue{}
// 	ret.e = e
// 	ret.value = e
// 	return ret
// }

//IntValue int numerical value
type IntValue struct {
	value int64
	e     int64
}

// NewIntValue new a int value
func NewIntValue(e int64) *IntValue {
	ret := &IntValue{}
	ret.e = e
	ret.value = e
	return ret
}
