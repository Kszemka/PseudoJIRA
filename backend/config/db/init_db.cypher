CREATE CONSTRAINT IF NOT EXISTS FOR (u:USER) REQUIRE (u.username) IS UNIQUE;
CREATE CONSTRAINT IF NOT EXISTS FOR (t:TASK) REQUIRE (t.name) IS UNIQUE;
CREATE CONSTRAINT IF NOT EXISTS FOR (c:CATEGORY) REQUIRE (c.name) IS UNIQUE;

// Create Users with Names and Surnames
CREATE (user1:USER {username: "alice_smith", name: "Alice", surname: "Smith"});
CREATE (user2:USER {username: "bob_jones", name: "Bob", surname: "Jones"});
CREATE (user3:USER {username: "charlie_davis", name: "Charlie", surname: "Davis"});
CREATE (user4:USER {username: "david_miller", name: "David", surname: "Miller"});
CREATE (user5:USER {username: "eve_jackson", name: "Eve", surname: "Jackson"});

// Create Categories
CREATE (bugCategory:CATEGORY {name: "Bug"});
CREATE (featureCategory:CATEGORY {name: "Feature"});

// Create Bugs with Specific Names
CREATE (bug1:TASK {name: "bug/Fix Database Tables", description: "Database tables need fixing"});
CREATE (bug2:TASK {name: "bug/Resolve Authentication Issue", description: "Authentication problem in the app"});
CREATE (bug3:TASK {name: "bug/Optimize Query Performance", description: "Queries are running slowly"});
CREATE (bug4:TASK {name: "bug/Address UI Bug", description: "UI elements misaligned"});
CREATE (bug5:TASK {name: "bug/Investigate Memory Leak", description: "Memory consumption issue"});

// Connect Bugs to Bug Category
MATCH (bug:TASK), (bugCategory:CATEGORY {name: "Bug"})
WHERE bug.name STARTS WITH "bug/"
CREATE (bug)-[:BELONGS_TO]->(bugCategory);

// Create Features with Specific Names
CREATE (feature1:TASK {name: "feature/Add Accessibility for App", description: "Improve app accessibility"});
CREATE (feature2:TASK {name: "feature/Implement Dark Mode", description: "Add dark mode feature to the app"});
CREATE (feature3:TASK {name: "feature/Enhance User Profiles", description: "Add more details to user profiles"});
CREATE (feature4:TASK {name: "feature/Integrate Third-party API", description: "Incorporate external API into the app"});
CREATE (feature5:TASK {name: "feature/Implement Push Notifications", description: "Enable push notifications"});

// Connect Features to Feature Category
MATCH (feature:TASK), (featureCategory:CATEGORY {name: "Feature"})
WHERE feature.name STARTS WITH "feature/"
CREATE (feature)-[:BELONGS_TO]->(featureCategory);

// Assign and Report Tasks to Users

MATCH (user:USER {username: "bob_jones"}), (bugTask:TASK {name: "bug/Fix Database Tables"})
CREATE (user)-[:REPORT]->(bugTask);
MATCH (user:USER {username: "alice_smith"}), (bugTask:TASK {name: "bug/Resolve Authentication Issue"})
CREATE (user)-[:REPORT]->(bugTask);
MATCH (user:USER {username: "alice_smith"}), (bugTask:TASK {name: "bug/Optimize Query Performance"})
CREATE (user)-[:REPORT]->(bugTask);
MATCH (user:USER {username: "charlie_davis"}), (bugTask:TASK {name: "bug/Address UI Bug"})
CREATE (user)-[:REPORT]->(bugTask);
MATCH (user:USER {username: "charlie_davis"}), (bugTask:TASK {name: "bug/Investigate Memory Leak"})
CREATE (user)-[:REPORT]->(bugTask);
MATCH (user:USER {username: "charlie_davis"}), (featureTask:TASK {name: "feature/Add Accessibility for App"})
CREATE (user)-[:REPORT]->(featureTask);
MATCH (user:USER {username: "bob_jones"}), (featureTask:TASK {name: "feature/Implement Dark Mode"})
CREATE (user)-[:REPORT]->(featureTask);
MATCH (user:USER {username: "david_miller"}), (featureTask:TASK {name: "feature/Enhance User Profiles"})
CREATE (user)-[:REPORT]->(featureTask);
MATCH (user:USER {username: "bob_jones"}), (featureTask:TASK {name: "feature/Integrate Third-party API"})
CREATE (user)-[:REPORT]->(featureTask);
MATCH (user:USER {username: "david_miller"}), (featureTask:TASK {name: "feature/Implement Push Notifications"})
CREATE (user)-[:REPORT]->(featureTask);


MATCH (user:USER {username: "alice_smith"}), (bugTask:TASK {name: "bug/Fix Database Tables"})
CREATE (user)-[:ASSIGN]->(bugTask);
MATCH (user:USER {username: "david_miller"}), (featureTask:TASK {name: "feature/Add Accessibility for App"})
CREATE (user)-[:ASSIGN]->(featureTask);
MATCH (user:USER {username: "charlie_davis"}), (bugTask:TASK {name: "bug/Fix Database Tables"})
CREATE (user)-[:ASSIGN]->(bugTask);
MATCH (user:USER {username: "bob_jones"}), (featureTask:TASK {name: "feature/Enhance User Profiles"})
CREATE (user)-[:ASSIGN]->(featureTask);
MATCH (user:USER {username: "david_miller"}), (featureTask:TASK {name: "feature/Implement Push Notifications"})
CREATE (user)-[:ASSIGN]->(featureTask);
MATCH (user:USER {username: "eve_jackson"}), (featureTask:TASK {name: "feature/Integrate Third-party API"})
CREATE (user)-[:ASSIGN]->(featureTask);
MATCH (user:USER {username: "eve_jackson"}), (bugTask:TASK {name: "bug/Address UI Bug"})
CREATE (user)-[:ASSIGN]->(bugTask);