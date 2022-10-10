CREATE TABLE borrowed
(
    borrowed_id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(225) NOT NULL,
    status_request VARCHAR(225) NOT NULL,
    book_name VARCHAR(225) NOT NULL,
    return_date TIMESTAMP,
    student_id INT,
    PRIMARY KEY(borrowed_id),
    FOREIGN KEY(student_id) REFERENCES student(student_id)
)ENGINE = InnoDB;