DROP TABLE IF EXISTS books;
CREATE TABLE books (
  id         INT AUTO_INCREMENT NOT NULL,
  title      VARCHAR(128) NOT NULL,
  PRIMARY KEY (`id`)
);

INSERT INTO books
  (title)
VALUES
  ('Blue Train'),
  ('Giant Steps'),
  ('Jeru'),
  ('Sarah Vaughan');

DROP TABLE IF EXISTS book_page;
CREATE TABLE book_page (
  id         INT AUTO_INCREMENT NOT NULL,
  book_id    INT NOT NULL,
  page_num   INT NOT NULL,
  content    VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY (`book_id`, `page_num`),
	FOREIGN KEY (`book_id`) REFERENCES books(`id`) ON DELETE CASCADE
);

INSERT INTO book_page
  (book_id, page_num, content)
VALUES
  (1, 1, 'Lorem Ipsum is simply dummy text of the printing and typesetting industry.'),
  (1, 2, 'Lorem Ipsum has been the industry standard dummy text ever since the 1500s.'),
  (2, 1, 'When an unknown printer took a galley of type and scrambled it to make a type specimen book.'),
  (2, 2, 'When an unknown printer took a galley of type and scrambled it to make a type specimen book.'),
  (2, 3, 'When an unknown printer took a galley of type and scrambled it to make a type specimen book.'),
  (2, 4, 'When an unknown printer took a galley of type and scrambled it to make a type specimen book.'),
  (2, 5, 'When an unknown printer took a galley of type and scrambled it to make a type specimen book.'),
  (3, 1, 'When an unknown printer took a galley of type and scrambled it to make a type specimen book.'),
  (4, 1, 'It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.'),
  (4, 2, 'It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.'),
  (4, 3, 'It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.'),
  (4, 4, 'It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.'),
  (4, 5, 'It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.'),
  (4, 6, 'It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.'),
  (4, 7, 'It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.'),
  (4, 8, 'It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged.');
