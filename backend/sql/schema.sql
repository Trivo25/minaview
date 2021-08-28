-- Schema for the MINAView Ecosystem overview+
-- I chose SQL for easy upgradablility and the option to go more into detail when it comes to things like voting, rating, etc.
-- might switch to elasticsearch if the need arises


CREATE TABLE services (
  sevice_id SERIAL,
  service_name varchar(255),
  logo_link varchar(255),
  website_link varchar(255),
  description text,
  creator varchar(255), -- TODO: add to Go structure
  PRIMARY KEY (sevice_id)
);

CREATE TABLE categories ( 
  category_id SERIAL,
  category_title varchar(255),
  category_key varchar(255),
  category_description text,
  PRIMARY KEY (category_id)
);

CREATE TABLE services_categories (
  id SERIAL,
  category_id int,
  service_id int,
  PRIMARY KEY (id)
);