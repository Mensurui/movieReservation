CREATE TABLE IF NOT EXISTS reservationcapacity (
	id SERIAL PRIMARY KEY,
	movie_premier_id INTEGER NOT NULL,
	current_capacity INTEGER NOT NULL DEFAULT 0,
	CONSTRAINT fk_movie_premier_id
		FOREIGN KEY (movie_premier_id)
		REFERENCES moviepremier(id)
);
