-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
  id         INT NOT NULL AUTO_INCREMENT,
  name       varchar(255) NOT NULL,
  email       varchar(255) NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (id),
  CONSTRAINT uc_email UNIQUE (email)
) ENGINE=InnoDB;



-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS users;
