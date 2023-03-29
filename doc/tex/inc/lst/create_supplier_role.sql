CREATE ROLE supplier WITH
    LOGIN
    NOSUPERUSER
    NOCREATEDB
    NOREPLICATION
    PASSWORD 'supplier'
    CONNECTION LIMIT -1;

GRANT SELECT ON public."Wines" TO supplier;
GRANT SELECT ON public."SupplierWines" TO supplier;
GRANT SELECT ON public."Sales" TO supplier;
GRANT INSERT ON public."Wines" TO supplier;
GRANT INSERT ON public."SupplierWines" TO supplier;
GRANT DELETE ON public."Wines" TO supplier;
GRANT DELETE ON public."SupplierWines" TO supplier;
GRANT UPDATE ON public."SupplierWines" TO supplier;
