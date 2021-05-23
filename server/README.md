Associations
Deploy

Tutorial Udemy React & Golang: A Practical Guide
https://github.com/myhendry/fiber_react_tutorial

/usr/local/mysql/bin/mysql -u root -p

show databases;

create database go_server;

CREATE USER 'user1'@'localhost' IDENTIFIED BY 'password';

// GRANT ALL PRIVILEGES ON \*.\* TO 'user1'@'localhost';
(\ as escape character only)


~/go/bin/realize start

password 1234

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

