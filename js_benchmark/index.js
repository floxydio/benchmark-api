const express = require("express");
const mysql = require("mysql2");
const dotenv = require("dotenv");
const app = express();
const port = 3001;

dotenv.config();

let connection;

function connect() {
    connection = mysql.createConnection({
        host: process.env.DB_IP,
        user: process.env.DB_USERNAME,
        password: process.env.DB_PASSWORD,
        database: process.env.DB_NAME,
    });

    connection.connect((err) => {
        if (err) {
            console.error("Error connecting to the database:", err.message);
            process.exit(1); 
        } else {
            console.log("Connected to the database");
        }
    });
}

connect();

app.get("/story", (req, res) => {
    const { sortBy } = req.query;
    let queryData = "";

    if (sortBy === "DESC") {
        queryData = "SELECT * FROM story ORDER BY story_id DESC LIMIT 20";
    } else {
        queryData = "SELECT * FROM story ORDER BY story_id ASC LIMIT 20";
    }

    connection.query(queryData, (err, results) => {
        if (err) {
            console.error("Query error:", err.message);
            res.status(500).json({ error: "Internal Server Error" });
        } else {
            res.status(200).json(results);
        }
    });
});

app.listen(port, () => {
    console.log(`Server is running on port ${port}`);
});
