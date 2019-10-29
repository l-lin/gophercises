package user

// User of a tweet
type User struct {
	ID   string `json:"id_str"`
	Name string `json:"name"`
}

// UniqueUsers filter unique users from a given user slice
func UniqueUsers(users []User) []User {
	result := []User{}
	m := make(map[string]bool, 0)
	for _, u := range users {
		if _, ok := m[u.ID]; !ok {
			m[u.ID] = true
			result = append(result, u)
		}
	}
	return result
}

// Merge users into a single slice
func Merge(users1 []User, users2 []User) []User {
	result := []User{}
	result = append(result, users1...)
	result = append(result, users2...)
	return UniqueUsers(result)
}

// Save users in a temp file
func Save(users []User) error {
	return nil
}
