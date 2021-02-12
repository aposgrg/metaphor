CREATE TABLE catalog.instance 
(
    id              INT GENERATED ALWAYS AS IDENTITY,
    name            VARCHAR(128) NOT NULL,
    description     VARCHAR(1000) NULL,
    system_id       INT NOT NULL,
    created_on      TIMESTAMP NOT NULL CONSTRAINT df_created_on DEFAULT CURRENT_TIMESTAMP,
    is_deleted      INT NOT NULL CONSTRAINT df_created_on DEFAULT 0,
    deleted_on      TIMESTAMP NULL,

    CONSTRAINT pk_instance PRIMARY KEY (id),
    CONSTRAINT fk_instance_system FOREIGN KEY (system_id) REFERENCES catalog.system(id),
    CONSTRAINT uc_instance_name UNIQUE (name, system_id, deleted_on)
);