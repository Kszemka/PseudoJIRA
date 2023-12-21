
# PseudoJIRA

The project was developed as a part of the 'Przetwarzanie danych w Chmurach Obliczeniowych' course at AGH, in the year 2023. Serving as a 'proof of concept,' it takes the form of a straightforward task and user management application with a user-friendly interface, reminiscent of JIRA. The essence of the project centers on employing the Neo4j graph database.

The primary objective of this project was: firstly, to craft a user-friendly application for the efficient management of tasks and users; and secondly, to gain practical experience with graph databases, leading to the choice of Neo4j as the underlying database technology. The application offers users the ability to:

- retrieve a list of tasks with the option to filter them based on their assignment status, distinguishing between those assigned and unassigned, and categorizing them as either bug tasks or feature tasks,
- modify existing tasks to update their details,
- generate new tasks to be added to the database,
- remove tasks from the database when no longer relevant,
- access a list of all users stored in the database,
- explore all tasks assigned to a particular user and those tasks reported by the chosen user.





## Implementation

### Database

The structure can be conceptualized as a graph where nodes represent entities (e.g. tasks, users) and relationships represent connections between them. Key elements include:

![Database Schema](/images/model-db.png)


The database that was created can be represented by the following diagram:

![Database general example no.1](/images/visualisation.png)

Users interacting with the application can perform various operations on the database:

- Tasks can be retrieved and categorized as either Bugs or Features based on the BELONGS_TO relationship,
- Users can view a list of tasks assigned to other users by following the ASSIGN relationship,
- Users can check tasks that was reported and by whom, by the REPORT relationship,
- Users have the capability to create, delete, and update task nodes in the graph. This involves adding new nodes, removing existing ones, and modifying attributes of existing tasks.

While the provided placeholders don't show an actual visualization, the following example visualization clearly presents nodes representing users and tasks connected by various relationships, how tasks are categorized (Bugs or Features), assigned to users, and reported by users.

![Database example no.2](/images/visualisation-complex.png)


### Backend

The backend of the application is developed in Go and consists of two key packages: `gorilla/mux` - for managing HTTP requests and `neo4j-go-driver` for database connectivity, specifically handling interactions with a Neo4j database. The backend functionality centers around basic CRUD operations.

Key components of the backend include:

- `entities`package defines models for database objects, as an example:
```
type User struct {
	Surname  string `json:"surname"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
```
- `controllers` package contains logic for handling HTTP requests and generating appropriate responses. An example:

```
func GetAllUsers(w http.ResponseWriter, request *http.Request) {
	users, err := services.ExecuteGetUsersQuery()
	if err != nil {
		log.Printf("Error retrieving tasks from the database: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	config.WriteHttpResponse(w, users)
}
```

- `services` package encapsulates additional business logic, working with models and mapping database results. For instance:

```
func ExecuteGetUsersQuery() ([]models.User, error) {
	db := config.GetDbConnection()
	var users []models.User

	result, err := database.RunQuery(db, database.GET_USERS_QUERY, nil)
	if err != nil {
		return nil, err
	}

	var user models.User

	for result.Next(db.Ctx) {
		userRecord := result.Record()
		u, _ := userRecord.Get("user")
		result, _ := json.Marshal(u.(neo4j.Node).Props)
		json.Unmarshal(result, &user)

		users = append(users, user)
	}

	return users, nil
}
```
The example illustrates the process of mapping records from the database to the defined Go type 'User,' resulting in a slice of 'User' types.

The `web` package houses all the logic related to handling HTTP requests for each endpoint, keeping it distinct from other components.

In essence, the backend uses the Gorilla Mux router for routing HTTP requests and the Neo4j Go driver for database connectivity. The logical components (entities, controllers, and services) collaborate to provide a seamless and modular structure for handling HTTP requests and interacting with the Neo4j database. 
The separation of concerns is evident in the distinct packages, enhancing code organization and maintainability.

### Frontend

Frontend, which was created with React.js, is SPA (Single Page Application). SPAs load a single HTML page and dynamically update the content as the user interacts with the application, without requiring a full page reload. In this case, the component handles the rendering of either a task list or a user list based on the selected tab.

The component fetches data from a backend API using the fetch function. It interacts with the backend at dedicated endpoints to retrieve task and user information, using hooks.

React hooks are functions that enable functional components to have state and lifecycle features that were previously only available in class components. 

The component supports basic CRUD (Create, Read, Update, Delete) operations for tasks. Users can create new tasks, edit existing tasks, assign tasks to users, and delete tasks. These operations involve making corresponding API requests to the backend.


## Demo

![App demo](/images/demo.gif)

