-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS citext;
CREATE EXTENSION IF NOT EXISTS moddatetime;

CREATE TABLE owners
(
  id             UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  full_name      VARCHAR(255) NOT NULL,
  license_number VARCHAR(50)  NOT NULL UNIQUE,
  phone          VARCHAR(20)  NOT NULL UNIQUE,
  email          VARCHAR(100) NOT NULL UNIQUE,
  created_at     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at     TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE vehicles
(
  id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  license_plate VARCHAR(10) NOT NULL UNIQUE,
  model         VARCHAR(50) NOT NULL,
  owner_id      UUID        REFERENCES owners (id) ON DELETE SET NULL,
  created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at    TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE fines
(
  id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  vehicle_id    UUID REFERENCES vehicles (id) ON DELETE CASCADE,
  issue_date    DATE           NOT NULL,
  due_date      DATE           NOT NULL,
  amount        NUMERIC(10, 2) NOT NULL,
  status        VARCHAR(50)    NOT NULL,
  description   TEXT,
  created_at    TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at    TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE payments
(
  id             UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  fine_id        UUID REFERENCES fines (id) ON DELETE CASCADE,
  paid_date      TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  amount         NUMERIC(10, 2) NOT NULL,
  payment_method VARCHAR(50)    NOT NULL,
  created_at     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at     TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE notifications
(
  id                  UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  owner_id            UUID REFERENCES owners (id) ON DELETE CASCADE,
  fine_id             UUID REFERENCES fines (id) ON DELETE CASCADE,
  sent_at             TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  notification_type   VARCHAR(50) NOT NULL,
  notification_status VARCHAR(50) NOT NULL,
  created_at     TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at     TIMESTAMP WITH TIME ZONE          DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX owners_license_number_index ON owners (license_number);

COMMENT ON TABLE owners IS 'Таблица для хранения информации о водителях';
COMMENT ON COLUMN owners.license_number IS 'регистрационный номер';

COMMENT ON TABLE vehicles IS 'Таблица для хранения информации об автомобилях';
COMMENT ON TABLE fines IS 'Таблица для хранения информации о штрафах';
COMMENT ON TABLE payments IS 'Таблица для хранения информации об оплатах штрафов';
COMMENT ON TABLE notifications IS 'Таблица для хранения информации об уведомлениях';

CREATE TRIGGER update_timestamp_owners BEFORE UPDATE ON owners FOR EACH ROW EXECUTE PROCEDURE moddatetime(updated_at);
CREATE TRIGGER update_timestamp_vehicles BEFORE UPDATE ON vehicles FOR EACH ROW EXECUTE PROCEDURE moddatetime(updated_at);
CREATE TRIGGER update_timestamp_fines BEFORE UPDATE ON fines FOR EACH ROW EXECUTE PROCEDURE moddatetime(updated_at);
CREATE TRIGGER update_timestamp_payments BEFORE UPDATE ON payments FOR EACH ROW EXECUTE PROCEDURE moddatetime(updated_at);
CREATE TRIGGER update_timestamp_notifications BEFORE UPDATE ON notifications FOR EACH ROW EXECUTE PROCEDURE moddatetime(updated_at);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS notifications CASCADE;
DROP TABLE IF EXISTS payments CASCADE;
DROP TABLE IF EXISTS fines CASCADE;
DROP TABLE IF EXISTS vehicles CASCADE;
DROP TABLE IF EXISTS owners CASCADE;
-- +goose StatementEnd
