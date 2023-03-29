CREATE OR REPLACE FUNCTION delete_user()
RETURNS TRIGGER AS
$$
DECLARE
    customer_rec RECORD;
BEGIN
    IF OLD."Role" = 'customer' THEN
        SELECT * INTO customer_rec
        FROM "Customers"
        WHERE "ID" = OLD."RoleID";

        IF customer_rec."BonusCardID" IS NOT NULL THEN
            DELETE FROM "BonusCards"
            WHERE "ID" = customer_rec."BonusCardID";

            UPDATE "Customers"
            SET "BonusCardID" = NULL
            WHERE "ID" = OLD."RoleID";

        END IF;
    END IF;

    RETURN OLD;
END
$$ LANGUAGE plpgsql;

CREATE TRIGGER delete_user_trigger
BEFORE DELETE ON "Users"
FOR EACH ROW EXECUTE FUNCTION delete_user();
