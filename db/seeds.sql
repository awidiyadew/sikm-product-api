INSERT INTO users (name, role, email, password, created_at, updated_at)
VALUES 
    ('John Doe', 'ADMIN','john.doe@mail.com', '123456', now(), now()),
    ('Amanda Doe', 'USER', 'amanda.doe@mail.com', '123456', now(), now());

INSERT INTO categories (name, created_at, updated_at)
VALUES 
    ('Snack', now(), now()),
    ('Minuman', now(), now()),
    ('Makanan', now(), now()),
    ('Fashion', now(), now());

INSERT INTO products (name, price, category_id, posted_by, created_at, updated_at)
VALUES 
    ('Beng-beng', 2500, 1, 1, now(), now()),
    ('Top', 1500, 1, 1, now(), now()),
    ('Larutan Kaki Tiga', 5000, 2, 1, now(), now());
