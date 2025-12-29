CREATE TABLE IF NOT EXISTS orders(
  id SERIAL PRIMARY KEY,
  uuid UUID UNIQUE DEFAULT gen_random_uuid(),
  customer_id INT REFERENCES customers(id),
  subtotal NUMERIC(10,2) NOT NULL,
  total NUMERIC(10, 2) NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW()
);