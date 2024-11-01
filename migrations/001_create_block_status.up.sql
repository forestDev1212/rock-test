-- Create the block_status table
CREATE TABLE  block_status (
    id SERIAL PRIMARY KEY,
    last_block_number BIGINT NOT NULL DEFAULT 0
);

-- Insert an initial record to track block numbers
INSERT INTO block_status (last_block_number) VALUES (100000);
