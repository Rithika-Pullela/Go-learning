package main

import (
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	//"net/smtp"
)

type Employee struct {
	gorm.Model
	Name         string
	Email        string
	DepartmentID uint
	Department   *Department
	CompanyID    uint
	Company      *Company
}

type Department struct {
	gorm.Model
	Name      string
	Manager   *Employee
	CompanyID uint
	Company   *Company
}

type Company struct {
	gorm.Model
	Name        string
	Employees   []Employee
	Departments []Department
}

func SendWelcomeEmail(mail string) {
	fmt.Printf("Mail Successfully  sent to %s", mail)
}

func (emp *Employee) AfterCreate() {
	SendWelcomeEmail(emp.Email)
}

// func sendWelcomeEmail(to, username string) error {
// 	from := "rithikapullela@gmail.com"
// 	password := "P@rithika211101"
// 	smtpServer := "smtp.gmail.com"
// 	smtpPort := 587

// 	// Set up authentication information
// 	auth := smtp.PlainAuth("", from, password, smtpServer)

// 	// Set up the email message
// 	message := []byte(fmt.Sprintf("Subject: Welcome to my website!\n\nDear %s,\n\nThank you for signing up for my website!", username))

// 	// Send the email
// 	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpServer, smtpPort), auth, from, []string{to}, message)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func main() {
	host := flag.String("host", "localhost", "PostgreSQL host")
	port := flag.String("port", "5432", "PostgreSQL port")
	user := flag.String("user", "postgres", "PostgreSQL user")
	password := flag.String("password", "", "PostgreSQL password")
	database := flag.String("database", "", "PostgreSQL database name")
	flag.Parse()

	// Construct connection string from flags
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", *host, *port, *user, *password, *database)

	// Open a connection to the database
	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	db.Debug().DropTable(&Employee{})
	db.Debug().DropTable(&Company{})
	db.Debug().DropTable(&Department{})
	db.Debug().AutoMigrate(&Employee{}, &Company{}, &Department{})
	fmt.Println("Connection Successful")
	company := Company{
		Name: "BeautifulCode",
		Employees: []Employee{
			{Name: "Rithika", Email: "rithika.pullela@beautifulcode.in"},
			{Name: "Rithin", Email: "rithin.pullela@beautifulcode.in"},
			{Name: "Sai", Email: "sai@beautifulcode.in"},
		},
		Departments: []Department{
			{Name: "Development"},
		},
	}

	db.Debug().Save(&company)

	e := Employee{Name: "Rithika"}
	manager := Employee{Name: "Sai"}
	dep := Department{Name: "Development"}

	db.Preloads("Department").Preloads("Company").Where(&manager).Find(&manager)
	db.Preloads("Department").Preloads("Company").Where(&e).Find(&e)
	db.Where(&dep).Find(&dep)

	manager.DepartmentID = dep.ID
	dep.Manager = &manager
	e.Department = &dep

	e2 := Employee{Name: "Rithin"}
	db.Preloads("Department").Preloads("Company").Where(&e2).Find(&e2)
	e2.Department = &dep

	db.Save(&e)
	db.Save(&e2)
	db.Debug().Find(&e)
	db.Debug().Find(&manager)

	var records []Employee
	db.Preload("Department").Preload("Company").Find(&records)

	for _, r := range records {
		fmt.Println(r.Name)
		fmt.Println(r.Email)
		fmt.Println(r.DepartmentID)
		fmt.Println(r.Department.Name)

	}

}
