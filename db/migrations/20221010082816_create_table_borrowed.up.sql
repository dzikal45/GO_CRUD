CREATE TABLE borrowed
(
    borrowed_id INT NOT NULL AUTO_INCREMENT,
    student_id INT,
    book_id INT,
    status_request VARCHAR(225) NOT NULL,
    book_name VARCHAR(225) NOT NULL,
    due_date DATE NOT NULL ,
    return_date DATE ,
    PRIMARY KEY(borrowed_id),
    FOREIGN KEY(student_id) REFERENCES student(student_id)   ON DELETE CASCADE,
    FOREIGN KEY(book_id) REFERENCES book(book_id)   ON DELETE CASCADE
)ENGINE = InnoDB;