### TechTrend
This is a mono-repo of tech-trend. Here will contain all backend and front-end files. This readme looks sucks now, I'll update later


#### Dir Structure

 - collector [where we scrape,parse and the data and store into a CSV. Then to a DB]
 - frontend [where all the front-end code lives]
 - backend [APIs lives there]

 ### Credentials: 
 This clickhouse installation included with [TabiX](https://tabix.io "Visit TabiX website")
 DB: http://localhost:8123
 host: http://127.0.0.1:8123 
 user: default
 pass: admin

## Roadmap (v1)
- [x]  Scrape count from BDJobs
- [x] Save max count for a day in Database
- [ ] Cron job or serverless function for invoking the collector
- [ ] Backend API for showing the graph
- [ ] Frontend for the graph with comparison
- [ ] Tracing for Backend API
- [x] Tracing for Collector

 #### Tech Stack
 - [ ] Backend
   - Golang
   - PostgresSQL
   - APIs are built upon [Echo Framework](https://echo.labstack.com)
   - Database query are using [Gorm](https://gorm.io/index.html)
   - Bugsnag for bug reporting
 - [ ] Collector CLI
   - Golang
   - PostgresSQL
   - Bugsnag for bugs
 - [ ] Frontend
   - Nextjs
   - Typescript