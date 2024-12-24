-- Создание перечисления для типов недвижимости
CREATE TYPE property_type AS ENUM ('apartment', 'house', 'land');

-- Таблица адресов
CREATE TABLE addresses (
    id SERIAL PRIMARY KEY,
    city VARCHAR(100),
    street VARCHAR(100),
    house_number VARCHAR(20),
    apartment_number VARCHAR(20)
);

-- Таблица координат
CREATE TABLE coordinates (
    id SERIAL PRIMARY KEY,
    latitude DECIMAL(9,6) CHECK (latitude >= -90 AND latitude <= 90),
    longitude DECIMAL(9,6) CHECK (longitude >= -180 AND longitude <= 180)
);

-- Основная таблица объектов недвижимости
CREATE TABLE properties (
    id SERIAL PRIMARY KEY,
    property_type property_type NOT NULL,
    address_id INTEGER REFERENCES addresses(id),
    coordinates_id INTEGER REFERENCES coordinates(id),
    floor INTEGER, -- для квартир
    rooms_count INTEGER, -- для квартир и домов
    area DECIMAL(10,2), -- для всех типов
    house_floors INTEGER, -- для домов
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Триггер для обновления updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$
 language 'plpgsql';

CREATE TRIGGER update_properties_updated_at
    BEFORE UPDATE ON properties
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Таблица предложений (для предотвращения удаления связанных объектов)
CREATE TABLE offers (
    id SERIAL PRIMARY KEY,
    property_id INTEGER REFERENCES properties(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Индексы для улучшения производительности фильтрации
CREATE INDEX idx_properties_type ON properties(property_type);
CREATE INDEX idx_addresses_city ON addresses(city);
CREATE INDEX idx_addresses_street ON addresses(street);
