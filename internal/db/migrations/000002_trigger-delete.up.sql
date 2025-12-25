CREATE OR REPLACE FUNCTION handle_todo_items_soft_delete()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.status = 3 AND (OLD.status IS NULl OR OLD.status !=3) THEN
        NEW.deleted_at = now();
    ELSIF NEW.status != 3 AND OLD.status = 3 THEN
        NEW.deleted_at = NULL;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_soft_delete
    BEFORE UPDATE ON todo_items
    FOR EACH ROW
    WHEN ( OLD.status IS DISTINCT FROM NEW.status )
EXECUTE FUNCTION handle_todo_items_soft_delete();