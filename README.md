
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
```sh
   git clone https://github.com/Anarogk/healme-API.git
```

2. Enter the magical directory:
```sh
   cd healme-API
```

3. Cast the spell of dependency installation:
```sh
   go mod tidy
```

4. Create a `.env` file and fill it with your secret incantations (database credentials, JWT secret, etc.)

5. Wave your wand (or just type in your terminal):
```sh
   go run cmd/server/main.go
```

Congratulations! Your API is now running faster than a Nimbus 2000!

## 🧙‍♀️ API Documentation

### Authentication Spells 🔐

#### Register a New User
```sh
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
```sh
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
```sh
GET /api/patients
```

#### Get Single Patient
```sh
GET /api/patients/:id
```

#### Update Patient
```sh
PUT /api/patients/:id
```

### Doctor Enchantments 👩‍⚕️

#### Get All Doctors
```sh
GET /api/doctors
```

#### Get Single Doctor
```sh
GET /api/doctors/:id
```

#### Update Doctor
```sh
PUT /api/doctors/:id
```

### Appointment Alchemy ⏰

#### Create Appointment
```sh
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
```sh
GET /api/appointments
```

#### Update Appointment
```sh
PUT /api/appointments/:id
```

#### Delete Appointment
```sh
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

```sh
go test ./...
```

## 🚢 Deployment

1. Summon a Docker container:
```sh
   docker-compose up --build
```

2. Deploy to the cloud using Vercel, because even APIs need a little magic in the sky.

## 🧙‍♂️ Contributing

Contributions are welcome! Just make sure your code is cleaner than Filch's mop and more organized than Hermione's study schedule.

## 📜 License

This project is licensed under the Spell-It-Yourself License. Use it wisely, and may your code always compile on the first try!


