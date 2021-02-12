CREATE TABLE catalog.table
(
    id              INT GENERATED ALWAYS AS IDENTITY,
    name            VARCHAR(128) NOT NULL,
    description     VARCHAR(1000) NULL,
    schema_id       INT NOT NULL,
    instance_id     INT NOT NULL,
    created_on      TIMESTAMP NOT NULL CONSTRAINT df_created_on DEFAULT CURRENT_TIMESTAMP,
    is_deleted      INT NOT NULL CONSTRAINT df_created_on DEFAULT 0,
    deleted_on      TIMESTAMP NULL,

    CONSTRAINT pk_table PRIMARY KEY (id),
    CONSTRAINT fk_table_schema FOREIGN KEY (schema_id) REFERENCES catalog.schema(id),
    CONSTRAINT fk_table_instance FOREIGN KEY (instance_id) REFERENCES catalog.instance(id),
    CONSTRAINT uc_table_name_schema UNIQUE (name, schema_id, deleted_on)
);