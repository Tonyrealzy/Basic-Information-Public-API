# Golang Public API

## 📌 Project Description

This is a simple public API built with **Golang** that returns:

- Your **email address**
- The **current datetime** in ISO 8601 format
- The **GitHub repository URL** of this project

This API is **CORS-enabled** and serves JSON responses over HTTP. It accepts only `GET` requests.

---

## 🚀 Getting Started

### **Prerequisites**

Ensure you have the following installed:

- **Go 1.23.5** ([Download Go](https://go.dev/dl/))
- **Git** ([Download Git](https://git-scm.com/downloads))

---

### **📥 Setup & Run Locally**

1️⃣ **Clone the Repository**

```sh
  git clone https://github.com/Tonyrealzy/Basic-Information-Public-API
  cd Basic-Information-Public-API
```

2️⃣ **Install Dependencies**

```sh
  go mod tidy
```

3️⃣ **Run the API**

```sh
  go run main.go
```

4️⃣ **Test Locally**
Open your browser or run:

```sh
curl -X GET http://localhost:8080/api/v1/basic-info
```

---

## 📖 API Documentation

### **🌍 Endpoint**

```plaintext
GET /api
```

### **📤 Response Format**

| Field        | Type   | Description                        |
| ------------ | ------ | ---------------------------------- |
| `email`      | string | Your email address                 |
| `datetime`   | string | Current datetime (ISO 8601 format) |
| `github_url` | string | GitHub repository URL              |

---

### **🛠 Example Response**

#### **✅ Successful Response**

```json
{
  "email": "umehobiarinze2@gmail.com",
  "datetime": "2025-01-29T12:00:00Z",
  "github_url": "https://github.com/Tonyrealzy/Basic-Information-Public-API"
}
```

---

## 🔗 Useful Resources

- **Hire Golang Developers:** [hng.tech/hire/golang-developers](https://hng.tech/hire/golang-developers)
- **Go Documentation:** [golang.org/doc](https://golang.org/doc/)

---

## 🌍 Live API Deployment (Render.com)

Once deployed, access your API at:

````plaintext
https://basic-information-public-api.onrender.com
```

---

## 📜 License
This project is licensed under the **MIT License**. Feel free to use and modify it.

---

## 🛠 Contributing
Pull requests are welcome! If you'd like to contribute, please **fork the repository**, create a new branch, and submit a PR.

---

### **📧 Contact**
For any questions or suggestions, reach out via email: `your.email@example.com`.

---

**🚀 Happy Coding!** 🎉

````
