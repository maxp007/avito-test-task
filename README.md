# avito-test-task
### Marketplace Json API

#### Demo server http://194.67.90.7:8080/api/adverts (see api methods below)

## Технологии
* Go 1.13
* PostgreSQL + PlpgSQL
* Redis
* Docker-Compose
* Easyjson
* Go testsing, Unit-тесты

### How To Run (Ubuntu 19.10)
    sudo apt update
    sudo apt-get install -y docker.io
    sudo apt-get install -y docker-compose
    sudo apt-get install -y git
    git clone https://github.com/maxp007/avito-test-task
    cd avito-test-task
    sudo docker-compose up --build 
## API methods
#### Create new advertisment

    POST /api/create
        {
        "title":"Продам гараж",
        "description":"Обычный гараж",
        "pictures":[
            "http://pics.com/main_pic.jpg",
            "http://pics.com/seco_pic.jpg",
            "http://pics.com/thir_pic.jpg"
        ],
        "price":1233
        }
        
    Response
        body
        {
            "id": {$int64},
            "error":{$string} //optional, if errors occured
        }     
        
#### Get advertisments page
    GET /api/adverts
    params
        page = $int64
        order={'date','price'}
        sort={'asc','desc'}

    Response
    {
        "adverts": [
            {
                "id": 10360,
                "title": "Продам гараж",
                "main_picture": "http://pics.com/main_pic.jpg",
                "price": 1233
            },
            {...},
            ]
    }        
    
#### Get advertisment by ID
    
    GET /api/advert/{id}/
    params
        fields=[pictures,description]
    
    Response
    {
       "id": 10360,
       "title": "Продам гараж",
       "description": "Обычный гараж",       //optional
       "pictures": [                         //optional
           "http://pics.com/main_pic.jpg",
           "http://pics.com/seco_pic.jpg",
           "http://pics.com/thir_pic.jpg"
       ],
       "main_picture": "http://pics.com/main_pic.jpg",
       "price": 1233
    }

