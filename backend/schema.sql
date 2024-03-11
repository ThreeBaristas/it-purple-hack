CREATE TABLE IF NOT EXISTS prices(
  category_id bigint NOT NULL,
  location_id bigint NOT NULL,
  price int NOT NULL, -- Цена на данную локацию и категорию
  matrix_id int -- ID матрицы. NULL, если это baseline
);

INSERT INTO prices(category_id, location_id, price, matrix_id) VALUES 
(1, 1, 500, 0), -- root, root, baseline
(20, 20, 600, 0), -- Spb, Vehicles, baseline
(21, 21, 300, 0); -- Petrogradskaya, Cars, baseline
