CREATE TABLE activities (
                            id BIGINT AUTO_INCREMENT,
                            email VARCHAR(255) NOT NULL,
                            title VARCHAR(255) NOT NULL,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
                            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
                            deleted_at TIMESTAMP NULL DEFAULT NULL,
                            PRIMARY KEY (id)
);