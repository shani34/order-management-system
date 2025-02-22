-- schema.sql
CREATE TABLE orders (
    order_id TEXT PRIMARY KEY,
    user_id INTEGER NOT NULL,
    item_ids TEXT NOT NULL, -- JSON array as string
    total_amount REAL NOT NULL,
    status TEXT NOT NULL DEFAULT 'Pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
