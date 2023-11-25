package productsController

import (
	"database/sql"
	"encoding/json"
	"net/http"

	database "github.com/alhazmy20/Go-WebAPI/database"
	models "github.com/alhazmy20/Go-WebAPI/models"
	utils "github.com/alhazmy20/Go-WebAPI/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var product models.Product
var db *sql.DB

func Index(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	db = database.OpenConnection()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		utils.HandleError(err)
		utils.SendJSONResponse(w, http.StatusInternalServerError, utils.ApiResponse{Message: "Error while querying the database"})
		return
	}
	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity); err != nil {
			utils.HandleError(err)
			utils.SendJSONResponse(w, http.StatusInternalServerError, utils.ApiResponse{Message: "Error while scanning the rows"})
			return
		}
		products = append(products, product)
	}
	jsonResponse, err := json.Marshal(products)
	if err != nil {
		utils.HandleError(err)
		return
	}
	utils.SendJSON(w, http.StatusOK, jsonResponse)
}

func Store(w http.ResponseWriter, r *http.Request) {
	db := database.OpenConnection()
	defer db.Close()
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		utils.HandleError(err)
	}
	res, err := db.Exec("INSERT INTO PRODUCTS (name, price, quantity) VALUES (?, ?, ?)", product.Name, product.Price, product.Quantity)
	if err != nil {
		utils.HandleError(err)
	}
	if rows, _ := res.RowsAffected(); rows > 0 {
		utils.SendJSONResponse(w, http.StatusOK, utils.ApiResponse{Message: "Product Inserted Successfully"})
	}
}
func Show(w http.ResponseWriter, r *http.Request) {
	db = database.OpenConnection()
	defer db.Close()
	id := mux.Vars(r)["id"]
	row := db.QueryRow("SELECT * FROM products WHERE id = ?", id)
	if err := row.Scan(&product.Id, &product.Name, &product.Price, &product.Quantity); err != nil {
		utils.SendJSONResponse(w, http.StatusNotFound, utils.ApiResponse{Message: "Product Not Found"})
	} else {
		jsonResponse, err := json.Marshal(product)
		if err != nil {
			utils.HandleError(err)
		}
		utils.SendJSON(w, http.StatusOK, jsonResponse)
	}
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	db = database.OpenConnection()
	defer db.Close()
	id := mux.Vars(r)["id"]
	res, err := db.Exec("DELETE FROM PRODUCTS WHERE id = ?", id)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusNotFound, utils.ApiResponse{Message: "Product Not Found"})
	} else {
		if rows, _ := res.RowsAffected(); rows > 0 {
			utils.SendJSONResponse(w, http.StatusOK, utils.ApiResponse{Message: "Product Deleted Successfully"})
		}
	}
}

func Update(w http.ResponseWriter, r *http.Request) {
	db = database.OpenConnection()
	defer db.Close()
	id := mux.Vars(r)["id"]
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&product); err != nil {
		utils.HandleError(err)
	}
	res, err := db.Exec("UPDATE PRODUCTS SET name = ?, price = ?, quantity = ? WHERE id = ?", product.Name, product.Price, product.Quantity, id)
	if err != nil {
		utils.HandleError(err)
	} else {
		if rows, _ := res.RowsAffected(); rows > 0 {
			utils.SendJSONResponse(w, http.StatusOK, utils.ApiResponse{Message: "Product Updated Successfully"})
		} else {
			utils.SendJSONResponse(w, http.StatusNotFound, utils.ApiResponse{Message: "Product Not Found"})
		}
	}
}
