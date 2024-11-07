CREATE TABLE customers (
  id VARCHAR(255) NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE promotions (
  id VARCHAR(255) NOT NULL,
  name VARCHAR(255) NOT NULL,
  discount_amount INT NOT NULL,
  slot_remaining_amount INT NOT NULL,
  submitted_at DATETIME NOT NULL,
  modified_at DATETIME NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE promotion_publications (
  promotion_id VARCHAR(255) NOT NULL,
  published_at DATETIME NOT NULL,
  PRIMARY KEY (promotion_id),
  CONSTRAINT fk_promotion_publications_promotions FOREIGN KEY (promotion_id) REFERENCES promotions (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE promotion_applications (
  customer_id VARCHAR(255) NOT NULL,
  promotion_id VARCHAR(255) NOT NULL,
  applied_at DATETIME NOT NULL,
  PRIMARY KEY (customer_id, promotion_id),
  CONSTRAINT fk_promotion_applications_customers FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT fk_promotion_applications_promotions FOREIGN KEY (promotion_id) REFERENCES promotions (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE coupons (
  id VARCHAR(255) NOT NULL,
  discount_amount INT NOT NULL,
  expires_at DATETIME NOT NULL,
  customer_id VARCHAR(255) NOT NULL,
  granted_at DATETIME NOT NULL,
  PRIMARY KEY (id),
  CONSTRAINT fk_coupons_customers FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE coupon_invalidations (
  coupon_id VARCHAR(255) NOT NULL,
  invalidated_at DATETIME NOT NULL,
  PRIMARY KEY (coupon_id),
  CONSTRAINT fk_coupon_invalidations_coupons FOREIGN KEY (coupon_id) REFERENCES coupons (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE coupon_redemptions (
  coupon_id VARCHAR(255) NOT NULL,
  redeemed_at DATETIME NOT NULL,
  PRIMARY KEY (coupon_id),
  CONSTRAINT fk_coupon_redemptions_coupons FOREIGN KEY (coupon_id) REFERENCES coupons (id) ON DELETE RESTRICT ON UPDATE RESTRICT
);
