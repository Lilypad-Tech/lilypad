
CREATE TABLE logs (
    id SERIAL PRIMARY KEY,
    type VARCHAR(255),
    timestamp TIMESTAMP WITH TIME ZONE,
    details TEXT
);


CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    module_id VARCHAR(255),
    job_id VARCHAR(255),
    status VARCHAR(255),
    duration VARCHAR(255),
    timestamp TIMESTAMP WITH TIME ZONE,
    details TEXT
);


CREATE OR REPLACE FUNCTION log_table_trigger_function()
RETURNS TRIGGER AS $$
DECLARE
    status TEXT;
BEGIN
    status := '';
    
    IF TG_OP = 'INSERT' OR TG_OP = 'UPDATE' THEN
        CASE NEW.type
            WHEN 'solver' THEN
                status := 'start';
            WHEN 'pending' THEN
                status := 'Pending';
            WHEN 'started' THEN
                status := 'Started';
            WHEN 'working' THEN
                status := 'Working';
            WHEN 'complete' THEN
                status := 'Complete';
            ELSE
                RAISE EXCEPTION 'Invalid log type: %', NEW.type;
        END CASE;

        -- Perform the operation on the events table
        INSERT INTO events (module_id, job_id, status, duration, timestamp, details)
        VALUES (NEW.type, '1' , status, '1', NEW.timestamp, NEW.details);
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;



CREATE TRIGGER log_table_trigger
AFTER INSERT OR UPDATE ON logs
FOR EACH ROW EXECUTE FUNCTION log_table_trigger_function();
