package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/trivo25/mina-view/backend/internal/util"
	"time"
)

func InsertService(service Service) error {
	conn, err := GetCon()
	if err != nil {
		return errors.New("An Error has occured")
	}
	defer conn.Close()
	service.ServiceInserted = int(time.Now().Unix())
	serviceString := service.ServiceName + time.Now().String()
	service.ServiceHash = util.GetHashOfService(serviceString)

	sqlStatement := `
			INSERT INTO services (service_name, logo_link, website_link, description, creator, hash, contact, inserted, accepted, islive)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, false, false);`
	_, err = conn.Exec(context.Background(), sqlStatement,
		&service.ServiceName,
		&service.ServiceLogo,
		&service.ServiceWebsite,
		&service.ServiceDescription,
		&service.ServiceCreator,
		&service.ServiceHash,
		&service.ServiceContact,
		&service.ServiceInserted,
	)
	if err != nil {
		return errors.New("An Error has occured")
	}

	var id int
	err = conn.QueryRow(context.Background(), "SELECT service_id FROM services WHERE hash = $1", service.ServiceHash).Scan(&id)
	/*	defer rows.Close()
		if err != nil {
			return errors.New("An Error has occured")
		}
		var id int
		for rows.Next() {
			fmt.Println(rows.Scan())
			err = rows.Scan(&id)
		}*/
	fmt.Println(service.ServiceHash)
	fmt.Println(id)

	return nil
}

func InsertCategory() {

}

func QueryCategories() ([]Category, error) {

	conn, err := GetCon()
	if err != nil {
		return nil, errors.New("An Error has occured")
	}
	defer conn.Close()

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

		err = conn.QueryRow(context.Background(),
			"SELECT COUNT(id) FROM services_categories WHERE category_id = $1;", c.CategoryID).Scan(&c.ServiceCount)

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

	conn, err := GetCon()

	if err != nil {
		return nil, errors.New("An Error has occured")
	}
	defer conn.Close()
	services := []Service{}
	rows, err := conn.Query(context.Background(),
		"SELECT services.service_id, services.service_name, services.logo_link, services.website_link, services.description, services.hash, services.inserted, services.creator, services.contact FROM services WHERE islive = true;")
	defer rows.Close()

	if err != nil {
		return nil, errors.New("An Error has occured")
	}
	for rows.Next() {
		s := Service{}
		err := rows.Scan(&s.ServiceID, &s.ServiceName, &s.ServiceLogo, &s.ServiceWebsite, &s.ServiceDescription, &s.ServiceHash, &s.ServiceInserted, &s.ServiceCreator, &s.ServiceContact)
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
