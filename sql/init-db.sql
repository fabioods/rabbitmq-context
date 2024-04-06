CREATE DATABASE IF NOT EXISTS `balances`;
CREATE DATABASE IF NOT EXISTS `wallet`;

Create table wallet.transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date);
Create table wallet.accounts (id varchar(255), client_id varchar(255), balance int, created_at date);
Create table wallet.clients (id varchar(255), name varchar(255), email varchar(255), created_at date);

INSERT INTO wallet.clients (id, name, email, created_at) VALUES ('1', 'John Doe', 'john@gmail.com', '2024-03-11');
INSERT INTO wallet.clients (id, name, email, created_at) VALUES ('2', 'Jane Doe', 'jane@gmail.com', '2024-03-11');
INSERT INTO wallet.accounts (id, client_id, balance, created_at) VALUES ('1', '1', 1000, '2024-03-11');
INSERT INTO wallet.accounts (id, client_id, balance, created_at) VALUES ('2', '2', 10000, '2024-03-11');

Create table balances.transactions (id varchar(255), account_id_from varchar(255), account_id_to varchar(255), amount int, created_at date);
Create table balances.accounts (id varchar(255), client_id varchar(255), balance int, created_at date);
Create table balances.clients (id varchar(255), name varchar(255), email varchar(255), created_at date);

INSERT INTO balances.clients (id, name, email, created_at) VALUES ('1', 'John Doe', 'john@gmail.com', '2024-03-11');
INSERT INTO balances.clients (id, name, email, created_at) VALUES ('2', 'Jane Doe', 'jane@gmail.com', '2024-03-11');
INSERT INTO balances.accounts (id, client_id, balance, created_at) VALUES ('1', '1', 1000, '2024-03-11');
INSERT INTO balances.accounts (id, client_id, balance, created_at) VALUES ('2', '2', 10000, '2024-03-11');
