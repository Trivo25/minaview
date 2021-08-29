package db

import (
	"context"
	"errors"
)

func QueryCategories() ([]Category, error) {

	conn := GetCon()
	categories := []Category{}
	rows, err := conn.Query(context.Background(), "SELECT categories.category_id, categories.category_title, categories.category_description, categories.category_key FROM categories;")
	defer rows.Close()
	if err != nil {
		return nil, errors.New("An Error has occured")
	}
	for rows.Next() {
		c := Category{}
		err = rows.Scan(&c.CategoryID, &c.CategoryTitle, &c.CategoryDescription, &c.CategoryKey)
		if err != nil {
			return nil, errors.New("An Error has occured")
		}

		categories = append(categories, c)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.New("An Error has occured")
	}
	return categories, nil
}

func QueryServices() ([]Service, error) {
	conn := GetCon()
	services := []Service{}
	rows, err := conn.Query(context.Background(),
		"SELECT services.service_id, services.service_name, services.logo_link, services.website_link, services.description, services.hash, services.inserted, services.creator FROM services WHERE islive = true;")
	defer rows.Close()

	if err != nil {
		return nil, errors.New("An Error has occured")
	}
	for rows.Next() {
		s := Service{}
		err := rows.Scan(&s.ServiceID, &s.ServiceName, &s.ServiceLogo, &s.ServiceWebsite, &s.ServiceDescription, &s.ServiceHash, &s.ServiceInserted, &s.ServiceCreator)
		if err != nil {
			return nil, errors.New("An Error has occured")
		}
		catRows, err := conn.Query(context.Background(), "SELECT categories.category_key FROM categories INNER JOIN services_categories ON services_categories.category_id = categories.category_id WHERE service_id = $1;", s.ServiceID)

		if err != nil {
			return nil, errors.New("An Error has occured")
		}
		categories := []string{}
		for catRows.Next() {
			key := ""
			err = catRows.Scan(&key)

			if err != nil {
				return nil, errors.New("An Error has occured")
			}
			categories = append(categories, key)

		}
		s.CategoryKeys = categories
		services = append(services, s)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.New("An Error has occured")
	}

	return services, nil
}
