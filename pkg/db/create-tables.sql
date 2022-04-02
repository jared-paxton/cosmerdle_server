-- user table
CREATE TABLE IF NOT EXISTS app_user (
	user_key serial PRIMARY KEY,
  user_id VARCHAR (100) UNIQUE,
	email VARCHAR ( 100 ) UNIQUE ,
	password VARCHAR ( 100 ),
	created_on TIMESTAMP NOT NULL,
  last_activity TIMESTAMP NOT NULL,
);

-- game state table
CREATE TABLE IF NOT EXISTS game_state (
  gs_key serial PRIMARY KEY,
  user_key INT NOT NULL,
  game_status INT NOT NULL,
  current_guess INT NOT NULL,
  current_day DATE NOT NULL,
  guesses TEXT[] NOT NULL,
  FOREIGN KEY (user_key)
      REFERENCES app_user (user_key)
);

-- TODO: add other tables like user stats, word tables, etc. 