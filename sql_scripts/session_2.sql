CREATE TABLE RealEstateObject (
    id SERIAL PRIMARY KEY,
    city VARCHAR(255),
    street VARCHAR(255),
    house_number VARCHAR(50),
    apartment_number VARCHAR(50),
    latitude NUMERIC CHECK (latitude >= -90 AND latitude <= 90),
    longitude NUMERIC CHECK (longitude >= -180 AND longitude <= 180)
);

CREATE TABLE Apartment (
    id SERIAL PRIMARY KEY,
    real_estate_object_id INT REFERENCES RealEstateObject(id),
    floor INT,
    number_of_rooms INT,
    area NUMERIC
);

CREATE TABLE House (
    id SERIAL PRIMARY KEY,
    real_estate_object_id INT REFERENCES RealEstateObject(id),
    number_of_floors INT,
    number_of_rooms INT,
    area NUMERIC
);

CREATE TABLE LandPlot (
    id SERIAL PRIMARY KEY,
    real_estate_object_id INT REFERENCES RealEstateObject(id),
    area NUMERIC
);

CREATE TABLE Offer (
    id SERIAL PRIMARY KEY,
    real_estate_object_id INT REFERENCES RealEstateObject(id) ON DELETE RESTRICT
);
