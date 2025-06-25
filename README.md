![Screenshot 2025-06-25 133546](https://github.com/user-attachments/assets/932c2d2f-15ac-4f44-841f-89ee5f4b6d05)

# GoLang_Assessment

This is a Golang-based web application for a hospital management system featuring a **Receptionist Portal** and a **Doctor Portal**. The system supports user authentication, patient registration, and role-based access to patient data.

---

## 🛠️ Tech Stack 👨‍💻

- **Backend**: Golang (Gin Framework)
- **Database**: PostgreSQL (using GORM ORM)
- **Authentication**: Simple token-based authentication (for demo purposes)

---
### API Documentation Link 📚

https://www.postman.com/joint-operations-geoscientist-68278105/golang-crud-api-documentation/collection/qj2usv1/api-documentation-for-golang-assessment?action=share&creator=15591734



## 🔐 Authentication

### Login (for both Receptionist and Doctor)

**Endpoint:** `POST /login`  
**Purpose:** Authenticate user and return token.

#### Request Body

```json
{
  "username": "recep1",
  "password": "pass123",
  "role": "receptionist"
}
```
#### Success Response  

```json
{
  "token": "simple-token-recep1-receptionist",
  "role": "receptionist"
}
```
### Login for (Doctor)

#### Request Body ✅

```json
{
    "username": "doctor1",
    "password": "pass456",
    "role": "doctor"
}
```

#### Success Response ✅ 

```json
{
    "token": "simple-token-doctor1-doctor",
    "role": "doctor"
}
```
### Screen Short (Proof of Work) ✅👨‍💻

![Screenshot 2025-06-25 133546](https://github.com/user-attachments/assets/446660ad-6ed8-45eb-ad24-7615cff79d45)
![Screenshot 2025-06-25 133456](https://github.com/user-attachments/assets/b28a0870-fe4e-4443-8ad2-5ad13f0d3856)
![Screenshot 2025-06-25 133429](https://github.com/user-attachments/assets/59c509ad-b8f4-48d7-b768-f92b84810080)
![Screenshot 2025-06-25 133346](https://github.com/user-attachments/assets/a0860982-a2bf-4c9e-88af-1853194d2f0f)
