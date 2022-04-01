INSERT INTO app_user (
  email, 
  password, 
  created_on 
) VALUES (
  'jaredpaxton.work@gmail.com',
  'dontguessmypassword',
  current_timestamp
);

INSERT INTO game_state (
  user_key,
  game_status,
  current_guess,
  current_day,
  guesses
) VALUES (
  1,
  0,
  1,
  '2022-03-29',
  '{"storm"}'
);

INSERT INTO app_user (
  user_id, 
  created_on,
  last_activity 
) VALUES (
  '123456789',
  current_timestamp,
  '2022-03-31'
);