CREATE TABLE product (
    id VARCHAR(50) PRIMARY KEY,
    marketplace VARCHAR(20) NOT NULL,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	image_url TEXT NOT NULL,
    price float NOT NULL,
    rating float NOT NULL,
    shop_name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NULL
);
