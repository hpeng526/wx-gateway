package fn

import "reflect"

/**
 * high class func map
 */
func Map(xs interface{}, f interface{}) interface{} {
	vf := reflect.ValueOf(f)
	vxs := reflect.ValueOf(xs)

	tys := reflect.SliceOf(vf.Type().Out(0))
	vys := reflect.MakeSlice(tys, vxs.Len(), vxs.Len())

	for i := 0; i < vxs.Len(); i++ {
		y := vf.Call([]reflect.Value{vxs.Index(i)})[0]
		vys.Index(i).Set(y)
	}
	return vys.Interface()
}
