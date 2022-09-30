# Test statistics service for BCraft

### Installation
#### Build a project:
```
docker-compose up -d --build
```

### API documentation
### Statistics creation API
```
POST <host>/statistics
Creates a statistic based on several inputs
> Views, clicks and cost are optional fields
```
#### JSON Body
```json
{
    "date": "YYYY-MM-dd",
    "views": 2500,
    "clicks": 1200,
    "cost": 2500
}
```
### Statistics List API
```
Shows all statistics based on search parameters or all statistics 
if no parameters given then by default will be sorted by the date
GET <host>/statistics?from=YYYY-MM-dd&to=YYYY-MM-dd&sort_by=<one_param: any field>
```
#### JSON Body
```json
{
  "statistics": [
    {
      "date": "2022-08-03",
      "views": 12,
      "clicks": 32,
      "cost": 1.1,
      "cpc": 0.034375,
      "cpm": 91.66667
    },
    {
      "date": "2022-07-03",
      "views": 123,
      "clicks": 321,
      "cost": 1.11,
      "cpc": 0.003457944,
      "cpm": 9.02439
    },
    {
      "date": "2022-01-03",
      "views": 1,
      "clicks": 1,
      "cost": 100,
      "cpc": 100,
      "cpm": 100000
    },
    {
      "date": "2021-02-03",
      "cpc": 0,
      "cpm": 0
    },
    {
      "date": "2021-01-03",
      "clicks": 1,
      "cost": 100,
      "cpc": 100,
      "cpm": 0
    }
  ]
}
```
#### Statistics Reset API
```
DELETE <host>/statistics
Delete all statistics from database
```