## Introduction👨‍💻

This program is written with GO to show how to read csv files and register them in the database

<a href="https://www.coffeebede.com/poulad"><img class="img-fluid" src="https://coffeebede.ir/DashboardTemplateV2/app-assets/images/banner/default-yellow.svg" /></a>
## Usage

**flags**
```
-config=config.json or --config=config.json
-file=file.csv or --file.csv=file.csv
```
config includes db information for insert csv
```json
{
  "username": "postgres",
  "password": "postgres",
  "host": "localhost",
  "name": "exel2sqldb",
  "ssl": "disable"
}
```

## Examples

**Windows**
```
go run main.go -file=C:\file\path\username.csv

OR

.\build.exe -file=C:\file\path\username.csv
```

**Linux**
```
go run main.go -file=/file/path/username.csv

OR

wine start -file=/file/path/username.csv
```


## Tech Stack

**Database:** postgresql

**Lang:** GO

    github.com/savioxavier/termlink v1.3.0
    github.com/lib/pq


## Authors

- [@pooulad](https://www.github.com/pooulad)
- [@hossein1376](https://www.github.com/hossein1376)
  
## Demo

![Screenshot 2023-10-24 211725](https://github.com/pooulad/csv-2-sql-example/assets/86445458/57dd39e1-9929-4a8c-8223-6f83167d23de)

