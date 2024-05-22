package immutability

import (
	"reflect"
	"testing"
)

func TestAppendImmutable(t *testing.T) {
	originalSlice := []int{1, 2, 3}
	element := 4

	// Call the function
	newSlice := AppendImmutable(originalSlice, element)

	// Check that originalSlice remains unchanged
	if !reflect.DeepEqual(originalSlice, []int{1, 2, 3}) {
		t.Errorf("originalSlice modified: got %v, want %v", originalSlice, []int{1, 2, 3})
	}

	// Check that newSlice contains the added element
	expected := []int{1, 2, 3, 4}
	if !reflect.DeepEqual(newSlice, expected) {
		t.Errorf("AppendImmutable() = %v, want %v", newSlice, expected)
	}
}

func TestUpdateAgeImmutable(t *testing.T) {
	originalPerson := Person{Name: "Alice", Age: 30}
	newAge := 35

	// Call the function
	updatedPerson := UpdateAgeImmutable(originalPerson, newAge)

	// Check that originalPerson remains unchanged
	if originalPerson.Age != 30 {
		t.Errorf("originalPerson modified: got %v, want %v", originalPerson.Age, 30)
	}

	// Check that updatedPerson has the new age
	if updatedPerson.Age != newAge {
		t.Errorf("UpdateAgeImmutable() = %v, want %v", updatedPerson.Age, newAge)
	}
}

func TestAddToMapImmutable(t *testing.T) {
	originalMap := map[string]int{"apple": 1, "banana": 2}
	newKey := "cherry"
	newValue := 3

	// Call the function
	newMap := AddToMapImmutable(originalMap, newKey, newValue)

	// Check that originalMap remains unchanged
	expectedOriginal := map[string]int{"apple": 1, "banana": 2}
	if !reflect.DeepEqual(originalMap, expectedOriginal) {
		t.Errorf("originalMap modified: got %v, want %v", originalMap, expectedOriginal)
	}

	// Check that newMap contains the new key-value pair
	expectedNew := map[string]int{"apple": 1, "banana": 2, "cherry": 3}
	if !reflect.DeepEqual(newMap, expectedNew) {
		t.Errorf("AddToMapImmutable() = %v, want %v", newMap, expectedNew)
	}
}

func TestGetPredefinedRoles(t *testing.T) {
	// Call the function
	roles := GetPredefinedRoles()

	// Check that the returned slice matches the expected roles
	expectedRoles := []string{"Admin", "User", "Guest"}
	if !reflect.DeepEqual(roles, expectedRoles) {
		t.Errorf("GetPredefinedRoles() = %v, want %v", roles, expectedRoles)
	}

	// Modify the returned slice and check that it doesn't affect the original
	roles[0] = "Modified"
	if reflect.DeepEqual(GetPredefinedRoles(), roles) {
		t.Error("GetPredefinedRoles() returned a slice that affects the original predefinedRoles")
	}
}

func TestGetAPIEndpoints(t *testing.T) {
	// Call the function
	endpoints := GetAPIEndpoints()

	// Expected endpoints
	expectedEndpoints := map[string]string{
		"auth": "/api/auth",
		"user": "/api/user",
	}
	if !reflect.DeepEqual(endpoints, expectedEndpoints) {
		t.Errorf("GetAPIEndpoints() = %v, want %v", endpoints, expectedEndpoints)
	}

	// Modify the returned map and check that it doesn't affect the original
	endpoints["auth"] = "modified"
	if reflect.DeepEqual(GetAPIEndpoints(), endpoints) {
		t.Error("GetAPIEndpoints() returned a map that affects the original apiEndpoints")
	}
}
