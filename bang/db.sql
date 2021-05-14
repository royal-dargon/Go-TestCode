CREATE DATABASE IF NOT EXISTS bang;

Use bang;

Create Table users (
    user_id BIGINT NOT NULL AUTO_INCREMENT,
    user_name VARCHAR(30) NOT NULL,
    user_password VARCHAR(50) NOT NULL,
    user_school VARCHAR(50) NOT NULL,
    user_college VARCHAR(50) NOT NULL,
    really_name VARCHAR(50)  NOT NULL,
    user_picture VARCHAR(80)  NULL,
    PRIMARY KEY (user_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

Create TABLE contest (
    contest_id BIGINT NOT NULL AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    contest_info TEXT NULL,
    contest_people VARCHAR(50) NOT NULL,
    createtime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    contest_url VARCHAR(80) NULL,
    people_college VARCHAR(50) NOT NULL,
    PRIMARY KEY(contest_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

Create TABLE review (
    review_id BIGINT NOT NULL AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    require_id BIGINT NOT NULL,
    content TEXT NOT NULL,
    date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (review_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

Create TABLE user_require (
    id BIGINT NOT NULL AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    require_id BIGINT NOT NULL,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

Create TABLE requires (
    require_id BIGINT NOT NULL AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    require_title VARCHAR(50) NOT NULL,
    require_number BIGINT NOT NULL,
    required_need VARCHAR(80) NULL,
    require_info TEXT NOT NULL,
    require_kind VARCHAR(30) NOT NULL,
    createtime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(require_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
