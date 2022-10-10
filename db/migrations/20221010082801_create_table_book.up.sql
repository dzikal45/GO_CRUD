CREATE TABLE book
(
    book_id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(225) NOT NULL,
    available INT NOT NULL CHECK(available IN (1,0)),
    PRIMARY KEY(book_id)
)ENGINE = InnoDB;
