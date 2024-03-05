CREATE TABLE todos (
                       id BIGINT AUTO_INCREMENT,
                       activity_group_id BIGINT NOT NULL,
                       title VARCHAR(255) NOT NULL,
                       is_active BOOLEAN DEFAULT TRUE NOT NULL,
                       priority VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
                       deleted_at TIMESTAMP NULL DEFAULT NULL,
                       PRIMARY KEY (id),
                       FOREIGN KEY (activity_group_id) REFERENCES activities(id)
);