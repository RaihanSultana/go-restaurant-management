CREATE TABLE IF NOT EXISTS order_items(
  id SERIAL PRIMARY KEY,
  order_id INT REFERENCES orders(id),
  item_id INT REFERENCES items(id),
  quantity INT NOT NULL, 
  total NUMERIC(10, 2) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW()
);