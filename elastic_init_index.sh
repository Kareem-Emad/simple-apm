echo "creating index request_info ..."
curl -X PUT "localhost:9200/request_info?pretty"

echo "seeding index mapping schema ...."
curl -X PUT "localhost:9200/request_info/_mapping?pretty" -H 'Content-Type: application/json' -d'
{
  "properties": {
    "created_at": {
      "type": "date", 
      "format": "yyyy-MM-dd HH:mm:ss"
    },
    "service_name": {
        "type": "keyword"
    },
    "url": {
        "type": "text"
    },
    "status_code": {
        "type": "short"
    },
    "response_time": {
        "type": "long"
    },
    "http_method": {
        "type": "keyword"
    }
  }
}
'
