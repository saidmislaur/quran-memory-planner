CREATE TABLE surahs (
    id BIGSERIAL PRIMARY KEY,
    number INT NOT NULL UNIQUE,
    name_arabic TEXT NOT NULL,
    name_english TEXT NOT NULL,
    name_russian TEXT NOT NULL,
    ayah_count INT NOT NULL
);

CREATE TABLE ayahs (
    id BIGSERIAL PRIMARY KEY,
    surah_id BIGINT NOT NULL REFERENCES surahs(id) ON DELETE CASCADE,
    number INT NOT NULL,
    arabic TEXT NOT NULL,
    translate TEXT NOT NULL,
    UNIQUE (surah_id, number)
);

CREATE TABLE words (
    id BIGSERIAL PRIMARY KEY,
    surah_id BIGINT NOT NULL REFERENCES surahs(id) ON DELETE CASCADE,
    ayah_id BIGINT NOT NULL REFERENCES ayahs(id) ON DELETE CASCADE,
    arabic TEXT NOT NULL,
    root TEXT,
    position INT NOT NULL,
    UNIQUE (ayah_id, position)
);

CREATE TABLE word_translations (
    id BIGSERIAL PRIMARY KEY,
    word_id BIGINT NOT NULL REFERENCES words(id) ON DELETE CASCADE,
    lang_code VARCHAR(5) NOT NULL,
    translation TEXT NOT NULL
);

CREATE TABLE word_progress (
    user_id BIGINT NOT NULL,
    word_id BIGINT NOT NULL,
    review_count INT NOT NULL DEFAULT 0,
    easy_count INT NOT NULL DEFAULT 0,
    last_reviewed_at TIMESTAMPTZ,
    next_review_at TIMESTAMPTZ NOT NULL,
    is_mastered BOOLEAN NOT NULL DEFAULT false,

    PRIMARY KEY (user_id, word_id)
)