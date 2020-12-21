# cubicasa_test
Cubicasa Assigment
# Assigment
You are building a functioning web service with Golang. The primary
resources to be used in this service are called “Hub”, “Team”, and
“Users” in the order of hierarchy. A short description is as below.
A Hub is an entity that associates Team depending on their geological
location.
A Team is an entity that associates Users based on their types.
A Users is an entity that holds the information of the human users.
Your job is to utilise those resources for building a system and provide
relevant information to the requests. A list of requirements/tasks are
as below.
• Build system in Golang, object-oriented style.
• Implement a Create for each hub, team, and users.
• Implement a Search which will return team and hub information.
• Implement a Join for users and team into hub (for the simplicity,
multiple affiliations is not allowed in this project).
• Provide a SQL script which creates tables needed for the web
service.
• Use Postgres for the database.
• Good to use docker/docker-compose for local development setup(not mandatory)

# Preparing
  - install postgresdb
  - modify cubicasa_test.config set userlogin database. if doest exist file, pls run go run ./server/main.go and file will be generate with default config
--------------
# How to run project
cd server
go run main.go
--------------
# JSON Structure
  - HUB: {"hub_id": "string", "name": "string", "geo_location": "string"}
  - TEAM: {"team_id": "string", "name": "string", "type": "string", "hub_id": "string"}
  - USER: {"user_id": "string", "role": "string", "email": "string", "team_id": "string"}
# API list
• Implement a Create for each hub, team, and users.
  - hub:  http://localhost:3100/hubs [POST]
  - team: http://localhost:3100/teams [POST]
  - user: http://localhost:3100/users [POST]
  
• Implement a Search which will return team and hub information.
  - http://localhost:3100/searchhubs?name="value" [GET]
  - http://localhost:3100/searchteams?name="value" [GET]

• Implement a Join for users and team into hub.
  - http://localhost:3100/jointeam/{team_id}/{hub_id} [PUT]
  - http://localhost:3100/joinuser/{user_id}/{team_id}/{role_id} [PUT]

• Provide a SQL script which creates tables needed for the web service.
  - http://localhost:3100/install [GET]
with script install insite folder ./setup/script.sql
  - http://localhost:3100/checkdb [GET]
check db exsited or not

• Use Postgres for the database. (yes)

• Good to use docker/docker-compose for local development setup(not mandatory):
  - build: cd server -> make linux -> make docker -> docker-compose up -d
  - continue test with api under
