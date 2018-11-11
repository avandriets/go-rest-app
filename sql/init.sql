DROP TABLE if exists nprl.public.Structures;

CREATE TABLE nprl.public.structures
(
  "noId"	NUMERIC PRIMARY KEY,
  "compounds"	TEXT,
  "smiles"	TEXT,
  "formular"	TEXT,
  "structure"	TEXT,
  "structure_type" TEXT,
  "bioactivity_type" TEXT,
  "activity"	TEXT,
  "journalName"	TEXT,
  "year"	NUMERIC,
  "volume"	NUMERIC,
  "page"	NUMERIC,
  "articleName"	TEXT,
  "articleNo"	NUMERIC,
  created_at timestamp  NOT NULL  DEFAULT current_timestamp,
  updated_at timestamp  NOT NULL  DEFAULT current_timestamp
);

COPY nprl.public.structures(
  "noId",
  "compounds",
  "smiles",
  "formular",
  "structure",
  "structure_type",
  "bioactivity_type",
  "activity",
  "journalName",
  "year",
  "volume",
  "page",
  "articleName",
  "articleNo"
) FROM '/docker-entrypoint-initdb.d/csv-data.csv' DELIMITER ',' CSV;
