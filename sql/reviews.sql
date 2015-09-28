
DROP TABLE IF EXISTS review_diffs;
DROP TABLE IF EXISTS reviews;

CREATE TABLE reviews(
	id SERIAL primary key,
	title TEXT,
	author TEXT
);

CREATE TABLE review_diffs(
	diff_id SERIAL primary key,
	review_id INT references reviews(id),
	gitdiff TEXT
);

CREATE OR REPLACE FUNCTION add_review(in_title TEXT, in_author TEXT, in_gitdiff TEXT) RETURNS INTEGER
AS $$

DECLARE
	r_id INTEGER;

BEGIN
	INSERT INTO reviews(title, author) VALUES (in_title, in_author) RETURNING id into r_id;
	INSERT INTO review_diffs (review_id, gitdiff) VALUES (r_id, in_gitdiff);
	RETURN r_id;
END;

$$
LANGUAGE plpgsql;
