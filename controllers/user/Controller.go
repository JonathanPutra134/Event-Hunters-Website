package user

import (
	"event-hunters/dto"
	"event-hunters/helpers"
	"event-hunters/repository"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func LandingPageController(c *fiber.Ctx) error {
	return c.Render("landingpage/index", fiber.Map{})
}

func LoginTypePageController(c *fiber.Ctx) error {
	return c.Render("logintype/index", fiber.Map{})
}

func LoginPageController(c *fiber.Ctx) error {
	// userType := c.FormValue("user_type")
	// c.Locals("userType", userType)
	return c.Render("loginpage/index", fiber.Map{})
}

func RegistrationPageController(c *fiber.Ctx) error {
	return c.Render("registrationpage/index", fiber.Map{})
}

func RegistrationHandler(c *fiber.Ctx) error {
	// Extract form values
	fullName := c.FormValue("FullName")
	email := c.FormValue("Email")
	password := c.FormValue("Password")
	phoneNumber := c.FormValue("PhoneNumber")
	address := c.FormValue("Address")
	latitude := c.FormValue("Latitude")
	longitude := c.FormValue("Longitude")
	var errorMessages []string

	latFloat, err := helpers.ValidateLatitude(latitude)
	if err != nil {
		errorMessages = append(errorMessages, "Invalid Latitude")
	}
	// Validate longitude
	lonFloat, err := helpers.ValidateLongitude(longitude)
	if err != nil {
		errorMessages = append(errorMessages, "Invalid Longitude")
	}

	usrRegistrationReq := dto.UserRegistrationRequest{
		FullName:    fullName,
		Email:       email,
		Password:    password,
		PhoneNumber: phoneNumber,
		Address:     address,
		Latitude:    latitude,
		Longitude:   longitude,
	}
	if len(errorMessages) > 0 {
		errorMessage := strings.Join(errorMessages, ", ")
		fmt.Println(errorMessage)
		return c.Render("registrationpage/index", fiber.Map{
			"alertType":    "danger",
			"alertMessage": errorMessage,
		})
	}
	err = repository.RegisterUser(usrRegistrationReq, latFloat, lonFloat)
	if err != nil {
		return c.Render("registrationpage/index", fiber.Map{
			"alertType":    "danger",
			"alertMessage": "REGISTRATION PROCESS FAILED",
		})

	}

	return c.Render("loginpage/index", fiber.Map{
		"alertType":    "success", // Corrected the key name
		"alertMessage": "Registration Successful",
	})
}

func MainPageController(c *fiber.Ctx) error {
	baseURL := c.BaseURL() + "/mainpage"
	urlPath := c.Path()
	if urlPath == "/mainpage/eventdetails" {
		return c.Render("mainpage/eventdetails/index", fiber.Map{"BaseURL": baseURL, "Finished": false})
	}
	if urlPath == "/mainpage/recommendation" {
		return c.Render("mainpage/recommendation/index", fiber.Map{"BaseURL": baseURL})
	}
	if urlPath == "/mainpage/search" {
		return c.Render("mainpage/search/index", fiber.Map{"BaseURL": baseURL})
	}
	if urlPath == "/mainpage/mytickets" {
		return c.Render("mainpage/mytickets/index", fiber.Map{"BaseURL": baseURL})
	}
	if urlPath == "/mainpage/ticketinformation" {
		return c.Render("mainpage/ticketinformation/index", fiber.Map{"BaseURL": baseURL, "Finished": true})
	}
	fmt.Println("MASUK SINI NIH BOS")
	return c.Render("mainpage/home/index", fiber.Map{
		"BaseURL": baseURL,
	})
}

func MainPageEventDetailsController(c *fiber.Ctx) error {
	return c.Render("mainpage/home/index", fiber.Map{})
}
