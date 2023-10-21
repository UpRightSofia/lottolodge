-- +goose Up
INSERT INTO users (id, updated_at, balance_e5) VALUES
    ('e861a88c-3023-49e4-8ce1-76896b87c238', '2023-10-21 15:30:00', 500000),
    ('11c54438-4c26-4d0d-a6e7-9b23b29299d0', '2023-10-21 14:45:00', 750000),
    ('0c7ff464-86c3-4952-9536-7f54b7267ea1', '2023-10-21 13:15:00', 250000),
    ('fb408c5f-4b0a-4da4-9ef3-d6e55a834a83', '2023-10-21 12:30:00', 125000),
    ('28c7b369-4e94-4ef2-84ef-930a597c660b', '2023-10-21 11:45:00', 600000);

INSERT INTO pools (id, details, created_at, is_active) VALUES
    ('d378f345-3727-4b2a-b816-2a1beba4c6c2', '{"drawn_numbers": [5, 10, 15, 20, 25, 30], "small_multiplier": 2, "big_multiplier": 3}', '2023-10-21 15:30:00', false),
    ('7a9ab991-117e-45d7-9c69-937b131415bb', '{"drawn_numbers": [7, 14, 21, 28, 35, 42], "small_multiplier": 3, "big_multiplier": 4}', '2023-10-21 14:45:00', false),
    ('c156b6ea-d2bf-4c61-bc67-e32bf7f1ca4c', '{"drawn_numbers": [3, 6, 9, 12, 15, 18], "small_multiplier": 2, "big_multiplier": 3}', '2023-10-21 13:15:00', false),
    ('f03de4dd-72df-4b9f-8e6a-54f01a27ca9e', '{"drawn_numbers": [8, 16, 24, 32, 40, 48], "small_multiplier": 3, "big_multiplier": 4}', '2023-10-21 12:30:00', false),
    ('a35e6fc7-2e26-4ea5-9470-d0b9d019cdde', '{"drawn_numbers": [2, 4, 6, 8, 10, 12], "small_multiplier": 2, "big_multiplier": 3}', '2023-10-21 11:45:00', true);

-- Tickets for User 1 in Pool 1 (with small_multiplier and big_multiplier)
INSERT INTO tickets (id, user_id, pool_id, details, is_hand_picked, is_used)
VALUES
    ('2f53e3a4-2b24-4a9f-8219-741a2f5d69b9', 'e861a88c-3023-49e4-8ce1-76896b87c238', 'd378f345-3727-4b2a-b816-2a1beba4c6c2', '{"numbers": [5, 10, 15, 20, 25, 30], "small_multiplier": 2, "big_multiplier": 4}', false, false),
    ('db22c975-55a7-42cc-91bb-c6a76ca5f1db', 'e861a88c-3023-49e4-8ce1-76896b87c238', 'd378f345-3727-4b2a-b816-2a1beba4c6c2', '{"numbers": [8, 16, 24, 32, 40, 48], "small_multiplier": 3, "big_multiplier": 5}', false, false),
    ('3d6b99ec-2764-4ea2-ba14-e6f8f98d3474', 'e861a88c-3023-49e4-8ce1-76896b87c238', 'd378f345-3727-4b2a-b816-2a1beba4c6c2', '{"numbers": [2, 4, 6, 8, 10, 12], "small_multiplier": 2, "big_multiplier": 3}', true, false);

-- Winnings User 1 Pool 1
INSERT INTO winnings (id, user_id, ticket_id, pool_id, prize_e5)
VALUES
    ('2beaef86-4550-43c1-8a05-e6c7366d73a6', 'e861a88c-3023-49e4-8ce1-76896b87c238', '2f53e3a4-2b24-4a9f-8219-741a2f5d69b9', 'd378f345-3727-4b2a-b816-2a1beba4c6c2', 2500),
    ('3a4a41f3-5ed0-4e07-8b9b-5438a92fda9e', 'e861a88c-3023-49e4-8ce1-76896b87c238', 'db22c975-55a7-42cc-91bb-c6a76ca5f1db', 'd378f345-3727-4b2a-b816-2a1beba4c6c2', 1800);

-- Tickets for User 1 in Pool 2 (with small_multiplier and big_multiplier)
INSERT INTO tickets (id, user_id, pool_id, details, is_hand_picked, is_used)
VALUES
    ('90bf74bb-e4e6-4e32-a683-eb7f9c5d7917', 'e861a88c-3023-49e4-8ce1-76896b87c238', '7a9ab991-117e-45d7-9c69-937b131415bb', '{"numbers": [7, 14, 21, 28, 35, 42], "small_multiplier": 3, "big_multiplier": 5}', false, false),
    ('a3f1ac01-9d31-4f10-8e94-742f7ee28dfc', 'e861a88c-3023-49e4-8ce1-76896b87c238', '7a9ab991-117e-45d7-9c69-937b131415bb', '{"numbers": [1, 2, 3, 4, 5, 6], "small_multiplier": 4, "big_multiplier": 6}', true, false);

