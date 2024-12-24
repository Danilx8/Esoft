-- Создание таблицы districts
CREATE TABLE districts (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    area POLYGON NOT NULL
);

-- Создание временной таблицы для импорта данных
CREATE TEMP TABLE districts_csv (
    name VARCHAR(100),
    area VARCHAR(10000)
);

-- Копирование данных из CSV
\COPY districts_csv FROM '/csv/Session 2/districts.csv' WITH (FORMAT CSV, HEADER);

-- Вставка данных с преобразованием строки в polygon
INSERT INTO districts (name, area)
SELECT 
    name,
    area::polygon
FROM districts_csv;

