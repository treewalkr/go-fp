package immutability

// Demonstrates immutability by returning a new slice with an appended element
func AppendImmutable(slice []int, element int) []int {
	newSlice := make([]int, len(slice)+1)
	copy(newSlice, slice)
	newSlice[len(slice)] = element
	return newSlice
}

// Demonstrates immutability for structs by creating a copy and modifying the copy
type Person struct {
	Name string
	Age  int
}

func UpdateAgeImmutable(p Person, newAge int) Person {
	return Person{Name: p.Name, Age: newAge}
}

// Example of avoiding mutation with maps
func AddToMapImmutable(original map[string]int, key string, value int) map[string]int {
	newMap := make(map[string]int, len(original)+1)
	for k, v := range original {
		newMap[k] = v
	}
	newMap[key] = value
	return newMap
}

// Demonstrates immutability with a nested struct (deep copy)
type Address struct {
	City  string
	State string
}

type Employee struct {
	Name    string
	Age     int
	Address Address
}

// Updates the City field in the Address without altering the original Employee
func UpdateCityImmutable(e Employee, newCity string) Employee {
	// Creates a new Employee with a modified Address
	return Employee{
		Name:    e.Name,
		Age:     e.Age,
		Address: Address{City: newCity, State: e.Address.State},
	}
}

// Immutable data example with a constant list of values
var predefinedRoles = []string{"Admin", "User", "Guest"}

// Returns a copy of predefinedRoles instead of modifying or exposing the original
func GetPredefinedRoles() []string {
	rolesCopy := make([]string, len(predefinedRoles))
	copy(rolesCopy, predefinedRoles)
	return rolesCopy
}

// Example of preventing global mutable state
// An immutable list of API endpoints that should not be altered
var apiEndpoints = map[string]string{
	"auth": "/api/auth",
	"user": "/api/user",
}

// Retrieve a copy of the API endpoints to ensure immutability
func GetAPIEndpoints() map[string]string {
	newMap := make(map[string]string, len(apiEndpoints))
	for k, v := range apiEndpoints {
		newMap[k] = v
	}
	return newMap
}
