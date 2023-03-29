CREATE OR REPLACE FUNCTION change_purchase_status()
RETURNS TRIGGER AS
$$
BEGIN
    UPDATE "Purchases"
    SET "Status" = 0
    WHERE "ID" = OLD."PurchaseID";

    RETURN OLD;
END
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_sale_trigger
BEFORE DELETE ON "Sales"
FOR EACH ROW EXECUTE FUNCTION change_purchase_status();
