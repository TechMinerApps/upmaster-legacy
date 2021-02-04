# Basic Requirements

## Basic Functions
- Accept metrics from upmaster-agent and store them into database
- Dispatch HTTP Endpoints to upmaster-agent
- Provide API for upmaster-frontend to visualize data
- Alert user when certain contraint meets

## Platform Requirements

## Commit Message Convention
Use AngularJS style commit message.

# Development

## Rough Structure

## Detailed Design
### Database Table Design
Two database is used in this project: InfluxDB and SQLite. Since UpMaster is designed for small teams, account information should be handled well by SQLite, while time serires data is stored in InfluxDB.
#### SQLite Table Design

**Users**
| ID       | Username | Alias  | Password      | Email  | Is_Admin | Endpoints   | Alerts      | Alert_Channels |
| -------- | -------- | ------ | ------------- | ------ | -------- | ----------- | ----------- | -------------- |
| Main Key | String   | String | Hashed String | String | Bool     | One-to-many | One-to-many | One-to-many    |

**Endpoints**
| ID       | Name   | UserID      | URL    | Interval     | Is_Enabled |
| -------- | ------ | ----------- | ------ | ------------ | ---------- |
| Main Key | String | Foreign Key | String | Int (Second) | Bool       |

**Alert_Channels**
| ID       | Name   | UserID      | Type             | Config | Alerts      |
| -------- | ------ | ----------- | ---------------- | ------ | ----------- |
| Main Key | String | Foreign Key | Int (With Marco) | byte[] | One-to-many |

**Alerts**
| ID       | UserID      | Alert_Channel_ID | Status                  |
| -------- | ----------- | ---------------- | ----------------------- |
| Main Key | Foreign Key | Foreign Key      | Int (Alerting/Resolved) |

**Configs**
Used to store dynamic InfluxDB configuration
| ID       | Key    | Value                  |
| -------- | ------ | ---------------------- |
| Main Key | String | Byte[] (After Marshal) |

**OAuth**

Used to provide storage for OAuth Server

#### InfluxDB Design

Measurement: **up_status**

| Time | Is_Up            | Node             | EndpointID    |
| ---- | ---------------- | ---------------- | ------------- |
| -    | Bool (Field Key) | String (Tag Key) | Int (Tag Key) |

### Initialization Process
The initialization process is designed to be **idempotent**. It collects configuration from config file or environment variable and reconfigure UpMaster.

The process will do the following steps:
- Initialize SQLite: Using GORM operations to do database initialization, idempotent is achieved by which.
- Update Admin Info: Create/Update a admin user according to config file.
- Initialize InfluxDB: **influxdb-client-go** by InfluxDB Official is used as client. Database connection info is retrieved from SQLite. The database should already be create before this step. Measurement `up_status` will be created if not exists.

### API Design Principle

## Test Methods
