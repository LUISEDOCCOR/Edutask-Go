package waitingList_model

import (
	"edutask/src/database"
)

func AddEmail(email string) error {
	error := database.RedisClient.RPush(database.Ctx, "emails", email).Err()
	return error
}

func GetEmailsCounter() (int64, error) {
	counter, err := database.RedisClient.LLen(database.Ctx, "emails").Result()
	return counter, err
}
