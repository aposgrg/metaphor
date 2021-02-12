CREATE TABLE catalog.system 
(
    id              INT GENERATED ALWAYS AS IDENTITY,
    name            VARCHAR(128) NOT NULL,
    description     VARCHAR(1000) NULL,
    created_on      TIMESTAMP NOT NULL CONSTRAINT df_created_on DEFAULT CURRENT_TIMESTAMP,
    is_deleted      INT NOT NULL CONSTRAINT df_created_on DEFAULT 0,
    deleted_on      TIMESTAMP NULL,

    CONSTRAINT pk_system PRIMARY KEY (id),
    CONSTRAINT uc_system_name UNIQUE (name, deleted_on)
);