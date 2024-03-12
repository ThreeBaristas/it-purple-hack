CREATE TABLE IF NOT EXISTS prices(
  category_id bigint NOT NULL,
  location_id bigint NOT NULL,
  price int NOT NULL, -- Цена на данную локацию и категорию
  matrix_id int, -- ID матрицы. NULL, если это baseline
  CONSTRAINT PK_price PRIMARY KEY (category_id, location_id, matrix_id)
);

