// http request handler functions
package handlers

import (
	"encoding/json"
	"fmt"
	"server/auth"
	"server/store"
	"strings"

	"github.com/gin-gonic/gin"
)

// User + token combo
type Auth struct {
	User  store.CleanUser
	Token string
}

// Login unmarshal format
type LoginStruct struct {
	Name     string
	Password string
}

// Function and object aliases
var memoryLayer = store.Store.Get()
var saveToFile = store.Store.Save

// Auth constructor
func SignAuth(user *store.CleanUser) Auth {
	token, err := auth.Sign(*user)
	if err != nil {
		fmt.Println(err)
		return Auth{}
	}

	return Auth{*user, token}
}

// Login handler
func Login(c *gin.Context) {
	var user LoginStruct
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	// verify user
	correct := memoryLayer.CorrectCredentials(user.Name, user.Password)

	if !correct {
		c.JSON(200, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": SignAuth(memoryLayer.GetUser(user.Name)),
	})
}

// Add user handler (same logic as login)
func SignUp(c *gin.Context) {
	var user LoginStruct
	err := c.BindJSON(&user)

	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	// verify user
	added := memoryLayer.AddUser(user.Name, user.Password, "")
	saveToFile()

	if !added {
		c.JSON(200, gin.H{
			"error": "Username already exists",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": SignAuth(memoryLayer.GetUser(user.Name)),
	})
}

func SetAPIKey(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(200, gin.H{
			"error": "Not logged in",
		})
		return
	}
	var data map[string]string
	err := c.BindJSON(&data)

	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	// update preferences
	updateUser := memoryLayer.SetAPIKey(user.(store.CleanUser).Name, data["key"])
	validKey := memoryLayer.GetDevices(user.(store.CleanUser).Name)
	saveToFile()

	if strings.Contains(validKey, "Invalid API") {
		// invalid key
		memoryLayer.SetAPIKey(user.(store.CleanUser).Name, "")
		saveToFile()
		c.JSON(200, gin.H{
			"error": "Invalid API key",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": SignAuth(updateUser),
	})
}

func GetDevices(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(200, gin.H{
			"error": "Not logged in",
		})
		return
	}
	// get devices
	devices := memoryLayer.GetDevices(user.(store.CleanUser).Name)

	data := make(map[string]interface{})
	err := json.Unmarshal([]byte(devices), &data)

	if err != nil {
		c.JSON(200, gin.H{
			"error": devices,
		})
		return
	}

	c.JSON(200, gin.H{
		"result": data,
	})
}

func ReverseGeocode(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(200, gin.H{
			"error": "Not logged in",
		})
		return
	}

	var data map[string]string
	err := c.BindJSON(&data)

	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}

	lookup := memoryLayer.ReverseGeocode(user.(store.CleanUser).Name, data["lat"], data["lon"])

	data_ := make(map[string]interface{})
	err = json.Unmarshal([]byte(lookup), &data_)

	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"result": data_,
	})
}

// Need to be logged in to view or alter preferences
func Preferences(c *gin.Context) {
	user, exists := c.Get("user")
	if !exists {
		c.JSON(200, gin.H{
			"error": "Not logged in",
		})
		return
	}

	cleaned := user.(store.CleanUser)

	if c.Request.Method == "GET" {
		preferences := memoryLayer.GetPreferences(cleaned.Name)
		fmt.Println(preferences)
		c.JSON(200, gin.H{
			"result": preferences,
		})
	} else if c.Request.Method == "POST" {
		var data map[string]interface{}
		err := c.BindJSON(&data)

		if err != nil {
			c.JSON(200, gin.H{
				"error": err.Error(),
			})
			return
		}

		// update preferences
		memoryLayer.SetPreferences(cleaned.Name, data)
		saveToFile()

		c.JSON(200, gin.H{
			"result": "Success",
		})
	}
}
