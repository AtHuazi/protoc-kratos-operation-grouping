package pb

var TestServerOperationGroup = map[string][]string{
	"auth": {
		"cart.service.v1.Cart.GetCart",
	},
	"auth2": {
		"cart.service.v1.Cart.GetCart",
	},
}

func GetTestServerOperationByGroup(group string) []string {
	return TestServerOperationGroup[group]
}
