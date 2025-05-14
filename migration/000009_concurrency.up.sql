ALTER TABLE reservations
ADD CONSTRAINT unique_user_movie UNIQUE(user_id, movie_premier_id);
