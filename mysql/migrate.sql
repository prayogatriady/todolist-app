CREATE TABLE users (
    id BIGINT NOT NULL AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    INDEX (id),
    PRIMARY KEY (id)
);

CREATE TABLE lists (
    id BIGINT NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    user_id BIGINT NOT null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    INDEX (user_id),
    PRIMARY KEY (id),
    foreign key (user_id) references users (id) ON DELETE RESTRICT on update CASCADE
);

CREATE TABLE tasks (
    id BIGINT NOT NULL AUTO_INCREMENT,
    description TEXT NOT NULL,
    completed BOOLEAN DEFAULT FALSE,
    list_id BIGINT NOT null,
    user_id BIGINT NOT null,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL,
    INDEX (list_id),
    PRIMARY KEY (id),
    foreign key (list_id) references lists (id) ON DELETE RESTRICT ON update CASCADE,
    foreign key (user_id) references users (id) ON DELETE RESTRICT on update CASCADE
);