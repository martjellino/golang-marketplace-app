CREATE TABLE products (
  product_id SERIAL PRIMARY KEY,
  seller_id INT,
  name VARCHAR(255),
  price INT,
  image_url VARCHAR(255),
  stock INT,
  condition VARCHAR(255),
  tags VARCHAR(255),
  is_purchaseable BOOLEAN,
  purchase_count INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (seller_id) REFERENCES users(user_id)
);