CREATE TABLE catalog.column
(
    id              INT GENERATED ALWAYS AS IDENTITY,
    name            VARCHAR(128) NOT NULL,
    description     VARCHAR(1000) NULL,
    ordinal         SMALLINT NOT NULL,
    data_type       VARCHAR(15) NOT NULL,
    length          SMALLINT NULL,
    precision       SMALLINT NULL,
    scale           SMALLINT NULL,
    nullable        BOOLEAN NOT NULL CONSTRAINT df_nullable DEFAULT FALSE,
    column_default  VARCHAR(80) NULL,
    table_id        INT NOT NULL,
    created_on      TIMESTAMP NOT NULL CONSTRAINT df_created_on DEFAULT CURRENT_TIMESTAMP,
    is_deleted      INT NOT NULL CONSTRAINT df_created_on DEFAULT 0,
    deleted_on      TIMESTAMP NULL,

    CONSTRAINT pk_column PRIMARY KEY (id),
    CONSTRAINT fk_column_table FOREIGN KEY (table_id) REFERENCES catalog.table(id),
    CONSTRAINT uc_column_name_table UNIQUE (name, table_id, deleted_on)
);