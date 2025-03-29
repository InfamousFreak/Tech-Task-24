# Food_Delivery_Website




## Stack_used

Backend - Golang; GORM; GOFiber


## Features

- JWT Auth, Google Auth for customers and Restaurants Admins
- Food menu pages
- Profile page, cart page
- Suggestions of dishes will be shown
- Food to be searched based on tags
- Payment gateway integration (virtual for now)
- 

## Setup Instructions
### **1. Install Golang**
1. Download and install Go from [https://golang.org/dl/](https://golang.org/dl/)
2. Verify installation:
   ```sh
   go version
   ```

### **2. Set Up PostgreSQL Database**
1. Install PostgreSQL:
   - **Linux** (Ubuntu/Debian):
     ```sh
     sudo apt update && sudo apt install postgresql postgresql-contrib
     ```
   - **MacOS**:
     ```sh
     brew install postgresql
     ```
   - **Windows**:
     Download and install from [https://www.postgresql.org/download/](https://www.postgresql.org/download/)
2. Start PostgreSQL service:
   ```sh
   sudo service postgresql start
   ```
3. Create a new database:
   ```sh
   sudo -u postgres psql
   CREATE DATABASE food_delivery;
   ```
4. Set up a PostgreSQL user:
   ```sh
   CREATE USER food_admin WITH ENCRYPTED PASSWORD 'securepassword';
   ALTER ROLE food_admin SET client_encoding TO 'utf8';
   ALTER ROLE food_admin SET default_transaction_isolation TO 'read committed';
   ALTER ROLE food_admin SET timezone TO 'UTC';
   GRANT ALL PRIVILEGES ON DATABASE food_delivery TO food_admin;
   ```

### **3. Install GoFiber & GORM**
```sh
go mod init food_delivery

go get github.com/gofiber/fiber/v2
go get gorm.io/gorm
go get gorm.io/driver/postgres
```

### **4. Install Authentication Libraries**
```sh
go get github.com/dgrijalva/jwt-go
go get golang.org/x/oauth2
go get golang.org/x/oauth2/google
```

### **5. Clone the Repository & Setup Backend**
```sh
git clone https://github.com/yourusername/food_delivery.git
cd food_delivery
```
- **Create a `.env` file** with:
  ```env
  DB_HOST=localhost
  DB_USER=food_admin
  DB_PASSWORD=securepassword
  DB_NAME=food_delivery
  JWT_SECRET=your_secret_key
  GOOGLE_CLIENT_ID=your_google_client_id
  GOOGLE_CLIENT_SECRET=your_google_client_secret
  ```
- **Run the server**:
  ```sh
  go run main.go
  ```

### **6. Set Up Frontend**
1. Navigate to `frontend/` folder:
   ```sh
   cd frontend
   ```
2. Open `index.html` in a browser:
   ```sh
   open index.html  # MacOS
   start index.html  # Windows
   xdg-open index.html  # Linux
   ```

### **7. Running the Full Application**
- **Start the backend**:
  ```sh
  go run main.go
  ```
- **Open the frontend** in a browser:
  ```sh
  http://localhost:8080/homepage.html
  ```

## API Endpoints
| Method | Endpoint | Description |
|--------|----------------|-------------|
| POST   | `/api/auth/signup` | User signup |
| POST   | `/api/auth/login` | User login |
| GET    | `/api/menu` | Fetch menu items |
| POST   | `/api/cart/add` | Add item to cart |
| POST   | `/api/order/place` | Place order |

## To-Do (Future Enhancements)
- WebSockets for live updates
- Real-time order tracking
- Deploying to cloud (Render, AWS, or DigitalOcean)

---
## Contributors
- **Your Name** - [GitHub](https://github.com/InfamousFreak)

---
## License
MIT License
make this the requested format 

