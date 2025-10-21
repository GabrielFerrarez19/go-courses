package bar

import "Modulo02/foo"

func TakeFoo(i foo.Interface) {
	i.Interface()
}
