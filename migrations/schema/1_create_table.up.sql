-- Check if the logs table already exists
DO
$$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'logs') THEN
        -- Create logs table
CREATE TABLE logs (
	id serial4 NOT NULL,
	"type" varchar(255) NULL,
	"timestamp" timestamptz NULL,
	details jsonb NULL,
	CONSTRAINT logs_pkey PRIMARY KEY (id),
	CONSTRAINT logs_unique UNIQUE ("timestamp")
);
CREATE INDEX logs_details_idx ON public.logs USING btree (details);
END IF;
END
$$;

-- Create logs_bulk table
CREATE TABLE logs_bulk (
	id serial4 NOT NULL,
	"type" varchar(255) NULL,
	"timestamp" timestamptz NULL,
	details jsonb NULL,
	CONSTRAINT logs_bulk_pkey PRIMARY KEY (id),
	CONSTRAINT logs_bulk_unique UNIQUE ("timestamp")
);
CREATE INDEX logs_bulk_details_idx ON public.logs_bulk USING btree (details);

-- public.jobs definition

-- Drop table

-- DROP TABLE jobs;

CREATE TABLE jobs (
	id serial4 NOT NULL,
	module_id varchar(255) NULL,
	job_id varchar(255) NULL,
	status varchar(255) NULL,
	"time_start" timestamptz NULL,
	"time_update" timestamptz NULL,
	details jsonb NULL,
	CONSTRAINT events_pkey PRIMARY KEY (id),
	CONSTRAINT events_unique UNIQUE (job_id)
);
CREATE INDEX events_details_idx ON public.jobs USING btree (details);

-- Define the trigger function to log data from logs table to events table
-- CREATE OR REPLACE FUNCTION log_table_trigger_function()
-- RETURNS TRIGGER AS $$
-- DECLARE
--     status TEXT;
-- BEGIN
--     status := '';
    
--     IF TG_OP = 'INSERT' OR TG_OP = 'UPDATE' THEN
--         CASE NEW.type
--             WHEN 'solver' THEN
--                 status := 'start';
--             WHEN 'pending' THEN
--                 status := 'Pending';
--             WHEN 'started' THEN
--                 status := 'Started';
--             WHEN 'working' THEN
--                 status := 'Working';
--             WHEN 'complete' THEN
--                 status := 'Complete';
--             -- ELSE
--             --     status := 'Invalid';
--             --     RAISE EXCEPTION 'Invalid log type: %', NEW.type;
--         END CASE;

--         -- Perform the operation on the events table
--         INSERT INTO jobs (module_id, job_id, status, duration, timestamp, details)
--         VALUES (NEW.type, '1', status, '1', NEW.timestamp, NEW.details);
--     END IF;

--     RETURN NEW;
-- END;
-- $$ LANGUAGE plpgsql;

-- Create trigger on logs table to call the trigger function after insert or update
-- CREATE TRIGGER log_table_trigger
-- AFTER INSERT OR UPDATE ON logs
-- FOR EACH ROW EXECUTE FUNCTION log_table_trigger_function();

-- Create the trigger function for inserting into logs table from logs_bulk
CREATE OR REPLACE FUNCTION insert_into_logs()
RETURNS TRIGGER AS $$
BEGIN
    -- Check if the record already exists in logs table
    IF NOT EXISTS (SELECT 1 FROM logs WHERE timestamp = NEW.timestamp) THEN
        INSERT INTO logs (type, timestamp, details)
        VALUES (NEW.type, NEW.timestamp, NEW.details);
        NOTIFY updates;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create a trigger on logs_bulk table to call the function after insert
CREATE TRIGGER insert_into_logs_trigger
AFTER INSERT ON public.logs_bulk
FOR EACH ROW
EXECUTE FUNCTION insert_into_logs();
