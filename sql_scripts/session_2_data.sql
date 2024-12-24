-- Создание временной таблицы для квартир
CREATE TEMP TABLE apartments_csv (
    Id INTEGER,
    Address_City VARCHAR(100),
    Address_Street VARCHAR(100),
    Address_House VARCHAR(20),
    Address_Number VARCHAR(20),
    Coordinate_latitude DECIMAL(9,6),
    Coordinate_longitude DECIMAL(9,6),
    TotalArea DECIMAL(10,2),
    Rooms INTEGER,
    Floor INTEGER
);

-- Создание временной таблицы для домов
CREATE TEMP TABLE houses_csv (
    Id INTEGER,
    Address_City VARCHAR(100),
    Address_Street VARCHAR(100),
    Address_House VARCHAR(20),
    Address_Number VARCHAR(20),
    Coordinate_latitude DECIMAL(9,6),
    Coordinate_longitude DECIMAL(9,6),
    TotalFloors INTEGER,
    TotalArea DECIMAL(10,2)
);

-- Создание временной таблицы для земельных участков
CREATE TEMP TABLE lands_csv (
    Id INTEGER,
    Address_City VARCHAR(100),
    Address_Street VARCHAR(100),
    Address_House VARCHAR(20),
    Address_Number VARCHAR(20),
    Coordinate_latitude DECIMAL(9,6),
    Coordinate_longitude DECIMAL(9,6),
    TotalArea DECIMAL(10,2)
);
-- Копирование данных из CSV файлов
\COPY apartments_csv FROM '/csv/Session 2/apartments.csv' WITH (FORMAT CSV, HEADER);
\COPY houses_csv FROM '/csv/Session 2/houses.csv' WITH (FORMAT CSV, HEADER);
\COPY lands_csv FROM '/csv/Session 2/lands.csv' WITH (FORMAT CSV, HEADER);

-- Заполнение таблицы адресов для квартир
WITH apartment_addresses AS (
    INSERT INTO addresses (city, street, house_number, apartment_number)
    SELECT DISTINCT 
        Address_City,
        Address_Street,
        Address_House,
        Address_Number
    FROM apartments_csv
    RETURNING id, city, street, house_number, apartment_number
),
-- Заполнение таблицы координат для квартир
apartment_coordinates AS (
    INSERT INTO coordinates (latitude, longitude)
    SELECT DISTINCT 
        Coordinate_latitude,
        Coordinate_longitude
    FROM apartments_csv
    RETURNING id, latitude, longitude
)
-- Заполнение таблицы properties для квартир
INSERT INTO properties (
    property_type,
    address_id,
    coordinates_id,
    floor,
    rooms_count,
    area
)
SELECT 
    'apartment'::property_type,
    a.id,
    c.id,
    ap.Floor,
    ap.Rooms,
    ap.TotalArea
FROM apartments_csv ap
JOIN apartment_addresses a ON 
    a.city = ap.Address_City AND 
    a.street = ap.Address_Street AND 
    a.house_number = ap.Address_House AND 
    a.apartment_number = ap.Address_Number
JOIN apartment_coordinates c ON 
    c.latitude = ap.Coordinate_latitude AND 
    c.longitude = ap.Coordinate_longitude;

-- Заполнение таблицы адресов для домов
WITH house_addresses AS (
    INSERT INTO addresses (city, street, house_number)
    SELECT DISTINCT 
        Address_City,
        Address_Street,
        Address_House
    FROM houses_csv
    RETURNING id, city, street, house_number
),
-- Заполнение таблицы координат для домов
house_coordinates AS (
    INSERT INTO coordinates (latitude, longitude)
    SELECT DISTINCT 
        Coordinate_latitude,
        Coordinate_longitude
    FROM houses_csv
    RETURNING id, latitude, longitude
)
-- Заполнение таблицы properties для домов
INSERT INTO properties (
    property_type,
    address_id,
    coordinates_id,
    house_floors,
    area
)
SELECT 
    'house'::property_type,
    a.id,
    c.id,
    h.TotalFloors,
    h.TotalArea
FROM houses_csv h
JOIN house_addresses a ON 
    a.city = h.Address_City AND 
    a.street = h.Address_Street AND 
    a.house_number = h.Address_House
JOIN house_coordinates c ON 
    c.latitude = h.Coordinate_latitude AND 
    c.longitude = h.Coordinate_longitude;

-- Заполнение таблицы адресов для земельных участков
WITH land_addresses AS (
    INSERT INTO addresses (city, street, house_number)
    SELECT DISTINCT 
        Address_City,
        Address_Street,
        Address_House
    FROM lands_csv
    RETURNING id, city, street, house_number
),
-- Заполнение таблицы координат для земельных участков
land_coordinates AS (
    INSERT INTO coordinates (latitude, longitude)
    SELECT DISTINCT 
        Coordinate_latitude,
        Coordinate_longitude
    FROM lands_csv
    RETURNING id, latitude, longitude
)
-- Заполнение таблицы properties для земельных участков
INSERT INTO properties (
    property_type,
    address_id,
    coordinates_id,
    area
)
SELECT 
    'land'::property_type,
    a.id,
    c.id,
    l.TotalArea
FROM lands_csv l
JOIN land_addresses a ON 
    a.city = l.Address_City AND 
    a.street = l.Address_Street AND 
    a.house_number = l.Address_House
JOIN land_coordinates c ON 
    c.latitude = l.Coordinate_latitude AND 
    c.longitude = l.Coordinate_longitude;
