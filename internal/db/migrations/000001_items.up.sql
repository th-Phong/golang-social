CREATE TABLE todo_items (
    id          INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY ,
    title       VARCHAR(150) NOT NULL,
    description TEXT DEFAULT NULL,
    image       JSONB        DEFAULT NULL,
    status      SMALLINT     DEFAULT 1 NOT NULL CHECK ( status IN (1,2,3) ),
    created_at  TIMESTAMPTZ   DEFAULT now(),
    updated_at  TIMESTAMPTZ   DEFAULT now(),
    deleted_at  TIMESTAMPTZ   DEFAULT NULL
);

COMMENT ON COLUMN todo_items.status IS 'Status: 1-Doing, 2-Done, 3-Deleted';
COMMENT ON COLUMN todo_items.deleted_at IS 'Soft delete timestamp: NULL mean not deleted';

CREATE INDEX IF NOT EXISTS idx_todo_items_status ON todo_items(status);
CREATE INDEX IF NOT EXISTS idx_todo_items_deleted_at ON todo_items(deleted_at);

CREATE OR REPLACE FUNCTION update_items_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language plpgsql;

CREATE TRIGGER set_updated_at
BEFORE UPDATE ON todo_items
FOR EACH ROW
EXECUTE FUNCTION update_items_updated_at();
