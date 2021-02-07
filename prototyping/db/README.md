## Intro

`metaphor\prototyping\db` is focused on exploring ways of interacting with databases to address needs and use cases identified during the development of metaphor.

As of now, the module is divided in two sections:
- `db_docker`: contains the Docker configs
- `db_client`: contains illustrative and exploratory client code

## Database setup with Docker: `db\db_docker`

We are going to use Docker Compose, which is a tool for defining and running multi-container Docker applications. It uses YAML configuration files to start and manage application's services.

1. Get Docker: 
    - [install Docker Compose](https://docs.docker.com/compose/install/)
    - [or install Docker Desktop (Mac and Windows), which includes Compose.](https://docs.docker.com/desktop/)
2. Review the database configuration:
   - Open `metaphor\prototyping\db\db_docker\docker-compose.yml`
3. Start the services:
    ```bash
    $ cd \metaphor\prototyping\db\db_docker
    $ docker-compose up
    ```
    This will pull or build the images defined in `docker-compose.yml`. At this point, it will download just the postgres image (it needs to do this only once) and then it will start the database. 

    You can start a second terminal session to test the database, for example:
    ```bash
    $ cd \metaphor\prototyping\db\db_docker
    $ docker-compose run db bash
    $ psql --host=db --username=db_user --dbname=demo_db
    $ \d
    $ create table foo (id int);
    $ \d
    $ drop table foo
    ```

## Interact with the database: `db\db_client`

This module serves as a little playground that can be used to test the database set up and started in the previous step.

1. Get dependencies:

    ```bash
    $ cd \metaphor\prototyping\db\db_client
    $ go get "github.com/jmoiron/sqlx"
    $ go get "github.com/lib/pq"
    ```

2. Compile and run:
    ```bash
    $ cd \metaphor\prototyping\db\db_client
    $ go run main.go
    ```
    This _should_ print messages on what was attempted along with results.
