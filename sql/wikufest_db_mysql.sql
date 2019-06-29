CREATE TABLE IF NOT EXISTS user (
    id int not null AUTO_INCREMENT,
    `password`  varchar(255),
    fullname  varchar(255),
    email  varchar(255) NOT NULL UNIQUE,
    reg_date DATETIME,
    upd_date DATETIME,
    is_active boolean,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS user_detail (
    user_id int,
    class_name  varchar(255),
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE IF NOT EXISTS course (
    id int not null AUTO_INCREMENT,
    is_private_class boolean,
    course_name varchar(255),
    description text,
    additional_link text,
    is_active boolean,
    reg_date DATETIME,
    upd_date DATETIME,
    PRIMARY KEY(id)
);

CREATE TABLE IF NOT EXISTS user_coupon (
    id int not null AUTO_INCREMENT,
    user_id integer,
    coupon_name varchar(50),
    coupon_code varchar(100),
    is_redeemed boolean,
    is_active boolean,
    reg_date DATETIME,
    upd_date DATETIME,
    PRIMARY KEY(id),
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE IF NOT EXISTS course_session (
    id int not null AUTO_INCREMENT,
    course_id int,
    room_name varchar(255),
    room_location varchar(255),
    quota int,
    available_quota int,
    start_time DATETIME,
    end_time DATETIME,
    is_active boolean,
    reg_date DATETIME,
    upd_date DATETIME,
    PRIMARY KEY(id),
    FOREIGN KEY (course_id) REFERENCES course(id)
);

CREATE TABLE IF NOT EXISTS course_session_enrollment (
    course_session_id int,
    user_id int,
    reg_date DATETIME,
    FOREIGN KEY (course_session_id) REFERENCES course_session(id),
    FOREIGN KEY (user_id) REFERENCES user(id)
);

CREATE TABLE IF NOT EXISTS course_instructor (
    course_id integer,
    instructor_id integer,
    reg_date DATETIME,
    FOREIGN KEY (course_id) REFERENCES course(id),
    FOREIGN KEY (instructor_id) REFERENCES user(id)
);