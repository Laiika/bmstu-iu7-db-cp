CREATE ROLE "admin" WITH
    LOGIN
    NOSUPERUSER
    NOCREATEDB
    NOREPLICATION
    PASSWORD 'admin'
    CONNECTION LIMIT -1;

GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA PUBLIC TO "admin";
