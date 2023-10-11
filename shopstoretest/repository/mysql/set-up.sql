create table users (
                       userID int not null auto_increment,
                       name varchar(255) not null,
                       hashed_password varchar(255) not null,
                       credit int unsigned default 0,
                       phone_number varchar(255) not null,
                       created_at  timestamp default current_timestamp,
                       PRIMARY KEY (userID)
);