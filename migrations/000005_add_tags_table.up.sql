CREATE TABLE tags (
  tag_id SERIAL PRIMARY KEY,
  tag_name VARCHAR(50) UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE product_tags (
  product_id INT,
  tag_id INT,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (product_id, tag_id),
  FOREIGN KEY (product_id) REFERENCES products(product_id),
  FOREIGN KEY (tag_id) REFERENCES tags(tag_id)
);