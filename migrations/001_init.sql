CREATE TABLE IF NOT EXISTS submissions (
    id          VARCHAR(36) PRIMARY KEY,
    language    VARCHAR(20) NOT NULL,
    code        TEXT NOT NULL,
    status      VARCHAR(20) NOT NULL DEFAULT 'pending',
    created_at  TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS test_results (
    id                VARCHAR(36) PRIMARY KEY,
    submission_id     VARCHAR(36) REFERENCES submissions(id),
    testcase_number   INT NOT NULL,
    status            VARCHAR(20) NOT NULL,
    expected_output   TEXT,
    actual_output     TEXT,
    execution_time_ms INT,
    memory_used_kb    INT
);