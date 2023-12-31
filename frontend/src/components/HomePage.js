import React, { useEffect, useState, useCallback } from "react";
import "./HomePage.css";

const HomePage = () => {
  const [selectedTab, setSelectedTab] = useState("");
  const [showEditModal, setShowEditModal] = useState(false);
  const [edittaskName, setEdittaskName] = useState(null);
  const [newName, setNewName] = useState("");
  const [newDescription, setNewDescription] = useState("");
  const [taskList, setTaskList] = useState([]);
  const [userList, setUserList] = useState([]);
  const [showAssignModal, setShowAssignModal] = useState(false);
  const [selectedUser, setSelectedUser] = useState(null);
  const [showDeleteModal, setShowDeleteModal] = useState(false);
  const [showTasksModal, setShowTasksModal] = useState(false);
  const [tasksForUser, setTasksForUser] = useState([]);
  const [selectedUsername, setSelectedUsername] = useState(null);

  const [showCreateTaskModal, setShowCreateTaskModal] = useState(false);
  const [createTaskName, setCreateTaskName] = useState("");
  const [createTaskDescription, setCreateTaskDescription] = useState("");
  const [createTaskReporter, setCreateTaskReporter] = useState("");
  const [createTaskCategory, setCreateTaskCategory] = useState("");
  const [filterByBug, setFilterByBug] = useState(false);
  const [filterByFeature, setFilterByFeature] = useState(false);
  const [filterByAssigned, setFilterByAssigned] = useState(false);
  const [filterByUnassigned, setFilterByUnassigned] = useState(false);
  const [filterAssignButton,setButtonAssign] = useState(null)

  useEffect(() => {
    fetchUserList();
  }, []);

  const fetchTaskList = useCallback(async () => {
    try {
      let endpoint = "/api/all";

      if (filterByBug) {
        endpoint = "/api/bugs";
      } else if (filterByFeature) {
        endpoint = "/api/features";
      } else if (filterByAssigned) {
        endpoint = "/api/assigned";
      } else if (filterByUnassigned) {
        endpoint = "/api/unassigned";
      }

      const response = await fetch(`${endpoint}`);

      if (response.ok) {
        const data = await response.json();
        setTaskList(data);
      }
    } catch (error) {
      console.error("Error fetching task list:", error);
    }
  }, [filterByBug, filterByFeature, filterByAssigned, filterByUnassigned]);

  useEffect(() => {
    fetchTaskList();
  }, [
    filterByBug,
    filterByFeature,
    filterByAssigned,
    filterByUnassigned,
    fetchTaskList,
  ]);

  const fetchUserList = async () => {
    try {
      const response = await fetch("/api/users");
      if (response.ok) {
        const data = await response.json();
        setUserList(data);
      }
    } catch (error) {
      console.error("Error fetching user list:", error);
    }
  };

  const handleTabClick = (tab) => {
    setSelectedTab(tab);
    if (tab === "task") {
      fetchTaskList();
    } else if (tab === "users") {
      fetchUserList();
      setUserList([]);
    }
  };

  const handleFilterChange = (filter) => {
    switch (filter) {
      case "bug":
        setFilterByBug(!filterByBug);
        setFilterByFeature(false);
        setFilterByAssigned(false);
        setFilterByUnassigned(false);
        break;
      case "feature":
        setFilterByFeature(!filterByFeature);
        setFilterByBug(false);
        setFilterByAssigned(false);
        setFilterByUnassigned(false);
        break;
      case "assigned":
        setFilterByAssigned(!filterByAssigned);
        setFilterByUnassigned(false);
        setFilterByBug(false);
        setFilterByFeature(false);
        break;
      case "unassigned":
        setFilterByUnassigned(!filterByUnassigned);
        setFilterByAssigned(false);
        setFilterByBug(false);
        setFilterByFeature(false);
        break;
      default:
        break;
    }
  };

  const renderTaskList = () => {
    return (
      <div>
        <label className="checkbox-container">
          Assigned
          <input
            type="checkbox"
            checked={filterByAssigned}
            onChange={() => handleFilterChange("assigned")}
          />
        </label>
        <label className="checkbox-container">
          Unassigned
          <input
            type="checkbox"
            checked={filterByUnassigned}
            onChange={() => handleFilterChange("unassigned")}
          />
        </label>
        <label className="checkbox-container">
          Bugs
          <input
            type="checkbox"
            checked={filterByBug}
            onChange={() => handleFilterChange("bug")}
          />
        </label>
        <label className="checkbox-container">
          Features
          <input
            type="checkbox"
            checked={filterByFeature}
            onChange={() => handleFilterChange("feature")}
          />
        </label>
        {taskList.length === 0 ? (
          <p>No tasks available.</p>
        ) : (
          taskList.map((task) => (
            <div className="list" key={task.name}>
              <p>{task.name}</p>
              <button onClick={() => handleAssignClick(task.name)}>
                Assign
              </button>
              <button onClick={() => handleDeleteClick(task.name)}>
                Delete
              </button>
              <button onClick={() => handleEditClick(task.name)}>Edit</button>
            </div>
          ))
        )}
        {showCreateTaskModal && (
          <div className="modal-overlay">
            <div className="modal">
              <h4>Create New Task</h4>
              <label>
                Task Name:
                <input
                  type="text"
                  value={createTaskName}
                  onChange={(e) => setCreateTaskName(e.target.value)}
                />
              </label>
              <label>
                Description:
                <textarea
                  value={createTaskDescription}
                  onChange={(e) => setCreateTaskDescription(e.target.value)}
                />
              </label>
              <label>
                Reporter:
                <select
                  value={createTaskReporter}
                  onChange={(e) => setCreateTaskReporter(e.target.value)}
                >
                  <option value="">Select Reporter</option>
                  {userList.map((user) => (
                    <option key={user.name}>{user.username}</option>
                  ))}
                </select>
              </label>
              <label>
                Category:
                <div>
                  <label className="checkbox-container-modal">
                    Bug
                    <input
                      type="radio"
                      value="Bug"
                      checked={createTaskCategory === "Bug"}
                      onChange={() => setCreateTaskCategory("Bug")}
                    />
                  </label>
                  <label className="checkbox-container-modal">
                    Feature
                    <input
                      type="radio"
                      value="Feature"
                      checked={createTaskCategory === "Feature"}
                      onChange={() => setCreateTaskCategory("Feature")}
                    />
                  </label>
                </div>
              </label>
              <button onClick={handleCreateTaskSubmit}>Submit</button>
              <button onClick={() => setShowCreateTaskModal(false)}>
                Cancel
              </button>
            </div>
          </div>
        )}
        {showEditModal && (
          <div className="modal-overlay">
            <div className="modal">
              <h4>Edit Task</h4>
              <label>
                New Name:
                <input
                  type="text"
                  value={newName}
                  onChange={(e) => setNewName(e.target.value)}
                />
              </label>
              <label>
                New Description:
                <textarea
                  value={newDescription}
                  onChange={(e) => setNewDescription(e.target.value)}
                />
              </label>
              <button onClick={handleEditSubmit}>Submit</button>
              <button onClick={() => setShowEditModal(false)}>Cancel</button>
            </div>
          </div>
        )}
        {showDeleteModal && (
          <div className="modal-overlay">
            <div className="modal">
              <h4>Delete Task</h4>
              <p>Are you sure you want to delete this task?</p>
              <button onClick={handleDeleteSubmit}>Yes</button>
              <button onClick={handleDeleteCancel}>No</button>
            </div>
          </div>
        )}
        {showAssignModal && (
          <div className="modal-overlay">
            <div className="modal">
              <h4>Assign Task</h4>
              <label>
                Select User:
                <select
                  value={selectedUser}
                  onChange={(e) => setSelectedUser(e.target.value)}
                >
                  <option value="">Select User</option>
                  {userList.map((user) => (
                    <option key={user.name}>{user.username}</option>
                  ))}
                </select>
              </label>
              <button onClick={handleAssignSubmit}>Assign</button>
              <button onClick={() => setShowAssignModal(false)}>Cancel</button>
            </div>
          </div>
        )}
      </div>
    );
  };

  const handleCreateTaskClick = () => {
    setShowCreateTaskModal(true);
  };

  const handleCreateTaskSubmit = async () => {
    try {
      const requestBody = JSON.stringify({
        Task: {
          name: createTaskName,
          description: createTaskDescription,
          category: createTaskCategory,
        },
        User: { username: createTaskReporter },
      });

      const response = await fetch("/api/createTask", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: requestBody,
      });

      if (response.ok) {
        console.log("Task created successfully");
        fetchTaskList();
      } else {
        console.error("Failed to create task");
      }
    } catch (error) {
      console.error("Error creating task:", error);
    }

    setShowCreateTaskModal(false);
    setCreateTaskName("");
    setCreateTaskDescription("");
    setCreateTaskReporter("");
    setCreateTaskCategory("");
  };

  const handleAssignClick = async (taskName) => {
    setEdittaskName(taskName);

    try {
      const response = await fetch("/api/users");
      if (response.ok) {
        const data = await response.json();
        setUserList(data);
      } else {
        console.error("Error fetching user list");
      }
    } catch (error) {
      console.error("Error fetching user list:", error);
    }

    setShowAssignModal(true);
  };

  const handleAssignSubmit = async () => {
    try {
      const selectedTask = taskList.find((task) => task.name === edittaskName);

      if (!selectedTask) {
        console.error(`Task with ID ${edittaskName} not found`);
        return;
      }

      const requestBody = JSON.stringify({
        User: { username: selectedUser },
        Task: { name: selectedTask.name },
      });

      const response = await fetch(`/api/assignTask`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: requestBody,
      });

      if (response.ok) {
        console.log(
          `Task with ID ${selectedTask.id} assigned successfully to User ID ${selectedUser}`
        );
        fetchTaskList();
      } else {
        console.error(
          `Failed to assign task with ID ${selectedTask.id} to User ID ${selectedUser}`
        );
      }
    } catch (error) {
      console.error("Error assigning task:", error);
    }

    setShowAssignModal(false);
    setEdittaskName(null);
    setSelectedUser(null);
  };

  const handleDeleteClick = (taskName) => {
    setEdittaskName(taskName);
    setShowDeleteModal(true);
  };

  const handleDeleteSubmit = async () => {
    try {
      const selectedTask = taskList.find((task) => task.name === edittaskName);

      if (!selectedTask) {
        console.error(`Task ${edittaskName} not found`);
        return;
      }

      const requestBody = JSON.stringify({
        name: selectedTask.name,
      });

      const response = await fetch(`/api/deleteTask`, {
        method: "DELETE",
        headers: {
          "Content-Type": "application/json",
        },
        body: requestBody,
      });

      if (response.ok) {
        console.log(`Task with ID ${selectedTask.id} deleted successfully`);
        fetchTaskList();
      } else {
        console.error(`Failed to delete task with ID ${selectedTask.id}`);
      }
    } catch (error) {
      console.error("Error deleting task:", error);
    }

    setShowDeleteModal(false);
    setEdittaskName(null);
  };

  const handleDeleteCancel = () => {
    setShowDeleteModal(false);
    setEdittaskName(null);
  };

  const handleEditClick = (taskName) => {
    setEdittaskName(taskName);
    setShowEditModal(true);

    const selectedTask = taskList.find((task) => task.name === taskName);
    if (selectedTask) {
      setNewName(selectedTask.name);
      setNewDescription(selectedTask.description);
    }
  };

  const handleEditSubmit = async () => {
    try {
      const selectedTask = taskList.find((task) => task.name === edittaskName);

      if (!selectedTask) {
        console.error(`Task ${edittaskName} not found`);
        return;
      }

      const requestBody = JSON.stringify({
        taskName: selectedTask.name,
        name: newName,
        description: newDescription,
      });

      const response = await fetch(`/api/updateTask`, {
        method: "PUT",
        headers: {
          "Content-Type": "application/json",
        },
        body: requestBody,
      });

      if (response.ok) {
        console.log(`Task with ID ${selectedTask.id} edited successfully`);
        fetchTaskList();
      } else {
        console.error(`Failed to edit task with ID ${selectedTask.id}`);
      }
    } catch (error) {
      console.error("Error editing task:", error);
    }

    setShowEditModal(false);
    setEdittaskName(null);
    setNewName("");
    setNewDescription("");
  };
  
  const renderUserList = () => {
    return (
      <div>
        {userList.length === 0 ? (
          <p>No users available.</p>
        ) : (
          userList.map((user) => (
            <div className="list" key={user.username}>
              <p> {user.username}</p>
              <p>
                {" "}
                {user.name} {user.surname}{" "}
              </p>
              <button onClick={() => handleShowTasksClick(user.username)}>
                Show Assigned
              </button>
              <button onClick={() => handleShowReportedClick(user.username)}>
                Show Reported
              </button>
            </div>
          ))
        )}
  
        {showTasksModal && (
          <div className="modal-overlay">
            <div className="modal">
              <h4>Tasks for User: {selectedUsername}</h4>
              {tasksForUser ? (
                <ul>
                  {tasksForUser.map((task) => (
                    <li key={task.name}>
                      {task.name}
                      {filterAssignButton && (
                        <button onClick={() => handleUnassignTask(task.name)}>
                          Unassign
                        </button>
                      )}
                    </li>
                  ))}
                </ul>
              ) : (
                <p>No tasks for this user.</p>
              )}
              <button onClick={handleExitTasksView}>Exit</button>
            </div>
          </div>
        )}
      </div>
    );
  };
  
  const handleShowTasksClick = async (username) => {
    try {
      const requestBody = JSON.stringify({ username });

      const response = await fetch("/api/getUsersTasks", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: requestBody,
      });

      if (response.ok) {
        const data = await response.json();
        setTasksForUser(data);
        setSelectedUsername(username);
        setShowTasksModal(true);
      } else {
        console.error(
          `Failed to fetch tasks for user with username ${username}`
        );
      }
    } catch (error) {
      console.error("Error fetching tasks:", error);
    }
    setButtonAssign(true);
  };

  const handleUnassignTask = async (taskName) => {
    try {
      const requestBody = JSON.stringify({
        Task: { name: taskName },
        User: { username: selectedUsername },
      });
  
      const response = await fetch("/api/unassignTask", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: requestBody,
      });
  
      if (response.ok) {
        console.log(
          `Task "${taskName}" unassigned successfully for user "${selectedUsername}"`
        );
  
        setTasksForUser((prevTasks) =>
          prevTasks.filter((task) => task.name !== taskName)
        );
      } else {
        console.error(
          `Failed to unassign task "${taskName}" for user "${selectedUsername}"`
        );
      }
    } catch (error) {
      console.error("Error unassigning task:", error);
    }
  };
  

  const handleShowReportedClick = async (username) => {
    try {
      const requestBody = JSON.stringify({ username: username });

      const response = await fetch("/api/getReportedTasks", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: requestBody,
      });

      if (response.ok) {
        const data = await response.json();
        setTasksForUser(data);
        setSelectedUsername(username);
        setShowTasksModal(true);
      } else {
        console.error(
          `Failed to fetch tasks for user with username ${username}`
        );
      }
    } catch (error) {
      console.error("Error fetching tasks:", error);
    }
    setButtonAssign(false);
  };

  const handleExitTasksView = () => {
    setShowTasksModal(false);
    setTasksForUser([]);
    setSelectedUsername(null);
    setButtonAssign(null);
  };

  return (
    <div>
      <h2>
        Projekt 'Przetwarzanie danych w chmurach obliczeniowych - Krzeminska' 2023
      </h2>
      <h3> Repository <a href="https://github.com/Kszemka/PseudoJIRA.git">GitHub</a></h3>
      <div>
        <button onClick={() => handleTabClick("task")}>Task List</button>
        <button onClick={() => handleTabClick("users")}>User List</button>
        {selectedTab === "task" && (
          <button onClick={handleCreateTaskClick}>Create New Task</button>
        )}
        {selectedTab === "task" ? renderTaskList() : null}
      </div>
      {selectedTab === "users" ? renderUserList() : null}
    </div>
  );
};

export default HomePage;