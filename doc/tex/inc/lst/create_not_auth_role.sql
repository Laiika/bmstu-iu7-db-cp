CREATE ROLE not_auth_user WITH
    LOGIN
    NOSUPERUSER
    NOCREATEDB
    NOREPLICATION
    PASSWORD 'not_auth_user'
    CONNECTION LIMIT -1;

GRANT SELECT ON public."Wines" TO not_auth_user;
GRANT SELECT ON public."SupplierWines" TO not_auth_user;

GRANT INSERT ON public."Users" TO not_auth_user;
