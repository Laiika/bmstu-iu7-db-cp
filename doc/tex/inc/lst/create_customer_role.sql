CREATE ROLE customer WITH
    LOGIN
    NOSUPERUSER
    NOCREATEDB
    NOREPLICATION
    PASSWORD 'customer'
    CONNECTION LIMIT -1;

GRANT SELECT ON public."Wines" TO customer;
GRANT SELECT ON public."SupplierWines" TO customer;
GRANT SELECT ON public."Suppliers" TO customer;
GRANT SELECT ON public."Purchases" TO customer;
GRANT SELECT ON public."BonusCards" TO customer;

GRANT INSERT ON public."BonusCards" TO customer;
GRANT INSERT ON public."Purchases" TO customer;
GRANT INSERT ON public."Sales" TO customer;

GRANT DELETE ON public."Sales" TO customer;

GRANT UPDATE ON public."Purchases" TO customer;
