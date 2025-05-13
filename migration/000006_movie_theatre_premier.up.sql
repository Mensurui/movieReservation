CREATE TABLE IF NOT EXISTS moviepremier(
	id SERIAL PRIMARY KEY,
	showtime TIMESTAMP WITH TIME ZONE,
	price NUMERIC NOT NULL,
	movie_id INTEGER NOT NULL,
	theatre_id INTEGER NOT NULL,


	CONSTRAINT fk_movie_id
		FOREIGN KEY(movie_id)
		REFERENCES movies(id),

	CONSTRAINT fk_theatre_id
		FOREIGN KEY(theatre_id)
		REFERENCES theatre(id)
)
