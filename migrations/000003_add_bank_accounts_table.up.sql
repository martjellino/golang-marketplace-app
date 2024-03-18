CREATE TABLE bank_accounts (
  account_id SERIAL PRIMARY KEY,
  user_id INT,
  bank_name VARCHAR(15),
  account_name VARCHAR(15),
  account_number VARCHAR(15),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(user_id)
);