-- Winning Sample for User 1 in Pool 2
INSERT INTO winnings (id, user_id, ticket_id, pool_id, prize_e5)
VALUES
    ('457b9610-3e3e-42ed-9f52-57e07e4f325d', 'e861a88c-3023-49e4-8ce1-76896b87c238', '90bf74bb-e4e6-4e32-a683-eb7f9c5d7917', '7a9ab991-117e-45d7-9c69-937b131415bb', 3200),
    ('f83ffab9-e4ac-4aa9-b456-e6fbcf9d4c3a', 'e861a88c-3023-49e4-8ce1-76896b87c238', 'a3f1ac01-9d31-4f10-8e94-742f7ee28dfc', '7a9ab991-117e-45d7-9c69-937b131415bb', 2500);

-- Tickets for User 1 in Pool 3 (with small_multiplier and big_multiplier)
INSERT INTO tickets (id, user_id, pool_id, details, is_hand_picked, is_used)
VALUES
    ('c82a3e41-3b86-47a4-b81d-963cb17e06d1', 'e861a88c-3023-49e4-8ce1-76896b87c238', 'c156b6ea-d2bf-4c61-bc67-e32bf7f1ca4c', '{"numbers": [3, 6, 9, 12, 15, 18], "small_multiplier": 2, "big_multiplier": 4}', false, false);


-- No winnings for User 1 in Pool 3

-- Tickets for User 1 in Pool 4 (with small_multiplier and big_multiplier)
INSERT INTO tickets (id, user_id, pool_id, details, is_hand_picked, is_used)
VALUES
    ('3be6c4b3-6aa1-49b3-9a7e-432739c03e42', 'e861a88c-3023-49e4-8ce1-76896b87c238', 'f03de4dd-72df-4b9f-8e6a-54f01a27ca9e', '{"numbers": [8, 16, 24, 32, 40, 48], "small_multiplier": 3, "big_multiplier": 5}', false, false),
    ('a6e8cfd0-92ce-4f90-a1bb-31ea6ff35f56', 'e861a88c-3023-49e4-8ce1-76896b87c238', 'f03de4dd-72df-4b9f-8e6a-54f01a27ca9e', '{"numbers": [4, 8, 12, 16, 20, 24], "small_multiplier": 2, "big_multiplier": 4}', true, false),
    ('b072e0a5-51ce-4e57-8643-33de86a98a91', 'e861a88c-3023-49e4-8ce1-76896b87c238', 'f03de4dd-72df-4b9f-8e6a-54f01a27ca9e', '{"numbers": [6, 12, 18, 24, 30, 36], "small_multiplier": 3, "big_multiplier": 5}', false, false),
    ('c49736f4-30b1-454b-9a3b-41a69504d27a', 'e861a88c-3023-49e4-8ce1-76896b87c238', 'f03de4dd-72df-4b9f-8e6a-54f01a27ca9e', '{"numbers": [9, 18, 27, 36, 45, 54], "small_multiplier": 2, "big_multiplier": 4}', true, false);

-- Winning Samples for User 1 in Pool 4
INSERT INTO winnings (id, user_id, ticket_id, pool_id, prize_e5)
VALUES
    ('fa2008a2-5ed0-47c4-b6a7-57e07e4f325d', 'e861a88c-3023-49e4-8ce1-76896b87c238', '3be6c4b3-6aa1-49b3-9a7e-432739c03e42', 'f03de4dd-72df-4b9f-8e6a-54f01a27ca9e', 2200),
    ('a1c1e3b9-e4ac-4ba9-b6d6-e6fbcf9d4c3a', 'e861a88c-3023-49e4-8ce1-76896b87c238', 'a6e8cfd0-92ce-4f90-a1bb-31ea6ff35f56', 'f03de4dd-72df-4b9f-8e6a-54f01a27ca9e', 1800),
    ('d4e9e6c7-7b63-4f1a-9a4a-41a69504d27a', 'e861a88c-3023-49e4-8ce1-76896b87c238', 'b072e0a5-51ce-4e57-8643-33de86a98a91', 'f03de4dd-72df-4b9f-8e6a-54f01a27ca9e', 2400);

-- Sample data for settings table
INSERT INTO settings (ticket_prize_e5, payout_percent) VALUES
    (1000, 70),
    (1500, 75),
    (2000, 80),
    (2500, 85),
    (3000, 90);

-- +goose Down