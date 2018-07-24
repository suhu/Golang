package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type SiteData struct {
	SiteId   int
	SiteName string
	Country  string
}

func main() {

	sites, err := GetSites()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sites[0].SiteName)
	fmt.Println(sites[1].SiteName)

}

func GetSites() ([]*SiteData, error) {

	db := GetDbConnection()

	rows, err := db.Query("select * from websitedata.website where sitename like '%site%'")

	sites := make([]*SiteData, 0)

	if err != nil {
		return nil, err
	}

	for rows.Next() {

		site := new(SiteData)
		err := rows.Scan(&site.SiteId, &site.SiteName, &site.Country)
		if err != nil {
			return nil, err
		}
		sites = append(sites, site)
	}

	defer db.Close()

	return sites, nil
}

func GetDbConnection() *sql.DB {

	db, err := sql.Open("mysql", "root:suhumar123@tcp(127.0.0.1:3306)/websitedata")

	pingErr := db.Ping()

	if pingErr != nil {
		fmt.Println("Ping failed")
		panic(pingErr.Error())
	}

	// if there is an error opening the connection, handle it
	if err != nil {

		fmt.Println("error")
		panic(err.Error())
	}

	if err == nil {
		fmt.Println("Making a connection with mysql")
	}

	return db

}
