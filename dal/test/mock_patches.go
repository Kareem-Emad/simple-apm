package daltest

import "github.com/Kareem-Emad/simple-apm/dal"

var requests = []dal.RequestStats{{
	URL:                "https://google.com",
	CreatedAt:          "2020-04-02 00:10:01",
	Service:            "test_s",
	Status:             500,
	Method:             "GET",
	TimeInMilliseconds: 3541,
}, {
	URL:                "https://facebook.com",
	CreatedAt:          "2020-04-05 00:10:01",
	Service:            "test_x",
	Status:             401,
	Method:             "POST",
	TimeInMilliseconds: 6652,
}}

var insertStatments = []string{
	"INSERT INTO simple_apm.request_info (service_name, created_at,  method, url, status, response_time)   VALUES ('test_s', '2020-04-02 00:10:01', 'GET', 'https://google.com', 500, 3541)",
	"INSERT INTO simple_apm.request_info (service_name, created_at,  method, url, status, response_time)   VALUES ('test_x', '2020-04-05 00:10:01', 'POST', 'https://facebook.com', 401, 6652)",
}
