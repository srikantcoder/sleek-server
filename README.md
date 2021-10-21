# sleek-backend
Backend service implementation of the Sleek Cashback application

The service exposes following APIs:
1. /deals : API to retrieve available deals at any moment of time
	The deals are returned as an array of json objects with following schema:
  	{
  		deal_id: UUID string representing this deal uniquely amongst all deals.
			Type: String

  		retailer_id: UUID string representing which retailer is offering this deal.
			Type: String

  		retailer_name: String name of the retailer offering this deal.
                    	Type: String

  		retailer_domains: Array of strings indicating which domains the retailer is offering this deal on. The string will be the host name of the URL and if that host matches the host of a webpage, the deal is valid on that page.
                    e.g. ["bestbuy.com", "bestbuy.ca"]
                    Type: Array<String>

  		deal_type: One of "FIXED" or "PERCENTAGE" for a fixed dollar amount off the total purchase or a percentage amount off the total purchase, respectively. All deals given in deals.json will be "PERCENTAGE".
                    Type: "FIXED" | "PERCENTAGE"

  		deal_amount: If |deal_type| is "FIXED", this value will be a positive floating point number to 2 decimal points indicating how many dollars and cents are discounted off of this purchase the deal gives the user.
                    If |deal_type| is "PERCENTAGE", this value will be a floating point number in the range [0, 1] indicating how many percentage points are discounted off of this purchase. e.g. 0.1 = 10% off of the purchase.
			Type: Number
	}

2. /activate/{id}: A POST API call provided by the server to activate deal with id 'id'

3. /deals-status: A GET API call provided by the server to get the count of users who have activated the particular deal.

How to run the server?

1. Backend server is backed by mysql server. Therefore before running the backend server, one needs to run mysql server.
2. The existing code expects the mysql credentials to be <username: root, password: password>. Also, it expects a database by name 'sleek' to be present. If your mysql server credentials are different, please update the server code's clients/mysql_manager.go file dsn value (root:password@tcp(localhost:3306)/sleek?charset=utf8mb4&parseTime=True&loc=Local) accordingly.
3. Move to the cashback folder and run 'go run main.go' to start the server. The server starts listening at port 8082 of localhost.
