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
| ID       | Username | Alias  | Password      | Is_Admin | Endpoints   |
| -------- | -------- | ------ | ------------- | -------- | ----------- |
| Main Key | String   | String | Hashed String | Bool     | One-to-many |

**Endpoints**
| ID       | UserID      | URL    | Is_Enabled |
| -------- | ----------- | ------ | ---------- |
| Main Key | Foreign Key | String | Bool       |

**Config**
| ID       | Key    | Value                  |
| -------- | ------ | ---------------------- |
| Main Key | String | Byte[] (After Marshal) |

#### InfluxDB Design

## Test Methods
