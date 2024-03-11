CREATE TABLE bank_accounts (
  account_id SERIAL PRIMARY KEY,
  user_id INT,
  account_name VARCHAR(255),
  account_number VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(user_id)
);