CREATE TABLE student
(
    student_id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(225) NOT NULL,
    email VARCHAR(225) NOT NULL ,
    password BINARY(60) NOT NULL,
    address VARCHAR(225) NOT NULL,
    PRIMARY KEY(student_id)
)ENGINE = InnoDB;
ALTER TABLE student ADD UNIQUE (email);