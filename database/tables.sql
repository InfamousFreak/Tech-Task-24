CREATE TABLE menu_items ( 
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    name VARCHAR(100) NOT NULL, /*variable length character strings upto 100 chars, and name column cant be null*/
    description TEXT,
    price DECIMAL(10, 2) NOT NULL, /*10 indicates the total number of digits stored and 2 shows no. of digits to the right of decimal point*/
    tags VARCHAR(255)
);

CREATE TABLE user_profiles (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE cart_items (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    user_id INTEGER NOT NULL,
    menu_item_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user_profiles(id), /*foreign key constraint on a colummn means that the column should be referred to and correspond to another column mentioned in the references part*/
    FOREIGN KEY (menu_item_id) REFERENCES menu_items(id)
);

CREATE TABLE suggestions (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    user_id INTEGER NOT NULL,
    menu_item_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user_profiles(id),
    FOREIGN KEY (menu_item_id) REFERENCES menu_items(id)
);
