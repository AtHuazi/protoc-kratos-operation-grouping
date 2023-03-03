package pb

var groupOperations = map[string][]string{
    "auth": {
        "cart.service.v1.Cart.GetCart",
    },
    "auth2": {
        "cart.service.v1.Cart.GetCart",
    },
}

func GetOperations(group string) []string {
	return groupOperations[group]
}
