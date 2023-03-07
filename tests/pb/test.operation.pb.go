package pb

var TestServerOperationGroup = map[string][]string{
	"auth": {
		"/cart.service.v1.Cart/GetCart",
		"/cart.service.v1.Cart/DeleteCart",
	},
	"auth2": {
		"/cart.service.v1.Cart/GetCart",
		"/cart.service.v1.Cart/DeleteCart",
	},
}

func GetTestServerOperationByGroup(group string) []string {
	return TestServerOperationGroup[group]
}
