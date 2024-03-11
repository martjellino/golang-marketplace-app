CREATE TABLE purchases (
  purchase_id SERIAL PRIMARY KEY,
  buyer_id INT,
  account_id INT,
  product_id INT,
  qty INT,
  total_price INT,
  image_url VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (buyer_id) REFERENCES users(user_id),
  FOREIGN KEY (account_id) REFERENCES bank_accounts(account_id),
  FOREIGN KEY (product_id) REFERENCES products(product_id)
);