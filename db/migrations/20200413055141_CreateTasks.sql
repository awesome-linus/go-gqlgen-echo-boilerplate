-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE tasks (
  id         INT NOT NULL AUTO_INCREMENT,
  title      varchar(255) NOT NULL,
  notes      text NOT NULL,
  completed  tinyint(1) NOT NULL DEFAULT 0,
  due        DATE NOT NULL,
  user_id    INT NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL,
  deleted_at DATETIME NULL DEFAULT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (user_id) 
  REFERENCES users (id)
  ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE=InnoDB;



-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS tasks;
