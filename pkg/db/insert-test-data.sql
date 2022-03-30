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