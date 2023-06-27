package controller

import (
	"context"
	"encoding/json"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/go-chi/render"
	"github.com/jackc/pgx/v4/pgxpool"
	"io"
	"mobilePhoneEdu/internal/model"
	"net/http"
)

var (
	DB *pgxpool.Pool
)

func DeletePhone(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var phones model.Phone
	err = json.Unmarshal(body, &phones)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid unmarshall"))
		return
	}
	err = DeleteFromStorage(phones)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed database"))
	}
	w.WriteHeader(200)

}

func DeleteFromStorage(phones model.Phone) error {
	query := `delete from phone where id = $1`
	_, err := DB.Exec(context.Background(), query, phones.Id)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

/////////////////////////////

func UpdatePhone(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var phones model.Phone

	err = json.Unmarshal(body, &phones)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid unmarshall"))
		return
	}
	err = UpdatePhoneFromStorage(phones)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed database"))
		return
	}
	w.WriteHeader(200)
}

func UpdatePhoneFromStorage(phones model.Phone) error {

	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	sqUpdate := psql.Update("phone")
	//query := `update phone `

	if phones.Brand != "" {
		//query += `set brand = `
		//query += phones.Brand
		sqUpdate = sqUpdate.Set("brand", phones.Brand)
	}
	if phones.Model != "" {
		sqUpdate = sqUpdate.Set("model", phones.Model)
	}
	if phones.ModelYear != 0 {
		sqUpdate = sqUpdate.Set("modelYear", phones.ModelYear)
	}
	if phones.Diagonal != 0.0 {
		sqUpdate = sqUpdate.Set("diagonal", phones.Diagonal)
	}
	if phones.MemoryStorage != 0 {
		sqUpdate = sqUpdate.Set("memoryStorage", phones.MemoryStorage)
	}
	if phones.Ram != 0 {
		sqUpdate = sqUpdate.Set("ram", phones.Ram)
	}
	if phones.Weight != 0 {
		sqUpdate = sqUpdate.Set("weight", phones.Weight)
	}
	if phones.BatteryCapacity != 0 {
		sqUpdate = sqUpdate.Set("batteryCapacity", phones.BatteryCapacity)
	}
	if phones.Color != "" {
		sqUpdate = sqUpdate.Set("color", phones.Color)
	}
	if phones.Price != 0 {
		sqUpdate = sqUpdate.Set("price", phones.Price)
	}
	sqUpdate = sqUpdate.Where(sq.Eq{"id": phones.Id})
	sqlQuery, args, err := sqUpdate.ToSql()
	if err != nil {
		fmt.Println(err)
		return err
	}
	//fmt.Println(query)
	fmt.Println(sqlQuery)
	_, err = DB.Exec(context.Background(), sqlQuery, args...)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//////////////////////////

func CreatePhone(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	var phones model.Phone
	err = json.Unmarshal(body, &phones)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid unmarshall"))
		return
	}
	err = CreatePhoneInStorage(phones)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed database"))
		return
	}
	w.WriteHeader(200)
}

func CreatePhoneInStorage(phones model.Phone) error {
	query := `insert into phone (brand, model, modelYear, diagonal, memoryStorage, ram, weight, batteryCapacity, color, price) values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := DB.Exec(context.Background(), query, phones.Brand, phones.Model, phones.ModelYear, phones.Diagonal, phones.MemoryStorage, phones.Ram, phones.Weight, phones.BatteryCapacity, phones.Color, phones.Price)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//////////////////////////

func GetPhone(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json; charset=utf-8")
	results, err := GetPhoneFromStorage()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("failed database"))
		return
	}

	render.JSON(w, r, results)
	//w.WriteHeader(200)
}

func GetPhoneFromStorage() ([]model.Phone, error) {
	query := `select id, brand, model, modelYear, diagonal, memoryStorage, ram, weight, batteryCapacity, color, price from phone`
	rows, err := DB.Query(context.Background(), query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var results []model.Phone
	for rows.Next() {
		var temp model.Phone
		err := rows.Scan(&temp.Id, &temp.Brand, &temp.Model, &temp.ModelYear, &temp.Diagonal, &temp.MemoryStorage, &temp.Ram, &temp.Weight, &temp.BatteryCapacity, &temp.Color, &temp.Price)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		results = append(results, temp)
	}
	return results, nil
}
