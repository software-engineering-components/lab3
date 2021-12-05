DROP DATABASE IF EXISTS "restaurant";
CREATE DATABASE "restaurant";

CREATE TABLE IF NOT EXISTS "menu"(
  "id" SERIAL PRIMARY KEY,
  "name" varchar(80) NOT NULL,
  "price" INT NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "orders"(
  "id" SERIAL PRIMARY KEY,
  "table_number" INT NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "details" (
  "order_id" INT NOT NULL,
  "menu_id" INT NOT NULL,
  "quantity" INT NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY ("order_id", "menu_id"),
  CONSTRAINT fk_details_menu_id_menu_id FOREIGN KEY ("menu_id") REFERENCES menu ("id"),
  CONSTRAINT fk_details_order_id_orders_id FOREIGN KEY ("order_id") REFERENCES orders ("id")
);

INSERT INTO "menu"("name", "price")
VALUES
  ('Le Pigeon Burger. Le Pigeon – Portland', 10),
  ('Whiskey King Burger. Village Whiskey – Philadelphia, PA', 12),
  ('Chargrilled Burger with Roquefort Cheese', 30),
  ('The Company Burger', 20),
  ('Dyers Deep-Fried Burger', 23),
  ('The Lola Burger', 11),
  ('Cheeseburger', 21),
  ('Raw Steak Tartare Burger', 15);

INSERT INTO "orders"("table_number")
VALUES
  (1),
  (2),
  (3);

INSERT INTO "details"("order_id", "menu_id", "quantity")
VALUES
  (1, 1, 10),
  (2, 1, 12),
  (3, 2, 1);
