CREATE TABLE users (
    id int primary key AUTO_INCREMENT,
    name varchar(255) not null,
    phone_number varchar(255) not null,
    hashed_password text not null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)