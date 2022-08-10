# Demo-CRUD---REST-API-in-Go-using-Fiber-GORM
Tranning biết các sử dụng các package để connect db và curd cơ bản 

-------------------

Get driver Sqlite: go get "gom.io/driver/sqlite"
Get Gorm: go get "gorm.io/gorm"
Get Fiber: go get "github.com/gofiber/fiber/v2"

-------------------

Cài đặt Extensions SQLite on VS Code

-------------------

run project: air

-------------------

* API Postman:
1 Create Page: 
URL: http://localhost:3000//api/pages
Method: POST
Parameter: 
    "body":     "body",
    "id_page":   "idpage"
    
2 Get Page:
URL: http://localhost:3000//api/pages
Method: GET
Parameter: No

3 Get Single Page:
URL: http://localhost:3000//api/pages/id
Method: GET
Parameter: No

4 Update Page:
URL: http://localhost:3000//api/pages/id
Method: PUT
Parameter: 
    "body":     "body",
    "id_page":   "idpage"
    
5 Delete Page:
URL: http://localhost:3000//api/pages/id
Method: Delete
Parameter: No
