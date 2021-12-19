CREATE DATABASE IF NOT EXISTS app;
USE app;
CREATE TABLE IF NOT EXISTS books (
  isbn CHAR(20),
  title CHAR(150),
  publisher CHAR(150),
  pubdate CHAR(150),
  cover CHAR(150),
  author CHAR(150),
  subject_code CHAR(10)
)
