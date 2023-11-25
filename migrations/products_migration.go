package migrations

import (
	database "github.com/alhazmy20/Go-WebAPI/database"
	utils "github.com/alhazmy20/Go-WebAPI/utils"
)

func Create_Products_Table() {
	database.EnvFilePath = "../.env"
	db := database.OpenConnection()
	query := `CREATE TABLE IF NOT EXISTS products (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255),
		price FLOAT,
		quantity INT
	);`
	_, err := db.Exec(query)
	if err != nil {
		utils.HandleError(err)
	}
}

func Insert_Product() {
	database.EnvFilePath = "../.env"
	db := database.OpenConnection()
	query := `INSERT INTO PRODUCTS (name, price, quantity) VALUES ('Test_Product', 25.5, 500);`
	_, err := db.Exec(query)
	if err != nil {
		utils.HandleError(err)
	}
}

func Drop_Products_Table() {
	database.EnvFilePath = "../.env"
	db := database.OpenConnection()
	query := `DROP TABLE products;`
	_, err := db.Exec(query)
	if err != nil {
		utils.HandleError(err)
	}
}
