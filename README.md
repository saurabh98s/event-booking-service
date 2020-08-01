
## Sample Json Post request

{
    "Name": "opera aidda",
    "Startdate": 2343423,
    "Enddate": 324984,
    "Duration": 120,
    "Location": {
        "Name": "Opera House",
        "Address": "11, west street austrailia",
        "Country": "Austrialia",
        "Opentime": 7,
        "Closetime": 20,
        "Hall": [
            {
                "Name": "Cesar hall",
                "Location": "Second floor",
                "Capacity": 10
            },{
                "Name": "Montana hall",
                "Location": "Third floor",
                "Capacity": 1000
            }
        ]
    }
}`


## RABBIT MQ

At this point of the development we have started working on the message broker service <br>
RabbitMQ is a message broker: it accepts and forwards messages. You can think about it as a post office: when you put the mail that you want posting in a post box, you can be sure that Mr. or Ms. Mailperson will eventually deliver the mail to your recipient. In this analogy, RabbitMQ is a post box, a post office and a postman.

`go get -u github.com/streadway/amq`