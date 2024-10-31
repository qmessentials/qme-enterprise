CREATE TABLE schema_version (
    version_number VARCHAR (20) PRIMARY KEY,
    applied_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE application_user (
    user_id SERIAL PRIMARY KEY,
    user_name VARCHAR (20) NOT NULL,
    given_name VARCHAR (100) NOT NULL,
    family_name VARCHAR (100) NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    UNIQUE (user_name)
);

CREATE TABLE application_role (
    role_id SERIAL PRIMARY KEY,
    role_name VARCHAR (100) NOT NULL,
    UNIQUE (role_name)
);

CREATE TABLE user_role (
    user_id INT REFERENCES application_user (user_id),
    role_id INT REFERENCES application_role (role_id),
    PRIMARY KEY (user_id, role_id)
);

DO $$
DECLARE 
    admin_role_id INTEGER;
    admin_user_id INTEGER;

BEGIN 

    INSERT INTO application_role (role_name)
    VALUES ('Administrator') 
    RETURNING role_id INTO admin_role_id;

    INSERT INTO application_user (user_name, given_name)
    VALUES ('admin', 'Default Admin')
    RETURNING user_id INTO admin_user_id;

    INSERT INTO user_role (user_id, role_id)
    VALUES (admin_user_id, admin_role_id);

END $$;

INSERT INTO application_role (role_name)
VALUES ('Quality Specialist');

INSERT INTO application_role (role_name)
VALUES ('Operations Lead');

INSERT INTO application_role (role_name)
VALUES ('Tester');

INSERT INTO application_role (role_name)
VALUES ('Manager');

BEGIN TRANSACTION;
    
    DELETE FROM schema_version;

    INSERT INTO schema_version (version_number)
    VALUES ('0.1.0');

COMMIT;

