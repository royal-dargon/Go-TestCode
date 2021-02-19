create database if not exists library;

use library;

create table users (
	user_id BIGINT primary key not null AUTO_INCREMENT,
	user_name varchar(30) not null,
	motto varchar(50) not null,
	password varchar(16) not null,
	user_picture varchar(50) not null
);

create table Books (
	books_id BIGINT primary key not null AUTO_INCREMENT,
	books_name varchar(50) not null,
	books_picture varchar(50) not null,
	author varchar(20) not null,
	books_content text not null,
	click  BIGINT not null,
	kinds_id BIGINT not null
);

create table digest (
	Digest_id BIGINT primary key not null AUTO_INCREMENT,
	Digest_name varchar(20) not null,
	Digest_content text not null,
	Digest_title varchar(30) not null,
	Digest_idea text not null,
	books_id BIGINT not null,
        user_id BIGINT not null,
        DigestKinds_id BIGINT not null,
	locken BOOLEAN not null,
	time timestamp not null default current_timestamp on update current_timestamp
);

create table DigestKinds (
	digestkinds_id BIGINT primary key not null AUTO_INCREMENT,
	digestkinds_name char(20) not null
);

create table BooksKinds (
        BooksKinds_id BIGINT not null,
	BooksKinds_picture char(50) not null,
	BooksKinds_number char(8) not null
);

create table UtoB (
	id1 char(8) not null,
	user_id char(6) not null,
	Books_id char(8) not null
);


	
