package dataset

import (
	"event-hunters/helpers"
	"event-hunters/models"

	"github.com/volatiletech/null/v8"
)

func InitializeUsers() []models.User {
	UsersToInsert := []models.User{
		{
			Name:           null.StringFrom("John Doe"),
			Latitude:       null.StringFrom("-6.165"),
			Longitude:      null.StringFrom("106.6856"),
			Email:          null.StringFrom("john.doe@example.com"),
			Password:       helpers.HashPassword("securepassword1"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
	}

	return UsersToInsert
}
