-- Drop trigger if exist
DROP TRIGGER IF EXISTS set_updated_at on todo_items;

-- Drop function
DROP FUNCTION IF EXISTS update_items_updated_at();

-- Drop index (option)
DROP INDEX IF EXISTS idx_todo_items_status;
DROP INDEX IF EXISTS idx_todo_items_deleted_at;

-- Drop table
DROP TABLE IF EXISTS todo_items;
