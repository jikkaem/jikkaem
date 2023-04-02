package user_service

import "jikkaem/internal/shared/mongodb"

func GetUser(id string) {
	userColl, err := mongodb.GetColl("users")
	if err != nil {
		panic(err)
	}
	_ = userColl
}
