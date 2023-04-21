# URL Shortener API

Simple URL shortener API written in Go using the Gin framework. It allows users to generate short URLs for long URLs and redirect users to the original long URLs using the short URLs.

## Installation

1. Install Go on your machine. You can download Go from the official website.
2. Clone the repository to your local machine
3. Set up a MySQL database and create a table called short_urls with the following columns:
```sql
Copy code
CREATE TABLE short_urls (
  id INT PRIMARY KEY AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  tags VARCHAR(255),
  destination_url VARCHAR(100) NOT NULL,
  back_half VARCHAR(20) UNIQUE NOT NULL,
  created_at DATETIME NOT NULL,
  updated_at DATETIME NOT NULL
);
```
4. Replace .env_example into .env  
5. Update the database configuration in the .env file to match your MySQL database credentials