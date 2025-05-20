-- -- Insert test users if they don't exist
-- DO $$
-- BEGIN
--     IF NOT EXISTS (SELECT 1 FROM users WHERE username = 'john_doe') THEN
--         INSERT INTO users (username, password_hash) VALUES
--         ('john_doe', '$2a$10$1234567890123456789012');
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM users WHERE username = 'jane_smith') THEN
--         INSERT INTO users (username, password_hash) VALUES
--         ('jane_smith', '$2a$10$abcdefghijklmnopqrstuv');
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM users WHERE username = 'bob_wilson') THEN
--         INSERT INTO users (username, password_hash) VALUES
--         ('bob_wilson', '$2a$10$zyxwvutsrqponmlkjihgfe');
--     END IF;
-- END $$;

-- -- Insert test chats if they don't exist
-- DO $$
-- BEGIN
--     IF NOT EXISTS (SELECT 1 FROM chats WHERE name = 'General Chat') THEN
--         INSERT INTO chats (name, is_private) VALUES
--         ('General Chat', false);
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM chats WHERE name = 'John and Jane Private') THEN
--         INSERT INTO chats (name, is_private) VALUES
--         ('John and Jane Private', true);
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM chats WHERE name = 'Development Team') THEN
--         INSERT INTO chats (name, is_private) VALUES
--         ('Development Team', false);
--     END IF;
-- END $$;

-- -- Add users to chats if they aren't already participants
-- DO $$
-- DECLARE
--     general_chat_id INTEGER;
--     private_chat_id INTEGER;
--     dev_team_chat_id INTEGER;
--     john_id INTEGER;
--     jane_id INTEGER;
--     bob_id INTEGER;
-- BEGIN
--     -- Get chat IDs
--     SELECT id INTO general_chat_id FROM chats WHERE name = 'General Chat';
--     SELECT id INTO private_chat_id FROM chats WHERE name = 'John and Jane Private';
--     SELECT id INTO dev_team_chat_id FROM chats WHERE name = 'Development Team';

--     -- Get user IDs
--     SELECT id INTO john_id FROM users WHERE username = 'john_doe';
--     SELECT id INTO jane_id FROM users WHERE username = 'jane_smith';
--     SELECT id INTO bob_id FROM users WHERE username = 'bob_wilson';

--     -- Add participants if they don't exist
--     IF NOT EXISTS (SELECT 1 FROM chat_participants WHERE chat_id = general_chat_id AND user_id = john_id) THEN
--         INSERT INTO chat_participants (chat_id, user_id) VALUES (general_chat_id, john_id);
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM chat_participants WHERE chat_id = general_chat_id AND user_id = jane_id) THEN
--         INSERT INTO chat_participants (chat_id, user_id) VALUES (general_chat_id, jane_id);
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM chat_participants WHERE chat_id = general_chat_id AND user_id = bob_id) THEN
--         INSERT INTO chat_participants (chat_id, user_id) VALUES (general_chat_id, bob_id);
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM chat_participants WHERE chat_id = private_chat_id AND user_id = john_id) THEN
--         INSERT INTO chat_participants (chat_id, user_id) VALUES (private_chat_id, john_id);
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM chat_participants WHERE chat_id = private_chat_id AND user_id = jane_id) THEN
--         INSERT INTO chat_participants (chat_id, user_id) VALUES (private_chat_id, jane_id);
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM chat_participants WHERE chat_id = dev_team_chat_id AND user_id = john_id) THEN
--         INSERT INTO chat_participants (chat_id, user_id) VALUES (dev_team_chat_id, john_id);
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM chat_participants WHERE chat_id = dev_team_chat_id AND user_id = bob_id) THEN
--         INSERT INTO chat_participants (chat_id, user_id) VALUES (dev_team_chat_id, bob_id);
--     END IF;
-- END $$;

-- -- Insert test messages if they don't exist
-- DO $$
-- DECLARE
--     general_chat_id INTEGER;
--     private_chat_id INTEGER;
--     dev_team_chat_id INTEGER;
--     john_id INTEGER;
--     jane_id INTEGER;
--     bob_id INTEGER;
-- BEGIN
--     -- Get chat IDs
--     SELECT id INTO general_chat_id FROM chats WHERE name = 'General Chat';
--     SELECT id INTO private_chat_id FROM chats WHERE name = 'John and Jane Private';
--     SELECT id INTO dev_team_chat_id FROM chats WHERE name = 'Development Team';

--     -- Get user IDs
--     SELECT id INTO john_id FROM users WHERE username = 'john_doe';
--     SELECT id INTO jane_id FROM users WHERE username = 'jane_smith';
--     SELECT id INTO bob_id FROM users WHERE username = 'bob_wilson';

--     -- Insert messages if they don't exist (checking content to avoid duplicates)
--     IF NOT EXISTS (SELECT 1 FROM messages WHERE chat_id = general_chat_id AND user_id = john_id AND content = 'Hello everyone! Welcome to the General Chat!') THEN
--         INSERT INTO messages (chat_id, user_id, content) VALUES
--         (general_chat_id, john_id, 'Hello everyone! Welcome to the General Chat!');
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM messages WHERE chat_id = general_chat_id AND user_id = jane_id AND content = 'Hi John! Thanks for creating this chat.') THEN
--         INSERT INTO messages (chat_id, user_id, content) VALUES
--         (general_chat_id, jane_id, 'Hi John! Thanks for creating this chat.');
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM messages WHERE chat_id = general_chat_id AND user_id = bob_id AND content = 'Great to be here!') THEN
--         INSERT INTO messages (chat_id, user_id, content) VALUES
--         (general_chat_id, bob_id, 'Great to be here!');
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM messages WHERE chat_id = private_chat_id AND user_id = john_id AND content = 'Hi Jane, this is a private chat between us.') THEN
--         INSERT INTO messages (chat_id, user_id, content) VALUES
--         (private_chat_id, john_id, 'Hi Jane, this is a private chat between us.');
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM messages WHERE chat_id = private_chat_id AND user_id = jane_id AND content = 'Perfect, now we can discuss our project privately.') THEN
--         INSERT INTO messages (chat_id, user_id, content) VALUES
--         (private_chat_id, jane_id, 'Perfect, now we can discuss our project privately.');
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM messages WHERE chat_id = dev_team_chat_id AND user_id = john_id AND content = 'Welcome to the Development Team chat!') THEN
--         INSERT INTO messages (chat_id, user_id, content) VALUES
--         (dev_team_chat_id, john_id, 'Welcome to the Development Team chat!');
--     END IF;

--     IF NOT EXISTS (SELECT 1 FROM messages WHERE chat_id = dev_team_chat_id AND user_id = bob_id AND content = 'Looking forward to working with you all!') THEN
--         INSERT INTO messages (chat_id, user_id, content) VALUES
--         (dev_team_chat_id, bob_id, 'Looking forward to working with you all!');
--     END IF;
-- END $$; 