# Tutorial Udemy React & Golang: A Practical Guide
[Github Repo](https://github.com/myhendry/fiber_react_tutorial)

- [x] Belong To
- [x] Has One
- [x] One-to-Many
- [x] No Record Found Error Handling
- [x] Authentication
- [x] Middleware - IsAuthenticated
- [x] Get UserID from Auth
- [ ] Query
- [ ] SQL Builder
- [x] Time X
- [ ] Many-to-Many
- [x] Cascade Delete
- [ ] Cascade Update 
- [ ] Authorization
- [ ] Deploy

---
### ASSOCIATIONS
Belongs To (User - Role) Role Belongs to User
Create Role
Create User with role_id

Has One (User - Profile) User Has One Profile
Create User
Create Profile with user_id
### DATABASE
/usr/local/mysql/bin/mysql -u root -p

show databases;

create database go_server;

CREATE USER 'user1'@'localhost' IDENTIFIED BY 'password';

// GRANT ALL PRIVILEGES ON \*.\* TO 'user1'@'localhost';
(\ as escape character only)

~/go/bin/realize start

password 1234

```javascript

.realize.yaml

settings:
  legacy:
    force: false
    interval: 0s
schema:
  - name: covid
    path: .
    commands:
      install:
        status: true
        method: go build -o /Users/hendrylim/coding/lab/go/covid/server/build
      run:
        status: true
        method: /Users/hendrylim/coding/lab/go/covid/server/build
    watcher:
      extensions:
        - go
      paths:
        - /
      ignored_paths:
        - .git
        - .realize
        - vendor

```
