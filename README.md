# LINKSTER API

API endpoint: `http://139.59.209.223`     
     

### Create a new event 
| Endpoint                | Request | Example Body                                                                                                       |
|-------------------------|---------|---------------------------------------------------------------------------------------------------------------------|
| `/api/v1/create/event`  | POST    | ```json
{
    "name": "Beff Jezos",
    "wallet": "0x12312312312312123123",
    "profession": "CEO",
    "company": "Amazon",
    "telegram": "beffy"
}
``` |
| `/api/v1/create/profile`  | POST    | ```json
{
    "event": {
        "event-name": "Hackathon somewhere on Chirstmas",
        "event-date": "2023-12-24T01:02:00.977+01:00",
        "event-location": "Nowhere",
        "additional-info": "hello welcome to this event please have fun and everything bla bla bla bla bla"
    }
}
``` |
