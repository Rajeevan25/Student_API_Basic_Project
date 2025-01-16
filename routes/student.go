package routes

import (
	"errors"
	database "fiber_api/Database"
	"fiber_api/models"

	"github.com/gofiber/fiber/v2"
)

type Student struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseStudent(student models.Student) Student {
	return Student{
		ID:        student.ID,
		FirstName: student.FirstName,
		LastName:  student.LastName,
	}
}
//creating students
func CreateStudent(c *fiber.Ctx) error {
	var user models.Student
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	database.Database.Db.Create(&user)
	responseUser := CreateResponseStudent(user)

	return c.Status(200).JSON(responseUser)
}

// retrieving all students
func GetStudents(c *fiber.Ctx) error {
	users := []models.Student{}
	database.Database.Db.Find(&users)
	responseStudents := []Student{}
	for _, user := range users {
		responseStudent := CreateResponseStudent(user)
		responseStudents = append(responseStudents, responseStudent)
	}

	return c.Status(200).JSON(responseStudents)
}
func findStudent(id int, user *models.Student) error {
	database.Database.Db.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("student does not exist")
	}
	return nil
}

// finding student by ID 
func GetStudent(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.Student

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	if err := findStudent(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseStudent := CreateResponseStudent(user)

	return c.Status(200).JSON(responseStudent)
}

// deleting student
func DeleteStudent(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.Student

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findStudent(id, &user)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err = database.Database.Db.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}
	return c.Status(200).JSON("Successfully deleted User")
}

// updating student
func UpdateStudent(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var user models.Student

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is an integer")
	}

	err = findStudent(id, &user)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}
    // can not update ID
	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	var updateData UpdateUser

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateData.FirstName
	user.LastName = updateData.LastName

	database.Database.Db.Save(&user)

	responseUser := CreateResponseStudent(user)

	return c.Status(200).JSON(responseUser)

}
