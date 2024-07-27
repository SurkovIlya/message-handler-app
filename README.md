# message-handler-app

The application is running on a remote server. To access its methods, use {message-handler-app-host} = 79.174.85.168

## A local option to launch the application and work with it:
### Requirements
* Docker and Go
### Usage
Clone the repository with:
```bash
git clone github.com/SurkovIlya/message-handler-app
```
Copy the `env.example` file to a `.env` file.
```bash
cp .env.example .env
```
Update the postgres variables declared in the new `.env` to match your preference. 

Build and start the services with:
```bash
docker-compose up --build
```
### Message-handler-app API
<details>
<summary> <h4>{message-handler-app-host}/v1/receiving - send message in app</h4></summary>
  
#### Method: POST
#### Request: 
```json
{
	"value": "New message..."
}
```
#### Response:
```json
"OK"
```
</details>
<details>
<summary> <h4>{message-handler-app-host}/v1/getstatistics - Getting statistics on processed messages</h4></summary>
  
#### Method: GET

#### Response:
```json
{
	"handled": 7, // - count of messages processed
	"inProcess": 0 // - count of unprocessed messages
}
```
</details>



