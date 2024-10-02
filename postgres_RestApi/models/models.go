package models

import (
	"log"
	"postgres_RestApi/config"
	// "golang.org/x/tools/go/analysis/passes/defers"
)

type Barang struct {
	Id          int64  `json:"id"`
	Nama_barang string `json:"nama_barang"`
	Harga       int    `json:"harga"`
	Stok        int    `json:"stok"`
}

func GetAllItem() ([]Barang, error) {

	db := config.CreateConnection()

	defer db.Close()

	// sql query
	sqlStatement := "SELECT * FROM barang"

	// query untuk database
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal("ini err", err)
	}

	defer rows.Close()
	// variabel untuk semua data
	var barangS []Barang

	for rows.Next() {

		var barang Barang
		err = rows.Scan(&barang.Id, &barang.Nama_barang, &barang.Harga, &barang.Stok)

		if err != nil {
			log.Fatal("ini eror", err)
		}

		barangS = append(barangS, barang)

	}

	return barangS, err

}

func GetOneItem(id int64) (Barang, error) {
	// coneksi ke database
	db := config.CreateConnection()

	// menutup koneksi ketika akhir program
	defer db.Close()

	// sql query
	sqlStatement := "SELECT * FROM barang WHERE id = $1"

	var barang Barang

	// implementasi sql statement
	row := db.QueryRow(sqlStatement, id)

	// cek err
	err := row.Scan(&barang.Id, &barang.Nama_barang, &barang.Harga, &barang.Stok)
	if err != nil {
		log.Fatal(err)
	}

	// mengembalikan data barang dan err yang dicari
	return barang, err
}

func AddItem(barang Barang) int64 {
	// coneksi database
	db := config.CreateConnection()
	defer db.Close()

	// sql statement
	sqlStatement := "INSERT INTO barang (nama_barang, harga, stok ) VALUES ($1,$2,$3) RETURNING id"

	var id int64

	// eksekusi query sql
	err := db.QueryRow(sqlStatement, barang.Nama_barang, barang.Harga, barang.Stok).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	return id

}

func UpdateItem(id int64, barang Barang) int64 {
	// konseksi db
	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := "UPDATE barang SET nama_barang=$2, harga=$3, stok=$4 WHERE id=$1"

	res, err := db.Exec(sqlStatement, id, barang.Nama_barang, barang.Harga, barang.Stok)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	return rowsAffected
}

func DeleteOneItem(id int64) int64 {
	// koneksi database
	db := config.CreateConnection()

	defer db.Close()

	sqlStatement := "DELETE FROM barang WHERE id = $1"

	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatal(err)
	}

	return rowsAffected

}
