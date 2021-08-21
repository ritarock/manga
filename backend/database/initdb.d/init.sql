CREATE DATABASE IF NOT EXISTS app;
USE app;
CREATE TABLE IF NOT EXISTS books (
	isbn char(20),
	title char(100),
	publisher char(100),
	pubdate char(100),
	cover char(100),
	author char(100),
	subject_code char(10)
);
