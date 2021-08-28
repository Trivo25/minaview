package db

import (
	"context"
	"fmt"
)

func QueryServices() {
	conn := GetCon()
	services := []Service{}
	rows, err := conn.Query(context.Background(),
		"select services.service_id, services.service_name, services.logo_link, services.website_link, services.description from services")
	defer rows.Close()
	for rows.Next() {
		s := Service{}
		err = rows.Scan(&s.ServiceID, &s.ServiceName, &s.ServiceLogo, &s.ServiceWebsite, &s.ServiceDescription)
		if err != nil {
			panic(err)
		}

		categoryRows, er := conn.Query(context.Background(),
			"SELECT categories.category_id, categories.categry_title, categories.category_description, categories.categories_key"+
				"FROM categories"+
				"INNER JOIN services_categories ON services_categories.category_id = categories.category_id WHERE service_id = $1", &s.ServiceID)
		if er != nil {
			panic(err)
		}
		categories := []Category{}
		for categoryRows.Next() {
			c := Category{}
			err = rows.Scan(&c.CategoryID, &c.CategoryTitle, &c.CategoryDescription, &c.CategoryKey)
			if err != nil {
				panic(err)
			}
			categories = append(categories, c)
		}
		s.Categories = categories
		services = append(services, s)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println(services)
}
