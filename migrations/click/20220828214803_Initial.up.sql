 CREATE TABLE IF NOT EXISTS user_creation ( 
    user_id UInt64,
    name String,
    age UInt8,
    email String,
    time DateTime
) ENGINE = MergeTree
PARTITION BY toYYYYMM(time)
ORDER BY (user_id);


CREATE TABLE IF NOT EXISTS user_creation_stream (
    user_id UInt64,
    name String,
    age UInt8,
    email String,
    time DateTime
  ) ENGINE = Kafka('broker:9092', 'user_creation', 'click_1', 'JSONEachRow');
 

CREATE MATERIALIZED VIEW IF NOT EXISTS user_creation_v TO user_creation AS
SELECT 
 	user_id,
 	name,
    age,
    email,
    time 
FROM user_creation_stream;


