CREATE TABLE products (
  product_id SERIAL PRIMARY KEY,
  seller_id INT,
  name VARCHAR(60),
  price INT,
  image_url TEXT,
  stock INT,
  condition VARCHAR(6),
  is_purchaseable BOOLEAN,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (seller_id) REFERENCES users(user_id)
);