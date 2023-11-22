package main

import "errors"

type mapUser struct {
	name                 string
	phoneNumber          int
	scheduledForDeletion bool
}

func getUserMap(names []string, phoneNumbers []int) (map[string]mapUser, error) {
	userMap := make(map[string]mapUser)
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("names and phoneNumbers must have the same length")
	}

	for i := 0; i < len(names); i++ {
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = mapUser{
			name:        name,
			phoneNumber: phoneNumber,
		}
	}
	return userMap, nil
}

func deleteIfNecessary(userMap map[string]mapUser, name string) (deleted bool, err error) {
	existingUser, ok := userMap[name]
	if !ok {
		return false, errors.New("user does not exist")
	}

	if existingUser.scheduledForDeletion {
		delete(userMap, name)
		return true, nil
	}
	return false, nil
}

func mapTest() {
	// names := []string{"Xipeng", "A", "B"}
	// phoneNumbers := []int{1, 2, 3}
	// userMap, err := getUserMap(names, phoneNumbers)
	// if err != nil {
	// 	panic(err)
	// }
	// for name, user := range userMap {
	// 	println(name, user.phoneNumber)
	// }

	userMap := make(map[string]mapUser)
	userMap["Xipeng"] = mapUser{
		name:                 "Xipeng",
		phoneNumber:          1,
		scheduledForDeletion: false,
	}

	deleted, err := deleteIfNecessary(userMap, "Xipeng")
	if err != nil {
		panic(err)
	}
	println(deleted)

	userMap["A"] = mapUser{
		name:                 "A",
		phoneNumber:          2,
		scheduledForDeletion: true,
	}

	deleted, err = deleteIfNecessary(userMap, "A")
	if err != nil {
		panic(err)
	}
	println(deleted)

}
