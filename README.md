# User management 

This a example project to manage a user in a app
##### technology used :
  - Golang language use to write the code
  - Gorm is a fantastic ORM library for Golang you can get in [here](https://gorm.io/)
  - Mux framework for router
  - JWT (Json web token)

# Features!
##### - admin :
 - add user or register new admin
 - delete user
 - view detail user
 - update profile user
 - upgrade user to admin
  - upload photo
##### - user :
  - register
  - delete account
  - update profile
  - upload photo

# Api documentation
this is api's documentation if you want to use my h3h3h3 script
##### #Method POST
 - [REGISTER] url : http://localhost:8080/register
    ```json
    {
    	"name": "carmudi",
    	"email": "carmudi56sss@gmail.com",
    	"password": "123456",
    	"no_tlpn": "08881506482",
    	"gender": "laki-laki",
    	"address": "jln.Kepongpong",
    	"role":"user"
    }
    ```
 - [LOGIN] url : http://localhost:8080/login
    ```json
    {
	    "email":"carmudi@mail.com",
	    "password":"123456"
    }
    ```
 - [PHOTO] url : http://localhost:8080/user/photo/{id}
 - NB: I used multipart-file request to upload the profile photo so you can see my screenshot for it
![](https://i.ibb.co/C1WKk5Q/Screen-Shot-2020-02-27-at-15-48-10.png)

##### #Method GET
- [VIEW ALL] url : http://localhost:8080/user
- Output :
    ```json
    {
	"data": [
		{
			"id": 1,
			"create_at": "2020-02-26T14:27:44.893511+07:00",
			"update_at": "2020-02-26T23:58:19.734921+07:00",
			"delete_at": null,
			"name": "Carmudi",
			"email": "carmudi@n45ht.or.id",
			"no_tlpn": "08888888",
			"gender": "laki-laki",
			"address": "Jln. H. Dahlan, Ragunan, Jakarta, indonesia",
			"role": "admin",
			"photo": "default.jpeg"
		},
		{
			"id": 2,
			"create_at": "2020-02-26T14:28:19.327105+07:00",
			"update_at": "2020-02-27T00:12:25.896243+07:00",
			"delete_at": null,
			"name": "Api al Rahman",
			"email": "apirahman@gmail.com",
			"no_tlpn": "088899999",
			"gender": "laki-laki",
			"address": "Jln.H.Dahlan, Ragunan, Jakarta",
			"role": "user",
			"photo": "default.jpeg"
		}
	],
	"message": "Success",
	"status": true
    }
    ```
- [VIEW USER] url : http://localhost:8080/user/{id}
- Output :
    ```json
    {
	"data": {
		"id": 2,
		"create_at": "2020-02-26T14:56:07.853569+07:00",
		"update_at": "2020-02-26T21:16:09.555645+07:00",
		"delete_at": null,
		"name": "Api al Rahman",
		"email": "apirahman@gmail.com",
		"no_tlpn": "088899999",
		"gender": "laki-laki",
		"address": "jln.Kepongpong",
		"role": "user",
		"photo": "default.jpeg"
	},
	"message": "Success",
	"status": true
    }
    ```

##### #Method Delete
- [DELETE] url : http://localhost:8080/user/{id}
- Output :
    ```json
    {
        "data": null,
        "message": "Success delete account",
        "status": false
    }
    ```
    
##### #Method PUT
- [UPDATE] url : http://localhost:8080/user/{id}
- Request :
    ```json
    {
    	"name": "carmudi",
    	"email": "carmudi@n45ht.or.id",
    	"password": "123456",
    	"no_tlpn": "0888888888",
    	"gender": "laki-laki",
    	"address": "Jln. H. Dahlan, Ragunan, Jakarta, indonesia",
    	"role":"admin"
    }
    ```
- Output :
    ```json
    {
      "data": {
        "id": 2,
        "create_at": "2020-02-26T14:28:19.327105+07:00",
        "update_at": "2020-02-27T11:13:29.309942+07:00",
        "delete_at": null,
        "name": "carmudi",
        "email": "carmudi88@gmail.com",
        "no_tlpn": "08881506482",
        "password": "$2a$10$bqJRKrUOwCPwhyhKDLB/duyVQbkHiy4MjZionSd9.AU0sYcQSHS0",
        "gender": "laki-laki",
        "address": "Jln. H. Dahlan, Ragunan, Jakarta, indonesia",
        "role": "admin",
        "photo": "photo_639663_271014522020.jpg"
    },
    "message": "Update successfully",
    "status": true
    }
    ```

### #Installation

In first time you want to install golang language to run this script and
Install the dependencies and devDependencies and start the server.

```sh
$ cd example-user_management-with-golang-and-gorm
$ go run .
```

## License
 - --
 >:v
