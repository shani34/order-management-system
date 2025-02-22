# order-management-system
The database (orders.db) is created automatically when the application starts.
However, you can manually create the table by running:

sh
Copy
Edit
sqlite3 orders.db
Then execute:

sql
Copy
Edit
CREATE TABLE orders (
    order_id TEXT PRIMARY KEY,
    user_id INTEGER NOT NULL,
    item_ids TEXT NOT NULL, -- JSON array as string
    total_amount REAL NOT NULL,
    status TEXT NOT NULL DEFAULT 'Pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);
 API Endpoints
1. Create Order
Request:

sh
Copy
Edit
curl -X POST "http://localhost:8080/orders" \
-H "Content-Type: application/json" \
-d '{
  "user_id": 1,
  "order_id": "ORD123",
  "item_ids": [101, 102, 103],
  "total_amount": 250.50
}'
Response:

json
Copy
Edit
{
  "message": "Order received",
  "order_id": "ORD123"
}
 2. Check Order Status
Request:

sh
Copy
Edit
curl -X GET "http://localhost:8080/orders/ORD123/status"
Response:

json
Copy
Edit
{
  "order_id": "ORD123",
  "status": "Processing"
}
 3. Fetch Metrics
Request:

sh
Copy
Edit
curl -X GET "http://localhost:8080/metrics"
Response:

json
Copy
Edit
{
  "total_orders": 2,
  "average_processing_time": 165.00,
  "order_status_counts": {
    "pending": 0,
    "processing": 0,
    "completed": 2
  }
}


Design Decisions & Trade-offs
3-layered architecture
Improves modularity and maintainability.
SQLite for storage
In-memory queue
JSON storage for item_ids
Explicit updated_at updates

Assumptions: 
Orders are unique and identified by order_id.
The item_ids field is stored as a JSON string (e.g., "[101, 102, 103]").
Order processing is simulated with a delay of random 2-5 seconds before marking an order as "Completed."
No user authentication is required for now (can be added later).
The queue is in-memory (No persistence, so orders in the queue will be lost if the server restarts).