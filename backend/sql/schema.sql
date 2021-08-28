-- Schema for the MINAView Ecosystem overview+
-- I chose SQL for easy upgradablility and the option to go more into detail when it comes to things like voting, rating, etc.
-- might switch to elasticsearch if the need arises



CREATE TABLE services (
                          service_id SERIAL,
                          service_name varchar(255) DEFAULT '',
                          logo_link varchar(255) DEFAULT '',
                          website_link varchar(255) DEFAULT '',
                          description text DEFAULT '',
                          creator varchar(255) DEFAULT '',
                          hash varchar(255) DEFAULT '',
                          inserted bigint,
                          accepted boolean DEFAULT FALSE,
                          islive boolean DEFAULT FALSE,
                          PRIMARY KEY (sevice_id)
);

CREATE TABLE categories (
                            category_id SERIAL,
                            category_title varchar(255) DEFAULT '',
                            category_key varchar(255) DEFAULT '',
                            category_description text DEFAULT '',
                            PRIMARY KEY (category_id)
);

CREATE TABLE services_categories (
                                     id SERIAL,
                                     category_id int,
                                     service_id int,
                                     PRIMARY KEY (id)
);