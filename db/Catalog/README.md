# catalog
**catalog** is supposed to mimic a real-life example of a metaphor use-case. It represents a target database schema which stores metadata for tables and columns. The end-users are supposed to provide metadata regarding the systems/tables/columns that they want their target database to hold. An automated process can then read the normalised catalog structures and actually perform the create statements to sync up with the physical database.

## Table dictionary

`catalog.system`: Defines source systems that may contain multiple tables
`catalog.instance`: Defines instances of source systems
`catalog.schema`: Defines schemas that store tables
`catalog.table`: Defines tables of source systems that are stored in schemas
`catalog.column`: Defines attributes of columns within tables

## Folder structure
```bash
│   catalog.sql : "creates the catalog schema."
│
└───tables : "holds all table definitions within catalog"
        column.sql
        instance.sql
        schema.sql
        system.sql
        table.sql
```
