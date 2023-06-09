# Product API
### ERD
<img src="erd.png">

### Instructions
- Please fork this repository to your private github repo.
- Create branch with name `<your_sikm_id>_<name>`.
- Push your changes to your branch.
- For any "challenge/homework", mentee can submit PR to branch `master` of this repo (upstream). Any MR submitted before deadline will get an extra score on "nilai keaktifan" for subject "Web Application".

### Session 1
Objectives during session:
- [ ] Intro [Gin framework](https://github.com/gin-gonic/gin/blob/master/docs/doc.md)
- [ ] Gin router
- [ ] Gin [model binding and validation](https://gin-gonic.com/docs/examples/binding-and-validation/)
- [ ] Clean architecture
- [ ] Concept of [DTO](https://en.wikipedia.org/wiki/Data_transfer_object)
- [ ] SQL error handling with [pgconn.PgError](https://github.com/jackc/pgx/blob/master/pgconn/errors.go)

#### Feat 1: Get list products
- Should return status code `200` when success select data from DB.
- Response should contain list product with their category.

API Contract: `GET /product/list`, sample response:
```
[
    {
        "id": 1,
        "name": "Beng-beng",
        "price": 2500,
        "category_id": 1,
        "posted_by": 1,
        "category": {
            "id": 1,
            "name": "Snack"
        }
    },
    {
        "id": 2,
        "name": "Top",
        "price": 1500,
        "category_id": 1,
        "posted_by": 1,
        "category": {
            "id": 1,
            "name": "Snack"
        }
    }
]
```

#### Feat 2: Get product detail
- Should return status code `200` when product with given id found in database.
- Product detail should get product data with it's category and the user/creator of the product.
- Should return status code `404` when product with given id not found in database.

API Contract: `GET /product/:id`, sample response:
```
{
    "id": 1,
    "name": "Beng-beng",
    "price": 2500,
    "poster": {
        "id": 1,
        "name": "John Doe",
        "role": "ADMIN",
        "email": "john.doe@mail.com"
    },
    "category": {
        "id": 1,
        "name": "Snack"
    }
}
```

#### Feat 3: Post product
- Should return status code `201 (Created)` when insert data success.
- Request payload should be validated with [gin validator](https://gin-gonic.com/docs/examples/binding-and-validation/), the requirement is on  payload sample below. When the validation error should return status code `400`.
- Product's name should not contains words like "termurah", "terbaik", "diskon", and "promo". 
- When `category_id` or `posted_by` id is not exist in DB (ForeignKeyViolation sql error state `23503`), it should return error status code 400. For others postgres error code can be read [here](https://www.postgresql.org/docs/current/errcodes-appendix.html).

API Contract: `POST /product/add`, sample response:
```
{
    "message": "product added"
}
```

Sample request payload and the validation:
```
{
    "name": "Ayam Geprek", // required, min length 8 chars
    "price": 7000, // required, price more than 0
    "category_id": 1, // required
    "posted_by": 1 // required
}
```

#### [CHALLENGE] Feat 4: Update product by ID
- Should be able to update data in product include these fields: name, price, and category_id. Please create new `DTO` for the request payload.
- Product's name should not contains words like "termurah", "terbaik", "diskon", and "promo". 
- Should return status code `200` when update success.
- When product with the given `id` not found in DB, API should return response status code `404`.
- When `category_id` is not exist in DB (ForeignKeyViolation sql error state `23503`), it should return error status code 400. For others postgres error code can be read [here](https://www.postgresql.org/docs/current/errcodes-appendix.html).

API Contract: `PUT /product/update/:id`, sample response:
```
{
    "message": "product updated"
}
```

Sample request payload and the validation:
```
{
    "name": "Ayam Geprek", // required, min length 8 chars
    "price": 7000, // required, price more than 0
    "category_id": 1 // required
}
```

#### [CHALLENGE] Feat 5: Delete product by ID
- Should return status code `200` when delete success.
- When product with the given `id` not found in DB, API should return response status code `404`.

API Contract: `DELETE /product/delete/:id`, sample response:
```
{
    "message": "product deleted"
}
```

### Session 2: Middleware & JWT Authentication
Session based authentication:
<img src="session-based.png">

Token based (JWT) authentication:
<img src="token-based.png">

Ref: 
- https://jerrynsh.com/all-to-know-about-auth-and-cookies/#:~:text=JWT%20is%20simply%20a%20token,within%20your%20browser's%20Cookies%20storage.
- https://hackernoon.com/using-session-cookies-vs-jwt-for-authentication-sd2v3vci

Objectives during session:
- [ ] Understand concept of middleware
- [ ] Implement [middleware](https://github.com/gin-gonic/gin/blob/master/docs/doc.md#using-middleware) in Gin
- [ ] Understand [JWT authentication](https://jwt.io/introduction)
- [ ] Implement JWT authentication and integrate with cookie
- [ ] Implement JWT validation
- [ ] Using gin context to carry data between middleware in a request

#### Feat 1: Login and generate JWT token
- Should response cookie with key `session_token`, active for 15 minutes, and `http only` cookie when email and password is valid.
- JWT payload should contain required fields (check below JWT payload example), we will use package [go-jwt](github.com/golang-jwt/jwt) and follow [generate-token snippet](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-New-Hmac).
- Should validate request body with gin validator.
- Should return error `400` when user by email is not found.
- Should return error `400` when password is wrong.

API Contract: `POST /user/login`, sample response:
```
// Response Body
{
    "message": "login success"
}
```

Sample request payload and the validation:
```
{
    "email": "john.doe@mail.com", // required, valid email
    "password": "123456", // required, min length 6
}
```

Sample JWT payload:
```
{
  "iss": "product-api.com",
  "user_id": "1",
  "name": "John Doe",
  "email": "john.doe@mail.com",
  "iat": 1618985752,
  "exp": 1618986952,
  "scope": "user admin"
}
```

#### Feat 2: Token validation
Below is the Authorization rules/access on our API base on user's role:
- Unauthorized: client don't need to authenticated to access this API
    - `POST /user/login`
- Authorized User: client must be logged in and have role/scope `USER` or `ADMIN`.
    - `GET /product/list`
    - `GET /product/:id`
- Authorized Admin: client must be logged in and have role `ADMIN`.
    - `POST /product/add`
    - `PUT /product/:id`
    - `DELETE /product/:id`

Read docs on [how to validate token](https://pkg.go.dev/github.com/golang-jwt/jwt/v5#example-Parse-Hmac).

### Session 3: Deployment 

During this session we will deploy our service to [railway](https://railway.app/).

#### Steps to follows:
- Move configuration to env variables, we will use [godotenv](https://github.com/joho/godotenv) to read env from `.env` file.
- Pull to the latest changes on `master` branch or fork this repository with the [latest commit](https://github.com/awidiyadew/sikm-product-api/commit/adb67ad48184370e9878278dc29d6c4a66161056).
- Signup to [railway](https://railway.app/) with your github account.
- Create a new project on railway
- Create a database project with PostgreSQL on railway
- Create a new service and connect to your Github repo `product-api`
- Setup the environment variables on your `product-api` service so it can connect to the PostgreSQL.
- Setup free railway domain.
- Test the API with Postman.

### Bonus - Swagger Docs Tutorials
Swagger is an API documentation tools that follow OPEN API specs. Below is steps to setup swagger docs with swagger-ui.

1. Create your swagger yaml with the online [swagger editor](https://editor.swagger.io/), it's ery easy to create the docs, you can copy paste and edit the example from the `Swagger Petstore` example. After editing is done, then click `File` and `Save as YAML`.
1. Rename the downloaded `yaml` file to `swagger.yaml`.
1. Download latest release of the swagger-ui [here](https://github.com/swagger-api/swagger-ui/releases), please download as zip.
1. Extract the downloaded file, copy all the contents of `dist` directory to your project with directory name `swagger-ui` on your root project.
1. Copy file `swagger.yaml` from step no 2 to directory `swagger-ui` in your project.
1. Open file `swagger-ui/swagger-initializer.js` and update the options `url` with values:
    ```
    window.ui = SwaggerUIBundle({
        url: "swagger.yaml", // path to yaml file
        // ...rest of the code
    ```
1. Open `main.go` and register a static route to the `swagger-ui` directory.
    ```
    // register API `/docs` for the swagger-ui
    router.Static("/docs", "swagger-ui")
    ```
1. Run the apps and try to access your api on path `/docs`!.