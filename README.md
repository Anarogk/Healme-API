

# 🏥 HealthTech Wizardry API 🧙‍♂️

Welcome to the most magical healthcare API this side of Hogwarts! We've combined the healing powers of Go with the sorcery of modern web technologies to create an API so powerful, it might just cure the common cold (disclaimer: it won't).

## 🌟 Features

- **Appointment Scheduling**: Because time-turners are so last century.
- **Patient Management**: Keep track of your patients better than Dumbledore kept track of Harry.
- **Doctor Management**: Organize your healers like McGonagall organizes her class schedule.
- **User Authentication**: More secure than Gringotts, less grumpy than goblins.

## 🚀 Getting Started

### Prerequisites

- Go (version 1.16+)
- PostgreSQL (because even wizards need reliable databases)
- A sense of humor (critical for debugging)

### Installation

1. Clone the repo:
   ```bash
   git clone https://github.com/your-username/healthtech-wizardry-api.git
   ```
2. Enter the magical directory:
   ```bash
   cd healthtech-wizardry-api
   ```
3. Cast the spell of dependency installation:
   ```bash
   go mod tidy
   ```
4. Create a `.env` file and fill it with your secret incantations (database credentials, JWT secret, etc.)

5. Wave your wand (or just type in your terminal):
   ```bash
   go run cmd/server/main.go
   ```

Congratulations! Your API is now running faster than a Nimbus 2000!

## 🧙‍♀️ API Documentation

### Authentication Spells 🔐

#### Register a New User
```bash
POST /api/register
```
Body:
```json
{
  "email": "harry.potter@hogwarts.edu",
  "password": "ExpectoAPIum!",
  "role": "patient"
}
```

#### Login
```bash
POST /api/login
```
Body:
```json
{
  "email": "harry.potter@hogwarts.edu",
  "password": "ExpectoAPIum!"
}
```

### Patient Charms 🧑‍⚕️

#### Get All Patients
```bash
GET /api/patients
```

#### Get Single Patient
```bash
GET /api/patients/:id
```

#### Update Patient
```bash
PUT /api/patients/:id
```

### Doctor Enchantments 👩‍⚕️

#### Get All Doctors
```bash
GET /api/doctors
```

#### Get Single Doctor
```bash
GET /api/doctors/:id
```

#### Update Doctor
```bash
PUT /api/doctors/:id
```

### Appointment Alchemy ⏰

#### Create Appointment
```bash
POST /api/appointments
```
Body:
```json
{
  "patient_id": 1,
  "doctor_id": 2,
  "date_time": "2023-05-01T14:30:00Z",
  "description": "Wand arm feeling a bit wonky"
}
```

#### Get All Appointments
```bash
GET /api/appointments
```

#### Update Appointment
```bash
PUT /api/appointments/:id
```

#### Delete Appointment
```bash
DELETE /api/appointments/:id
```

## 📊 Schema Flowchart

```
+----------------+       +----------------+
|     Users      |       |   Patients     |
+----------------+       +----------------+
| id             |------>| id             |
| email          |       | user_id        |
| password       |       | first_name     |
| role           |       | last_name      |
+----------------+       | date_of_birth  |
                         | gender         |
                         | phone_number   |
                         | address        |
                         +----------------+
                                 ^
                                 |
                                 |
                         +----------------+
                         |   Doctors      |
                         +----------------+
                         | id             |
                         | user_id        |
                         | first_name     |
                         | last_name      |
                         | specialization |
                         | license_number |
                         +----------------+
                                 ^
                                 |
                         +----------------+
                         | Appointments   |
                         +----------------+
                         | id             |
                         | patient_id     |
                         | doctor_id      |
                         | date_time      |
                         | description    |
                         | status         |
                         +----------------+
```

## 🧪 Testing

Run tests faster than Snape takes points from Gryffindor:

```bash
go test ./...
```

## 🚢 Deployment

1. Summon a Docker container:
   ```bash
   docker-compose up --build
   ```
2. Deploy to the cloud using Vercel, because even APIs need a little magic in the sky.

## 🧙‍♂️ Contributing

Contributions are welcome! Just make sure your code is cleaner than Filch's mop and more organized than Hermione's study schedule.

## 📜 License

This project is licensed under the Spell-It-Yourself License. Use it wisely, and may your code always compile on the first try!

