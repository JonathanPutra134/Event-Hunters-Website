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
			Longitude:      null.StringFrom("-73.985428"),
			Latitude:       null.StringFrom("40.748817"),
			Email:          null.StringFrom("john.doe@example.com"),
			Password:       helpers.HashPassword("securepassword1"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		{
			Name:           null.StringFrom("Jonggun"),
			Longitude:      null.StringFrom("-73.985428"),
			Latitude:       null.StringFrom("40.748817"),
			Email:          null.StringFrom("john.doe@example.com"),
			Password:       helpers.HashPassword("securepassword1"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		{
			Name:           null.StringFrom("usagi_coser"),
			Longitude:      null.StringFrom("-73.985428"),
			Latitude:       null.StringFrom("40.748817"),
			Email:          null.StringFrom("john.doe@example.com"),
			Password:       helpers.HashPassword("securepassword1"),
			ProfilePicture: null.StringFrom("john_doe.jpg"),
			PhoneNumber:    null.StringFrom("+1234567890"),
		},
		// Add more users as needed
	}

	return UsersToInsert
}