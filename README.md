# avito-test-task
Marketplace Json API

##Технологии

* Go1.13
* PostgreSQL+PlpgSQL
* Docker
* easyjson, для сериализации структур
* Go tests

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
    codes 200, 500 
        body
        {
            "id": {$int64},
            "error":{$string} //optional, if code != 200 OK
        }     
        
#### Get advertisments page
    URL : /api/adverts
    params
        page = $int64
        order_date={'asc','desc'}
                OR 
        order_price={'asc','desc'}

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
    
    /api/advert/{id}/
    
    params
        fields=[pictures,description]
    
    Response
    {
       "id": 10360,
       "title": "Продам гараж",
       "description": "Обычный гараж",
       "pictures": [
           "http://pics.com/main_pic.jpg",
           "http://pics.com/seco_pic.jpg",
           "http://pics.com/thir_pic.jpg"
       ],
       "main_picture": "http://pics.com/main_pic.jpg",
       "price": 1233
    }
#### Install docker
    sudo apt-get update
    sudo apt-get install -y \
        apt-transport-https \
        ca-certificate \
        curl \
        gnupg-agent \
        software-properties-common
              
    curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
  
    sudo apt-key fingerprint 0EBFCD88   
    sudo add-apt-repository \
       "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
       $(lsb_release -cs) \
       stable"
####    
    sudo apt-get update
    sudo apt-get install docker-ce docker-ce-cli containerd.io
####   
    cd /path/to/repo/
    docker-compose up --build --force-recreate --remove-orphans
