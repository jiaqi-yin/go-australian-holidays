# go-australian-holidays
This app provides an endpoint to check Australian holidays for states

## Usage

```
curl --request GET 'http://localhost:8080/v1/is_holiday' \
--header 'Content-Type: application/json' \
--data-raw '{
    "state": "vic",
    "date": "20221225"
}'
```
