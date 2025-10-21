ALTER TABLE authors
ADD COLUMN name VARCHAR(255) NOT NULL;

ALTER TABLE authors
DROP COLUMN birth_year;


---- create above / drop below ----

DROP TABLE authors;

