package constant

const (
	ROLE_ADMIN     = "admin"
	ROLE_CASHIER   = "cashier"
	ROLE_WAREHOUSE = "warehouse"
)

var (
	mapRoleCanBeCreated = map[string]bool{
		ROLE_ADMIN:     false,
		ROLE_CASHIER:   true,
		ROLE_WAREHOUSE: true,
	}
)

func IsRoleCanBeCreated(role string) bool {
	return mapRoleCanBeCreated[role]
}
